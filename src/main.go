package main

import (
	"flag"
	"fmt"
	"game"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"runtime"
	"runtime/pprof"
	"strconv"
	"strings"
	"sync"
	"time"
	"unsafe"
)

/**
查看文件的编译源码
1. go build && main.go  //这个时候会生成一个 src.exe的文件
2. go tool objdump -s "main.main" src.exe
*/

const (
	January time.Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

var m = map[time.Month]int{
	January:   1,
	February:  2,
	March:     3,
	April:     4,
	May:       5,
	June:      6,
	July:      7,
	August:    8,
	September: 9,
	October:   10,
	November:  11,
	December:  12,
}

type HandleLock struct {
	sync.Mutex
	locked bool
}

type wfeProcInsModelImpl struct {
	sync.Mutex
	lockMap map[int64]HandleLock
}

func (m *wfeProcInsModelImpl) checkLockCap() {
	if len(m.lockMap) >= 3 { // 超过容量，重新分配
		m.Lock()
		if len(m.lockMap) >= 3 {
			for k, v := range m.lockMap {
				if !v.locked {
					delete(m.lockMap, k)
				}
			}
		}
		m.Unlock()
	}
}

var (
	//定义外部输入文件名字
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file.")
)
var i = 0

func main() {
	// 查看输出性能文件        go tool pprof ./Main cpu.pprof  交互式页面输入 web,直接浏览器查看
	f, err := os.Create("cpu.pprof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

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
	//tunny_goroutine_test()
	//time2_test()
	//fmt.Println(wild_code_test())
	//golang_replace_all_test()
	//calculateTime("s b")
	//url_test()
	lock_test()
}

func lock_test() {
	//runtime.GOMAXPROCS(1)
	//arr := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr := []int64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m := wfeProcInsModelImpl{}
	m.lockMap = make(map[int64]HandleLock, 0)
	for _, v := range arr {
		go handle(v, &m)
	}
}

func handle(v int64, m *wfeProcInsModelImpl) {
	m.checkLockCap()
	var lc HandleLock
	m.Lock()
	lc, ok := m.lockMap[v]
	if ok {
		lc.Lock()
		lc.locked = true
		i++
		fmt.Println("exist i ========= ", i)
	} else {
		m.lockMap[v] = HandleLock{}
		lc = m.lockMap[v]
		lc.Lock()
		lc.locked = true
		i++
		fmt.Println("noexist i ========= ", i)
	}
	m.Unlock()
	fmt.Println("---> i", i)
	lc.Unlock()
	lc.locked = false
}

func url_test() {
	urls := "http%3A%2F%2F10.0.1.133%3A6666%2FAutoWorkflow%2FSubmitProcIns%3FparentTaskHandelUrl%3Dhttp%3A%2F%2F10.0.1.133%3A6666%2FAutoWorkflow%2FTaskHandle%26bussinessId%3D1061524390132338714"
	fmt.Println("url before:", urls)
	urls, _ = url.QueryUnescape(urls)
	fmt.Println("url after:", urls)
}
func golang_replace_all_test() {
	s := "echo sss ${year} hello world ${year}+COUNT-${month}+COUNT-${day}+COUNT ${hour}+COUNT:${min}+COUNT:${second}+COUNT sss"
	if strings.Contains(s, "${year}") {
		yearStr := replaceStr("${year}")
		s = strings.Replace(s, "${year}", yearStr, -1)
		fmt.Println(s)
	}
	if strings.Contains(s, "${month}") {
		monthStr := replaceStr("${month}")
		s = strings.Replace(s, "${month}", monthStr, -1)
		fmt.Println(s)
	}
	if strings.Contains(s, "${day}") {
		dayStr := replaceStr("${day}")
		s = strings.Replace(s, "${day}", dayStr, -1)
		fmt.Println(s)
	}
	if strings.Contains(s, "${hour}") {
		hourStr := replaceStr("${hour}")
		s = strings.Replace(s, "${hour}", hourStr, -1)
		fmt.Println(s)
	}
	if strings.Contains(s, "${min}") {
		minStr := replaceStr("${min}")
		s = strings.Replace(s, "${min}", minStr, -1)
		fmt.Println(s)
	}
	if strings.Contains(s, "${second}") {
		secondStr := replaceStr("${second}")
		s = strings.Replace(s, "${second}", secondStr, -1)
		fmt.Println(s)
	}
	fmt.Println(s)
}

func replaceStr(timeStr string) (result string) {
	count := 0
	t := time.Now()
	year, month, day := t.Date()
	hour := t.Hour()
	min := t.Minute()
	second := t.Second()
	preCount := 0
	zeroPrefix := ""
	switch timeStr {
	case "${year}":
		preCount = year + count
	case "${month}":
		preCount = m[month] + count
		if preCount < 10 {
			zeroPrefix = "0"
		}
	case "${day}":
		preCount = day + count
		if preCount < 10 {
			zeroPrefix = "0"
		}
	case "${hour}":
		preCount = hour + count
		if preCount < 10 {
			zeroPrefix = "0"
		}
	case "${min}":
		preCount = min + count
		if preCount < 10 {
			zeroPrefix = "0"
		}
	case "${second}":
		preCount = second + count
		if preCount < 10 {
			zeroPrefix = "0"
		}
	default:

	}
	result = zeroPrefix + strconv.Itoa(preCount)
	return result
}

func wild_code_test() string {
	s := "echo sss ${${year} hello world ${${year}+COUNT}-${${month}-COUNT}-${${day}/COUNT} ${${hour}abcCOUNT}:${${min}+COUNT}:${${second}+COUNT} sss"
	//s := "hello world {${year}+1}"
	if !strings.Contains(s, "${${") {
		fmt.Println("s: ", s)
		return s
	}
	result := ""
	for {
		if len(s) <= 0 {
			break
		}
		postfix := ""
		index := strings.LastIndex(s, "${${")
		if index == -1 {
			return s + result
		}
		handleStr := s[index:len(s)] //${${second}+COUNT}
		postfixIndex := strings.LastIndex(handleStr, "}")
		if postfixIndex == -1 {
			return s + result
		}
		if postfixIndex != -1 && postfixIndex != len(handleStr)-1 {
			postfix = handleStr[postfixIndex+1 : len(handleStr)]
			handleStr = handleStr[:postfixIndex+1]
		}
		timestr := calculateTime(handleStr)
		result = timestr + postfix + result
		if index == 0 {
			s = ""
		} else if index == -1 {
			break
		} else {
			s = s[0:index]
		}
	}
	fmt.Println(s + result)
	return s + result
}

func calculateTime(handleStr string) string { // ${${month}+COUNT}
	if !isValidTimeString(handleStr) {
		return handleStr
	}
	handleStr = handleStr[2 : len(handleStr)-1] //  ${month}+COUNT
	rightIndex := strings.LastIndex(handleStr, "}")
	if rightIndex == -1 {
		return "${" + handleStr + "}"
	}
	timeWild := handleStr[0 : rightIndex+1]
	COUNT := handleStr[rightIndex+1 : len(handleStr)]
	timeStr := cgStrToTime(timeWild, COUNT)
	return timeStr
}

func cgStrToTime(timeStr, COUNT string) (result string) {
	if !isValidCOUNT(COUNT) {
		return "${" + timeStr + COUNT + "}"
	}
	fmt.Println("timeStr:", timeStr, "COUNT:", COUNT)
	t := time.Now()
	year, month, day := t.Date()
	hour := t.Hour()
	min := t.Minute()
	second := t.Second()
	preCount := 0
	zeroPrefix := ""
	switch timeStr {
	case "${year}":
		preCount = calculateCOUNT(year, COUNT)
	case "${month}":
		preCount = calculateCOUNT(m[month], COUNT)
	case "${day}":
		preCount = calculateCOUNT(day, COUNT)
	case "${hour}":
		preCount = calculateCOUNT(hour, COUNT)
	case "${min}":
		preCount = calculateCOUNT(min, COUNT)
	case "${second}":
		preCount = calculateCOUNT(second, COUNT)
	default:
		return "{" + COUNT + "}"

	}
	if preCount < 10 {
		zeroPrefix = "0"
	}
	result = zeroPrefix + strconv.Itoa(preCount)
	return result
}

func isValidCOUNT(COUNT string) bool {
	if len(COUNT) == 0 {
		return true
	}
	if COUNT[:1] != "+" && COUNT[:1] != "-" && COUNT[:1] != "*" && COUNT[:1] != "/" {
		return false
	}
	return true
}

func calculateCOUNT(preCount int, COUNT string) int {
	result := preCount
	if len(COUNT) == 0 {
		return result
	}
	expression := COUNT[1:]
	cal := calculateExpression(expression)
	switch COUNT[:1] {
	case "+":
		result = result + cal
	case "-":
		result = result - cal
	case "*":
		result = result * cal
	case "/":
		if cal == 0 {
			return result
		}
		result = result / cal
	}
	return result
}

func calculateExpression(expression string) int {
	result := 0
	return result
}

func isValidTimeString(timeStr string) bool {
	if !strings.Contains(timeStr, "${year}") && !strings.Contains(timeStr, "${month}") &&
		!strings.Contains(timeStr, "${day}") && !strings.Contains(timeStr, "${hour}") &&
		!strings.Contains(timeStr, "${min}") && !strings.Contains(timeStr, "${second}") {
		return false
	}
	return true
}

func restraint(preCount int, timeType string) int {
	switch timeType {
	case "Month":
		if preCount < 1 {
			preCount = 1
		}
		if preCount > 12 {
			preCount = 12
		}
	}

	return 1
}

func time2_test() {
	t := time.Now()
	y, m, d := t.Date()

	today := time.Now().Format("2006-01-02")
	datetime := time.Now().Format("20060102150405") //后面的参数是固定的 否则将无法正常输出

	fmt.Println("time is : ", t)
	fmt.Println("y m d is : ", y, m, d)
	fmt.Println("now is :", today)
	fmt.Println("now is :", datetime)

}

//func tunny_goroutine_test() {
//	numCPUs := runtime.NumCPU()
//	pool := tunny.NewFunc(numCPUs, func(payload interface{}) interface{} {
//		var result []byte
//
//		// TODO: Something CPU heavy with payload
//		if value,ok := payload.([]uint8);ok {  // interface{} 需要断言才能转换
//			fmt.Println(string(value))
//		}
//		return result
//	})
//	defer pool.Close()
//
//	http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
//		input, err := ioutil.ReadAll(r.Body)
//		if err != nil {
//			http.Error(w, "Internal error", http.StatusInternalServerError)
//		}
//		defer r.Body.Close()
//
//		result := pool.Process(input)
//
//		w.Write(result.([]byte))
//	})
//
//	http.ListenAndServe(":8080", nil)
//}

func new_object_test() {
	a := new(int64)
	fmt.Println(a)
}

func bit_calculate_test() {
	var in int64
	tye := reflect.TypeOf(in)
	fmt.Println("输出类型:", tye.Kind())

	bitResult := 1 << 7
	fmt.Println("输出位计算结果", bitResult)

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
	fmt.Printf("a's address %p\n", &a[0])
	fmt.Printf("b's address %p\n", &b[0]) //a[0]和b[0]指向的地址差了一个元素的大小，说明reslice后指针直接指向a[low]的地址
	b = append(b, 2, 3, 4, 5, 6, 7)       // 元素个数超过cap时会重新分配地址，地址变了
	fmt.Printf("b's address after append %p\n", &b[0])
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
	b := a // 如果改为 b := &a 那么三四行输出的value就一样了
	//值传递，输出的地址不一样
	fmt.Println("a's address:", &a, " a's value:", a)
	fmt.Println("b's address:", &b, " b's value:", b)

	change_value(&b)
	fmt.Println("a's address:", &a, " a's value:", a)
	fmt.Println("b's address:", &b, " b's value:", b)

}
func change_value(b *int) {
	*b = 3
}
func string_test() {
	s := "hello world"
	s = "hello world2"
	//reflect.ValueOf(&s).Pointer()  // reflect.ValueOf（）参数一定要是指针类型，否则会panic
	fmt.Println(s)

	b := []byte{1, 2, 3}
	b[1] = 3
	fmt.Println(b)

	str := "welcome to outofmemory.cn"
	fmt.Println("str[", str, "] 替换前地址：v%", &str)
	str = strings.Replace(str, " ", ",", -1)
	fmt.Println("str[", str, "] 替换后地址：v%", &str)
	//输出地址是一样的，？？？？ 这是因为str原先在内存有一个地址，这个是不变的，string改变后，变的是str指向的地址
}

func panic_test() {
	//panic_without_recover()
	panic_with_recover()
}

func panic_with_recover() {
	defer func() { //recover()用于将panic的信息捕捉。recover必须定义在panic之前的defer语句中。
		if err := recover(); err != nil { //注意这里用分号隔开，为什么不是 && ?
			fmt.Println("catch the panic:", err)
		}
		fmt.Println("execute 3")
	}()
	fmt.Println("execute 1")
	panic("code panic")
	fmt.Println("execute 2") //即使recover捕获了异常，这一句也不会执行
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
	fmt.Println("a地址:", &a)   // #1 输出指针自身的地址
	fmt.Println("a指向的地址:", a) // #2  输出指针变量指向的地址
	invoke_point(a)

}
func invoke_point(b *int) {
	fmt.Println("a参数地址:", &b)   // #3  输出和#1不一样，因为是值传递，复制了一个地址
	fmt.Println("a参数指向的地址:", b) //#4  输出和#2一样，因为指针指向的地址是一样的
}

func map_test() {
	var m = map[string]string{"name": "zhangsan", "pass": "zhangsan'pass"}
	value, ok := m["name"] //map取值可以返回两个值
	fmt.Println(value, ok)
	value, ok = m["name don't exist"]
	fmt.Println(value, ok)
	fmt.Println(&m)
	fmt.Println(1 << 0) // 1
}

func var_test() {
	var a = 3
	const b = 4
	fmt.Println(&a, a)
	//fmt.Println(&b,b)  //不能引用constant的地址，因为常量在编译预处理阶段直接展开，作为指令数据使用，没有地址
	arr := make([]int, 1)
	fmt.Println(reflect.TypeOf(arr))
}

func unsafe_pointer_uintptr_test() {
	a := [4]int64{5, 1, 2, 3}
	fmt.Println(unsafe.Sizeof(&a[0]))                                                                       // 每个元素占8个字节，64位
	fmt.Println(unsafe.Pointer(&a[0]), unsafe.Pointer(&a[1]), unsafe.Pointer(&a[2]), unsafe.Pointer(&a[3])) //输出每个元素的内存地址 16进制
	p := uintptr(unsafe.Pointer(&a[0]))
	fmt.Println(p) //输出a[0]的内存地址 10进制
	p1 := unsafe.Pointer(p)
	fmt.Println(*(*int64)(p1)) // 根据p1的地址输出对应的值 （* int64)表示先转化为 int64 指针， * 表示取值
}

func wait_group_test() {
	var waitGroup sync.WaitGroup   //定义一个同步等待的组
	waitGroup.Add(1)               //添加一个计数
	go game.ConnSocket(&waitGroup) //调用其他包的方法执行任务
	waitGroup.Wait()               //阻塞直到所有任务完成
	fmt.Println("main DONE!!!")
}

func close_channel_test() {
	ch := make(chan string, 2)
	go func(in chan<- string) {
		in <- "hello"
		//time.Sleep(time.Second * 7) //这里如果睡眠7秒再写入数据就会报错，因为这个时候channel已经被关闭了
		in <- "world"
	}(ch)

	go func(out <-chan string) {
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

	go func(in chan<- string) {
		in <- "hello"
	}(process1)

	go func(in chan<- string) {
		in <- "world"
	}(process2)

	time.Sleep(time.Second * 2) //如果没有这一句，会输出两个default operation,因为两个协程来不及往通道放数据
	for i := 0; i < 2; i++ {    //如果不加for循环，select找到一个匹配的条件就退出了
		select {
		case str := <-process1:
			fmt.Println("process1 data", str)
		case str := <-process2:
			fmt.Println("process2 data", str)
		default:
			fmt.Println("default operation")
		}
	}

}

func single_channel_test() {
	ch := make(chan string)

	go func(out chan<- string) {
		out <- "hello"
	}(ch)

	go func(in <-chan string) {
		fmt.Println(<-in)
		//fmt.Println( <- in)
		fmt.Println("execute here") //多加一个fmt.Println( <- in) 这里就阻塞不会执行，一直到main运行结束
	}(ch)
	time.Sleep(2 * time.Second)
}

func channel_test() {
	ch := make(chan string, 2)
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
