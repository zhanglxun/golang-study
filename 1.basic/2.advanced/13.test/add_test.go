package main

import "testing"

/**
go test 加包的名称，执行包下所有测试用例
go test 加测试源文件，执行所有测试用例
go test -run 选项 执行指定的测试用例
*/

/**
以_test.go 结尾
*testing.T  函数的单元测试，参数
*testing.B  机制测试或压力测试
*/
func TestAdd(t *testing.T) {
	var a = 10
	var b = 20
	t.Logf("a=%d b=%d \n", a, b)

	c := Add(a, b)
	if c != 30 {
		t.Fatalf("invalid a + b, c=%d", c)
	}
	t.Logf("a = %d b = %d sum=%d\n", a, b, c)
}

func TestSub(t *testing.T) {
	var a = 10
	var b = 20
	t.Logf("a = %d b = %d\n", a, b)
	c := Sub(a, b)
	if c != -10 {
		t.Fatalf("invalid a - b, c=%d", c)
	}
	t.Logf("a = %d b = %d sub=%d\n", a, b, c)
}
