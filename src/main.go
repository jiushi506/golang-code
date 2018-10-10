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
	"runtime"
	"log"
	"io/ioutil"
	"github.com/Jeffail/tunny"
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
	//pointer_test()
	//panic_test()
	//string_test()
	//value_translate_test()
	//array_slice_difference_test()
	//slice_test()
	//read_memory_stat_test()
	//time_test()
	//bit_calculate_test()
	//new_object_test()
	tunny_goroutine_test()
}

func tunny_goroutine_test() {
	numCPUs := runtime.NumCPU()
	pool := tunny.NewFunc(numCPUs, func(payload interface{}) interface{} {
		var result []byte

		// TODO: Something CPU heavy with payload
		if value,ok := payload.([]uint8);ok {  // interface{} 需要断言才能转换
			fmt.Println(string(value))
		}
		return result
	})
	defer pool.Close()

	http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		input, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
		}
		defer r.Body.Close()

		// Funnel this work into our pool. This call is synchronous and will
		// block until the job is completed.
		result := pool.Process(input)

		w.Write(result.([]byte))
	})

	http.ListenAndServe(":8080", nil)
}

func new_object_test() {
	a := new(int64)
	fmt.Println(a)
}

func bit_calculate_test() {
	var in int64
	tye := reflect.TypeOf(in)
	fmt.Println("输出类型:",tye.Kind())

    bitResult := 1 << 7
	fmt.Println("输出位计算结果",bitResult)

}

func time_test() {
	time1 := time.Now()
	d, _ := time.ParseDuration("-144h")
	time2 := time1.Add(d)
	fmt.Println(time2)

}

func bigBytes() *[]byte {
	s := make([]byte, 100000000)
	return &s
}

func read_memory_stat_test() {
	// 统计内存中在字节数组分配前后的大小对比
	//https://studygolang.com/articles/12008?fr=sidebar
	var mem runtime.MemStats

	fmt.Println("memory baseline...")

	runtime.ReadMemStats(&mem)
	log.Println(mem.Alloc)
	log.Println(mem.TotalAlloc)
	log.Println(mem.HeapAlloc)
	log.Println(mem.HeapSys)

	for i := 0; i < 10; i++ {
		s := bigBytes()
		if s == nil {
			log.Println("oh noes")
		}
	}

	fmt.Println("memory comparison...")

	runtime.ReadMemStats(&mem)
	log.Println(mem.Alloc)
	log.Println(mem.TotalAlloc)
	log.Println(mem.HeapAlloc)
	log.Println(mem.HeapSys)
}

func slice_test() {
	var a = make([]int, 6)
	b := a[1:3]
	fmt.Printf("a's address %p\n",&a[0])
	fmt.Printf("b's address %p\n",&b[0])  //a[0]和b[0]指向的地址差了一个元素的大小，说明reslice后指针直接指向a[low]的地址
	b = append(b,2,3,4,5,6,7) // 元素个数超过cap时会重新分配地址，地址变了
	fmt.Printf("b's address after append %p\n",&b[0])
}
func array_slice_difference_test() {
	//var a [2]int
	var b []int
	c := []int{}
	//fmt.Println(a == nil)  报错，数组不能和nil比较
	fmt.Println(b == nil)
	fmt.Println(c == nil)
}

func value_translate_test() {
	var a = 2
	b := a  // 如果改为 b := &a 那么三四行输出的value就一样了
	//值传递，输出的地址不一样
	fmt.Println("a's address:", &a, " a's value:",a)
	fmt.Println("b's address:", &b, " b's value:",b)

	change_value(&b)
	fmt.Println("a's address:", &a, " a's value:",a)
	fmt.Println("b's address:", &b, " b's value:",b)

}
func change_value(b *int) {
	*b = 3
}
func string_test() {
	s := "hello world"
	s = "hello world2"
	//reflect.ValueOf(&s).Pointer()  // reflect.ValueOf（）参数一定要是指针类型，否则会panic
	fmt.Println(s)

	b := []byte{1,2,3}
	b[1] = 3
	fmt.Println(b)

	str := "welcome to outofmemory.cn"
	fmt.Println("str[",str,"] 替换前地址：v%", &str)
	str = strings.Replace(str, " ", ",", -1)
	fmt.Println("str[",str,"] 替换后地址：v%", &str)
	//输出地址是一样的，？？？？ 这是因为str原先在内存有一个地址，这个是不变的，string改变后，变的是str指向的地址
}

func panic_test() {
	//panic_without_recover()
	panic_with_recover()
}

func panic_with_recover() {
	defer func() {  //recover()用于将panic的信息捕捉。recover必须定义在panic之前的defer语句中。
		if err := recover(); err != nil {  //注意这里用分号隔开，为什么不是 && ?
			fmt.Println("catch the panic:",err)
		}
		fmt.Println("execute 3")
	}()
	fmt.Println("execute 1")
	panic("code panic")
	fmt.Println("execute 2")  //即使recover捕获了异常，这一句也不会执行
}

func panic_without_recover() {
	defer func() {
		fmt.Println("execute 3")
	}()
	fmt.Println("execute 1")
	panic("code panic")
	fmt.Println("execute 2")
}
func pointer_test() {
	a := new(int)
	*a = 3
	fmt.Println("a地址:",&a)   // #1 输出指针自身的地址
	fmt.Println("a指向的地址:",a)  // #2  输出指针变量指向的地址
	invoke_point(a)

}
func invoke_point(b *int) {
	fmt.Println("a参数地址:",&b)  // #3  输出和#1不一样，因为是值传递，复制了一个地址
	fmt.Println("a参数指向的地址:",b)  //#4  输出和#2一样，因为指针指向的地址是一样的
}

func map_test() {
	var m = map[string]string{"name":"zhangsan","pass":"zhangsan'pass"}
	value,ok := m["name"]  //map取值可以返回两个值
	fmt.Println(value,ok)
	value,ok = m["name don't exist"]
	fmt.Println(value,ok)
	fmt.Println(&m)
	fmt.Println(1 << 0) // 1
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
