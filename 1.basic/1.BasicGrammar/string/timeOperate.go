package main

import (
	"fmt"
	"time"
)

// 时间格式打印
func testTime() {

	var now = time.Now()
	fmt.Printf("currenct time :%v \n", now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	fmt.Printf("%02d %02d %02d %02d %02d %02d\n", year, month, day, hour, minute, second)

	timestamp := now.Unix()
	fmt.Printf("%d\n", timestamp)

}

func testStamp(timestamp int64) {
	timeObj := time.Unix(timestamp, 0)

	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()

	fmt.Printf("%d\n", timestamp)
	fmt.Printf("%02d %02d %02d %02d %02d %02d\n", year, month, day, hour, minute, second)

}

//时间 管道 定时器
func testTicker() {
	ticker := time.Tick(2 * time.Second)
	for i := range ticker {
		fmt.Printf("%v \n", i)
		doTask()
	}
}
func doTask() {
	fmt.Println("do this task")
}

func testConst() {
	fmt.Printf("nano second %d\n", time.Nanosecond)
	fmt.Printf("micro second %d\n", time.Microsecond)
	fmt.Printf("milli second %d\n", time.Millisecond)
	fmt.Printf("second %d\n", time.Second)

}

func testFormat() {
	now := time.Now()
	timestr := now.Format("2006/01/02 15:04:05")
	fmt.Printf("time is : %s \n", timestr)
}

//程序执行时间计算
func testCost() {
	start := time.Now().UnixNano()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond)
	}
	end := time.Now().UnixNano()
	cost := (end - start) / 1000
	fmt.Printf("code run cost : %d us \n", cost)
}

func main() {

	//testTime()

	// 时间处理
	//testStamp(time.Now().Unix())

	//通过管道实现的定时器
	//testTicker()

	//时间打印
	//testConst()

	//时间格式 2006-01-02 15：04：05  必须这个格式模板
	//testFormat()

	//程序执行时间计算时长
	testCost()
}
