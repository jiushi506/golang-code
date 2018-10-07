package main

import (
	"fmt"
	"sync"
	"game"
	"time"
	"net/http"
	"strings"
	"os"
	"io"
	"reflect"
	"unsafe"
	"github.com/gin-gonic/gin"
	"github.com/DeanThompson/ginpprof"
)

/**
查看文件的编译源码
1. go build && main.go  //这个时候会生成一个 src.exe的文件
2. go tool objdump -s "main.main" src.exe
 */

func main() {
	//s := "hello" //这种简短赋值的方法只能用在函数内部
	//fmt.Println(s)
	//file_test()
	//interface_test()
	//file_create_delete_test()
	//defer_test()
	//channel_test()
	//single_channel_test()
	//select_channel_test()
	//close_channel_test()
	//wait_group_test()
	//unsafe_pointer_uintptr_test()
	//var_test()
	//map_test()
	ginpprof_test();
}

func ginpprof_test() {   //https://github.com/jiushi506/ginpprof
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// automatically add routers for net/http/pprof
	// e.g. /debug/pprof, /debug/pprof/heap, etc.
	ginpprof.Wrap(router)

	// ginpprof also plays well with *gin.RouterGroup
	// group := router.Group("/debug/pprof")
	// ginpprof.WrapGroup(group)

	router.Run(":8080")
}

func map_test() {
	var m = map[string]string{"name":"zhangsan","pass":"zhangsan'pass"}
	value,ok := m["name"]  //map取值可以返回两个值
	fmt.Println(value,ok)
	value,ok = m["name don't exist"]
	fmt.Println(value,ok)
	fmt.Println(&m)
}

func var_test() {
	var a = 3
	const b = 4
	fmt.Println(&a,a)
	//fmt.Println(&b,b)  //不能引用constant的地址，因为常量在编译预处理阶段直接展开，作为指令数据使用，没有地址
	arr := make([]int,1)
	fmt.Println(reflect.TypeOf(arr))
}

func unsafe_pointer_uintptr_test() {
	a := [4]int64{5, 1, 2, 3}
	fmt.Println(unsafe.Sizeof(&a[0]))  // 每个元素占8个字节，64位
	fmt.Println(unsafe.Pointer(&a[0]),unsafe.Pointer(&a[1]),unsafe.Pointer(&a[2]),unsafe.Pointer(&a[3])) //输出每个元素的内存地址 16进制
	p := uintptr(unsafe.Pointer(&a[0]))
	fmt.Println(p)   //输出a[0]的内存地址 10进制
	p1 := unsafe.Pointer(p)
	fmt.Println(*(* int64)(p1))  // 根据p1的地址输出对应的值 （* int64)表示先转化为 int64 指针， * 表示取值
}

func wait_group_test() {
	var waitGroup sync.WaitGroup //定义一个同步等待的组
	waitGroup.Add(1) //添加一个计数
	go game.ConnSocket(&waitGroup) //调用其他包的方法执行任务
	waitGroup.Wait() //阻塞直到所有任务完成
	fmt.Println("main DONE!!!")
}

func close_channel_test() {
	ch := make(chan string,2)
	go func(in chan <- string) {
		in <- "hello"
		//time.Sleep(time.Second * 7) //这里如果睡眠7秒再写入数据就会报错，因为这个时候channel已经被关闭了
		in <- "world"
	}(ch)

	go func(out <- chan string) {
		fmt.Println(<-out)
		time.Sleep(time.Second * 4)
		fmt.Println(<-out)
	}(ch)

	time.Sleep(time.Second * 3)
	close(ch)
	fmt.Println("close channel")
	time.Sleep(time.Second * 5)

}

func select_channel_test() {
	process1 := make(chan string)
	process2 := make(chan string)

	go func (in chan <- string) {
		in <- "hello"
	}(process1)

	go func (in chan <- string) {
		in <- "world"
	}(process2)

	time.Sleep(time.Second * 2) //如果没有这一句，会输出两个default operation,因为两个协程来不及往通道放数据
	for i:=0;i<2;i++{ //如果不加for循环，select找到一个匹配的条件就退出了
		select {
		case str := <- process1:
			fmt.Println("process1 data",str)
		case str := <- process2:
			fmt.Println("process2 data",str)
		default:
			fmt.Println("default operation")
		}
	}

}

func single_channel_test() {
	ch := make(chan string)

	go func(out chan<- string){
		out <- "hello"
	}(ch)

	go func(in <-chan string){
		fmt.Println( <- in)
		//fmt.Println( <- in)
		fmt.Println("execute here") //多加一个fmt.Println( <- in) 这里就阻塞不会执行，一直到main运行结束
	}(ch)
	time.Sleep(2 * time.Second)
}

func channel_test() {
	ch := make(chan string,2)
	//ch := make(chan string)
	go func() {
		// 当channel只能包含一个元素时，在执行到这一步的时候main goroutine才会停止阻塞
		str := <-ch
		fmt.Println("receive data：" + str)
	}()
	ch <- "hello"
	ch <- "world"
	time.Sleep(time.Second)
	fmt.Println("channel has send data")
}

func defer_test() {
	//多个 defer 按照 FILO( First In Last Out)的顺序执行
	defer fmt.Println("before")
	defer fmt.Println("after")
	divide(10, 0)
	fmt.Println("execute end") // divide报错后这句不会执行
}

func divide(a, b int) {
	if b == 0 {
		panic("diveded by zero")
	}
	fmt.Println("a/b = ", a/b)
}

func file_create_delete_test() {
	dir := "f://"
	fileUrl := "https://resourceuat-1252347619.cosgz.myqcloud.com/gift/6c31196c-e399-4986-994e-dc9e2516bb00.png"
	res, err := http.Get(fileUrl)
	if err != nil {
		panic(err)
	}
	lastIndex := strings.LastIndex(fileUrl, "/")
	filename := fileUrl[lastIndex+1 : len(fileUrl)]
	fmt.Println("Download filename --------------:", filename)
	filename = dir + filename
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Create file --------------:", filename)
	io.Copy(f, res.Body)
	f.Close()
	deleteFile(filename)
}

func deleteFile(filename string) {
	time.Sleep(5 * time.Second)
	err := os.Remove(filename)
	if err != nil {
		fmt.Println("delete fail! -->", err.Error())
	} else {
		fmt.Println("delete success ! --> ", filename)
	}
}

func interface_test() {
	type myType int
	var i int
	var j myType
	reflect.ValueOf(i).Interface()
	fmt.Println("i kind:", reflect.TypeOf(i).Kind())
	fmt.Println("j kind:", reflect.TypeOf(j).Kind())
	fmt.Println("j kind:", reflect.TypeOf(j))
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
	fmt.Println("generate value from reflect.Value:", v.Interface().(float64))
}

// https://github.com/cuebyte/The-Laws-of-Reflection   深入理解
func file_test() (res interface{}, err error) {
	var r io.Reader
	tty, err := os.Create("f:/20180827.txt")
	if err != nil {
		return nil, err
	}
	r = tty

	fmt.Println(r.Read)

	var w io.Writer
	w = r.(io.Writer)
	fmt.Println(w.Write)

	var empty interface{}
	empty = w
	fmt.Println(empty)
	return nil, nil
}
