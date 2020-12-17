package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

// Server HTTP Server
type Server struct {
	Addr string
	http *http.Server
}

// NewServer 创建一个 HTTP Server
func NewServer(addr string) *Server {
	mux := http.NewServeMux()
	mux.Handle("/", &handler{})
	http := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	return &Server{
		Addr: addr,
		http: http,
	}
}

// Start 开始服务
func (s *Server) Start() error {
	if err := s.http.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	log.Print("http server exit")
	return nil
}

// Shutdown 停止服务
func (s *Server) Shutdown(ctx context.Context) error {
	log.Print("http server shutdown")
	if err := s.http.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}

// BackgroundWorker 后台任务Worker
type BackgroundWorker struct {
	stop chan struct{}

	ctx    context.Context
	cancel context.CancelFunc
}

// NewBackgroundWorker 创建后台任务Worker
func NewBackgroundWorker() *BackgroundWorker {
	ctx, cancel := context.WithCancel(context.Background())
	return &BackgroundWorker{
		stop: make(chan struct{}),

		ctx:    ctx,
		cancel: cancel,
	}
}

// Start 后台Worker开始工作
func (w *BackgroundWorker) Start() error {
	log.Print("worker start")

	// 退出时通知调用方
	defer close(w.stop)

	log.Print("background worker start")

	// 定时器
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	// 模拟后台任务
	for {
		select {
		case <-ticker.C:
			// 模拟后台任务
			time.Sleep(500 * time.Millisecond)
			log.Printf("background do job")
		case <-w.ctx.Done():
			return nil
		}
	}
}

// Stop 停止后台任务
func (w *BackgroundWorker) Stop(ctx context.Context) error {
	// 让后台任务停下来
	w.cancel()

	select {
	case <-w.stop: // 等待成功退出
	case <-ctx.Done(): // 外部也可以强制退出
		log.Print("background worker exit with deadline")
		return ctx.Err()
	}

	log.Print("background worker exit")
	return nil
}

// Stop 停止 HTTP Server 和 后台任务
func Stop(s *Server, w *BackgroundWorker) error {
	g, _ := errgroup.WithContext(context.Background())

	// 设置5秒强制退出
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 停止后台任务
	g.Go(func() error {
		return w.Stop(ctx)
	})

	// 关闭 HTTP Server
	g.Go(func() error {
		return s.Shutdown(ctx)
	})

	// 等待全部完成
	return g.Wait()
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	// 启动 HTTP Server
	s := NewServer(":8080")
	g.Go(func() error {
		return s.Start()
	})

	// 启动后台任务
	w := NewBackgroundWorker()
	g.Go(func() error {
		return w.Start()
	})

	// 监听退出信号
	g.Go(func() error {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		select {
		case sig := <-sigs:
			// 收到关闭信号
			log.Printf("signal caught: %s, ready to quit...", sig.String())
		case <-ctx.Done():
			// HTTP Server 或 后台任务启动失败 停止监听信号
			log.Println("monitor signal exit")
		}

		// 停止 HTTP Server 和 后台任务
		return Stop(s, w)
	})

	// 等待所有 goroutine 退出
	if err := g.Wait(); err != nil {
		log.Printf("server exit: %+v", err)
		return
	}

	// 服务优雅关闭
	log.Println("server graceful exit")
}
