//当前程序的包名
package main

//导入的包
import (
	"fmt"
	//"os"
	//"time"
	//"math"
	//"sort"
)

//显式类型定义,全局
const number int = 30
const (
	userName 	= "admin"
	address string  = "中关村"
	age      	= "30"
)

//全局变量
var address string = "北京中关村"
var (
	userType      = 1
	userFromTitle = "未知来源"
)

//一般类型全局变量
type names int8
type (
	newType int
	type1   float32
	type2   string
	type3   byte
)

//结构体声明
type comms struct {
	pak string
	num int
}

type persons struct {
	comms    //嵌入结构体
	userName string
	age      int
}

//接口声明
type golangs interface{}

//构造函数,init 函数也经常被用在当一个程序开始之前调用后台执行的 goroutine，如下面这个例子当中的 backend()：
func init() {
	//fmt.Println("我是构造函数")
}

func main() {

	//调用函数
	//echoShow()
	//foreachs()
	//sliceDemo()
	//mapDemo()
	//functions()
	//structDemo()
	//time.Sleep(1 * time.Second) //暂停2秒

	//并发
	c := make(chan bool)
	go func() {
		fmt.Println("A")
		c <- true
	}()
	<-c
}

//输出
func echoShow() {

	//获取int类型的长度，需要导入math包
	//fmt.Println(math.MinInt8)
	//fmt.Println(math.MaxInt32)

	//获取a的内存地址
	//fmt.Println(p)
	//获取a的值
	//fmt.Println(*p)

	//日期输出
	//t := time.Now()
	//fmt.Printf("日期：%4d-%02d-%02d %02d:%02d:%02d\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	//fmt.Printf("A") //输出A
	//fmt.Printf('A') //输出67

	// 指针-----------------------/
	//打印il的内存地址
	/*
	   var i1 = 5
	   fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)
	   var intP *int //定义一个指针
	   intP = &i1    //指针赋值
	   fmt.Printf("The value at memory location %p is %d\n", intP, *intP)

	   s := "good bye"
	   var p *string = &s
	   *p = "ciao"
	   print("skdfjslkd");
	   fmt.Printf("Here is the pointer p: %p\n", p)  // prints address
	   fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	   fmt.Printf("Here is the string s: %s\n", s)   // prints same string
	*/

	/*var k1 int = 8
	  var k2 int = 9
	  fmt.Printf("k1: %2d ,k2: %2d \n", k1, k2)
	  var uname string = "这是一个demo字符串\n"
	  fmt.Printf("string %s", uname)*/

	//%d 用于格式化整数
	//（%x 和 %X 用于格式化 16 进制表示的数字
	//%g 用于格式化浮点型（%f 输出浮点数，%e 输出科学计数表示法）
	//%0d 用于规定输出定长的整数，其中开头的数字 0 是必须的
	//%n.mg 用于表示数字 n 并精确到小数点后 m 位，除了使用 g 之外，还可以使用 e 或者 f，例如：使用格式化字符串 %5.2e 来输出 3.4 的结果为 3.40e+00
	/*fmt.Printf("32 bit int is: %d\n", n)
	  fmt.Println("字符串：", "我是string"+"连接符")
	  fmt.Println(os.Getenv("HOME"))
	  fmt.Println("Hello, 世界\n")
	  fmt.Println("这是一个字符串888")*/

	/*
	   var userName string = "liguibing"
	   var age int = 10
	   var address = "北京海定中关村"
	   number := "编码"

	   fmt.Printf("userName:%s", userName)
	   fmt.Printf("age:%d\n", age)
	   fmt.Printf("address:%s\n", address)
	   fmt.Printf("未知参数类型：%s\n", number)
	   //print(number)
	   fmt.Println("变量：", userName, age, address, number)
	*/
}

//常量
func constDemo() {
	//隐式类型定义
	const age = 15

	//显式类型定义
	const number int = 30
	const bb = "abc"
	const userName string = "liguibing"

	//数字型的常量是没有大小和符号的，并且可以使用任何精度而不会导致溢出：
	const Ln2 = 0.693147180559945309417232121458176568075500134360255254120680009

	//常量赋值
	const beef, two, c = "eat", 2, "veg"
	const Monday, Tuesday, Wednesday, Thursday, Friday, Saturday = 1, 2, 3, 4, 5, 6
	const (
		Unknown = 0
		Female  = 1
		Male    = 2
		Address = "北京市海淀区中关村"
	)
}

//基本类型和运算符
func numbers() {
	//类型转换
	//var a float32 = 100.1
	//将a转换为整型
	//b := int(a) //将输出100

	//整型转换为字符串,需要利用 strconv 包
	//var a int = 100
	//b := strconv.Itoa(a)

	//将字符串->整数
	//var c string = "100"
	//d := strconv.Atoi(c)

	/*
	   //整数：
	   int8（-128 -> 127）
	   int16（-32768 -> 32767）
	   int32（-2,147,483,648 -> 2,147,483,647）
	   int64（-9,223,372,036,854,775,808 -> 9,223,372,036,854,775,807）
	   //无符号整数：
	   uint8（0 -> 255）
	   uint16（0 -> 65,535）
	   uint32（0 -> 4,294,967,295）
	   uint64（0 -> 18,446,744,073,709,551,615）
	   //浮点型（IEEE-754 标准）：
	   float32（+- 1e-45 -> +- 3.4 * 1e38）
	   float64（+- 5 * 1e-324 -> 107 * 1e308）
	   //赋值
	   var b2 int32
	   var n int16 = 34
	*/
}

