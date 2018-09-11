package test

import (
	"testing"
	"fmt"
)
// 性能测试  go test -bench="."
func Test_test1(t *testing.T) {
	fmt.Println("execute here")
}

//go test -bench=Benchmark_Division  执行指定方法的性能测试
func Benchmark_Division(t *testing.B) {
    fmt.Println("print message")
}

func Benchmark_method2(t *testing.B) {
	fmt.Println("print message2")
}