//变量
func variate() {
	/*
		       var a,b,c int = 1,3,4
		       a,b,c := 1,2,3

			   var x, y *int
			   var xxx, yyy, zzz int
			   var address string = "北京中关村"
			   var num int
			   var aa bool = false
			   var (
			       aaaa int
			       bbbb bool
			       strs string = "这是变量"
			   )
	*/
}

//数组
func array() {
	//数组
	//a := [3]int{1, 2} //长度不足3个元素时后面用0不足
	//fmt.Println(a)

	//b := [20]int{19: 8} //将8放到最后一位
	//fmt.Println(b)

	//c := [...]int{1, 2, 34, 5} //已经知道数组长度
	//fmt.Println(c)
	//fmt.Println(c[3]) //获取某一个元素

	//d := [...]int{999: 88} //第100个元素为88
	//fmt.Println(d)

	//数组指针
	//e := [...]int{99: 1}
	//var p *[100]int = &e
	//fmt.Println(p)

	//二维数组,下面是 包含了两个元素，每个元素的长度为3个
	/*f := [2][3]int{
	      {1, 2, 3},
	      {4, 5, 6}}
	  fmt.Println(f)
	  fmt.Println(f[1][2])
	*/

	//字符串数组，用法与整数一样
	g := [3]string{"AAA", "BBB", "CCC"}
	fmt.Println(g)

	//s3 := []string{"a", "A"}
	//fmt.Println(s3)

	//s4 := []string{'a','b','c'}
	//fmt.Println(s4)

	//s5 := [...]string{'a','b','c'}
	//fmt.Println(s5)

	//冒泡排序
	arr := [...]int{21, 36, 76, 2, 5, 9, 435}
	num := len(arr)
	fmt.Println(arr)

	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if arr[j] > arr[i] {
				tmp := arr[j]
				arr[j] = arr[i]
				arr[i] = tmp
			}
		}
	}
	fmt.Println(arr)
}

//切片 slice
func sliceDemo() {
	//定义一个切片
	//var a []int
	//fmt.Println(a)

	/*b := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//fmt.Println(b) //[1 2 3 4 5 6 7 8 9 10]
	//取区间的数据
	c := b[5:10]
	fmt.Println(c) //[6 7 8 9 10]
	d := b[3:]
	fmt.Println(d) //[4 5 6 7 8 9 10]
	e := b[:7]
	fmt.Println(e) //[1 2 3 4 5 6 7]
	*/

	//类型，包含的元素，存储容量 ,但容量不够时，容量会自动翻倍增加,容量可以不设置，但为了效率，可根据情况设置
	//s1 := make([]int, 3, 10) //slice的整型数组，3个元素，slice的初始容量:长度为10
	//分别打印：长度，容量，值
	//fmt.Println(len(s1), cap(s1), s1) //3 10 [0 0 0]

	//s2 := []byte{'a', 'b', 'c', 'd', 'e', 'f', '/', 'A', 'B'}
	//fmt.Println(s2) //这里打印的是acsll码 ，[97 98 99 100 101 102]
	//fmt.Println(s2[5:])
	//fmt.Println(string(s2)) //转换为字符串，abcdef/
	//s3 := []string{"a", "A"}
	//fmt.Println(s3)

	//切片增加元素
	//s4 := make([]int, 3, 6)
	//fmt.Println(len(s4), cap(s4)) // 3,6
	//s4切片增加元素
	//s4 = append(s4, 43, 435, 634, 53535)
	//fmt.Println(len(s4), cap(s4), s4) //7 12 [0 0 0 43 435 634 53535]

}

//map 的使用
func mapDemo() {
	//map 01
	//var a map[int]string
	// a = map[int]string{} 或者 a = make(map[int]string)

	// map 02
	// a := make(map[int]string)

	//map 03,键未整数，值为字符串
	//a := map[int]string{}
	// 键、值都为字符串
	//a1 := map[string]string{}

	//赋值
	/*a[1] = "ok"
	a[2] = "AAA"
	m := a[1]
	fmt.Println(a) // map[1:ok]
	fmt.Println(m) // ok
	//删除指定元素
	delete(a, 2)
	fmt.Println(a)*/

	/*
		sm := make([]map[int]string, 5)
		//fmt.Println(sm) //[map[] map[] map[] map[] map[]]
		//初始化map ,key=>value
		for k, _ := range sm {
			sm[k] = make(map[int]string, 1)
			sm[k][1] = "测试赋值"
			fmt.Println(sm[k])
		}
		fmt.Println(sm)*/
	/*
		sm2 := make([]map[int]string, 5)
		for k, v := range sm2 {
			//sm2[k] = make(map[int]string, 1)
			//sm2[k][1] = "测试赋值"
			fmt.Println(k)
			fmt.Println(v)
		}
		fmt.Println(sm2)
	*/

	/*
		sm3 := make([]map[int]string, 5)
		for i := range sm3 {
			sm3[i] = make(map[int]string, 1)
			sm3[i][1] = "测试赋值"
		}
		fmt.Println(sm3) //[map[1:测试赋值] map[1:测试赋值] map[1:测试赋值] map[1:测试赋值] map[1:测试赋值]]
	*/

	//间接排序(按key排序),map是无序性,需要导入soft包
	/*
		m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
		s := make([]int, len(m))
		i := 0
		for k, _ := range m {
			s[i] = k
			i++
		}
		sort.Ints(s)
		fmt.Println(s)
	*/

	/*
		m1 := map[int]string{1: "a", 2: "b", 3: "c"}
		//将m1转换为如下：
		//newret := map[string]int{"a": 1, "b": 2, "c": 3}
		m2 := make(map[string]int)
		for k, v := range m1 {
			m2[v] = k
		}
		fmt.Println(m1)
		fmt.Println(m2)
	*/
}

//循环
func foreachs() {
	//普通循环
	for i := 0; i < 5; i++ {
		//跳出循环
		if i > 3 {
			break
		}
		//跳出本次循环
		if i == 2 {
			continue
		}
		fmt.Printf("This is the %d iteration\n", i)
	}

}

//if判断
func ifDemo() {
	var isbool = 9
	if isbool == 1 {
		print("1")
	} else if isbool == 2 {
		print("2")
	} else if isbool > 3 && isbool < 5 {
		print("&&条件")
	} else if isbool == 6 || isbool == 7 {
		print("或条件")
	} else {
		print("AAA")
	}

	//初始化赋值，a=1，b=2
	if a, b := 1, 2; a > 1 {
		fmt.Println("AAA")
	} else if b > 3 {
		fmt.Println("BBB")
	}
	// ps 这里的 a，b赋值只能在if判断语句中使用，无法在if之外使用，switch也是如此

}

//switch分之判断
func switchDemo() {
	//demo 1
	var snum int = 2
	switch snum {
	case 1:
		print("1111")
	case 2:
		print("2222")
	default:
		print("未知")
	}

	//demo2
	var num int = 5
	switch {
	case num == 1:
		print("num==1")
	case num >= 3:
		print("num==3")
	default:
		print("未知")
	}
}

//Break、continue、lable、goto的使用
// lable 跳出代码块
// goto 调整执行位置，然后继续执行
func bclg() {
	/*
			//lable使用
		LABLE1:
			for {
				for i := 0; i < 10; i++ {
					if i > 3 {
						//这里的break会永远循环下去，因为外层是一个无限循环
						//break
						//这里会跳出两个循环，执行后面的代码
						break LABLE1
					}
				}
			}
	*/
	//goto使用
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				//这里的break会永远循环下去，因为外层是一个无限循环
				//break
				//这里会跳出两个循环，执行后面的代码
				goto LABLE1
			}
		}
	}
LABLE1:
}

//函数的使用
func functions() {
	//print("开始调用函数\n")
	//A(8, "test")
	deferDemo()

}

//带参数函数
func A(a int, b string) {
	fmt.Println(a)
	fmt.Println(b)
}

//带返回值,返回一个整数
func B(a int, b string) int {
	fmt.Println(a)
	fmt.Println(b)
	return 1
}

//返回多个
func C(a int, b int) (string, int) {
	fmt.Println(a)
	fmt.Println(b)
	return "A", 1
}

func D() (a, b, c int) {
	a, b, c = 1, 2, 3 //上面已经初始化，不需要 var或者：了
	return            //这里会自动返回a，b，c参数值，也可以写为：return a,b,c
}

//不确定参数,这里就可以传入多个参数
func E(a ...int) {
	print(a)
}
func F(a ...string) {
	print(a)
}
func G(a []int) {

}

//传递指针类型
func H(a *int) {
	*a = 2
	fmt.Println(*a)
}

//defer的使用，go没有try异常处理，
func deferDemo() {
	//panic("B")
	defer fmt.Println("A")
}

//结构体的使用
func structDemo() {
	/*a := persons{}
	//赋值
	a.userName = "张三"
	a.age = 30
	fmt.Println(a)
	*/

	//更简洁的赋值
	/*b := persons{
		userName: "admin",
		age:      30,
	}
	fmt.Println(b)*/

	//匿名结构体
	/*c := struct {
		number  int
		address string
	}{
		number:  10086,
		address: "中关村",
	}
	fmt.Println(c.address)
	*/

	//嵌入结构
	//初始化： 原始结构体名称:原始结构体名称{sex:30}
	//赋值：   当前结构体名称.sex = 13
	d := persons{}
	d.age = 3
	d.pak = "admin"
	d.num = 30
	fmt.Println(d.num)

}

//并发
func Gos() {
	fmt.Println("A")
}
