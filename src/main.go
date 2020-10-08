package main

//import "fmt"

import (
	"mygo/src/model"
	"fmt"
	"unsafe"
	"strconv"
	"time"
	"math/rand"
)

type myfunType = func(int)int //自定义数据类型

var(
	//funn 这是一个匿名函数
	funn = func(a int,b int) int{
		return a-b
	}
)

func fb(n int) int {
	if(n==1||n==2){
		return 1;
	}else{
		return fb(n-1)+fb(n-2)
	}
}

func getFb(funvar func(int)int,n int) int{
	return funvar(n)
}

func simpleOperation(a int,b int) (sum int,sub int){
	sum = a+b;
	sub = a-b;
	return
}

func addSum(t0 int,args... int) int {
	var sum = t0;
	var i = 0
	for i = 0;i<len(args);i++ {
		sum+=args[i];
	}

	return sum
}

//AddUpper 闭包
func AddUpper() func(int) int{
	var n = 10
	var str = "Hello"
	return func(x int) int{
		n = n+x;
		str+="a"
		fmt.Println("str= ",str)
		return n
	}
}

func main() {

	var i = 10
	i = 11

	var c byte = 255

	fmt.Println(i, c)

	//输出查看变量的类型
	fmt.Printf("i =  %T \n", i)

	fmt.Println()

	//输出查看变量占用的字节数
	fmt.Printf("i 的字节数为 %d \n",unsafe.Sizeof(i))

	//单个字符输出
	var c1  = 'A'//一般的字符例如ascii码表中的字符0-255范围内的字符可以直接使用byte进行存储
	var c2  int = '薛'//此处如果使用byte的话会导溢出
	fmt.Printf("%c %c \n",c1,c2)


	//字符串
	var str0 string = "Hello World"
	var str1 string = `Hello \n World`//使用 ``就可以直接转移字符串中的特殊字符，比如说\n，“”等
	var str2 string = "Hello World " +
	"my girl"  //连接较长的字符串需要将“+”放到末尾，这样系统自动不给末尾添加“;”
	fmt.Println(str0)
	fmt.Println(str1)
	fmt.Println(str2)


	//变量类型的转换：必须使用显示转换，无法自动转换
	var t1 float64 = 1
	var t2 int = int(t1)
	var t3 = 1.111111
	fmt.Printf("t3 = %v t2 = %d\n",t3,t2)//浮点型的输出需要使用%v

	var t4 int32 = 1;
	var t5 int64  = 2;
	var t6 = int64(t4)+t5;//两个不同类型的变量是不可以相加的，正如上面所述go不会帮我们进行自动进行类型转换
	//var t7 int8 = 128//直接就溢出了
	fmt.Printf("%d\n",t6)


	//变量转为字符串
	var t7 int64 = 12;
	var str string;
	str = fmt.Sprintf("%v",t7);
	fmt.Printf("%T  str = %v\n",str,str)

	str = strconv.FormatInt(t7,10);
	fmt.Printf("%T  str = %v\n",str,str)


	//字符串转为变量
	var t8 int64 
	t8,_= strconv.ParseInt(str,10,64)  //此处有两个返回值，我们只需要第一个，如果某个返回值不需要可以使用_来代替
	fmt.Println(t8)

	fmt.Println(model.HeroName)

	//i++ i-- 只允许单独使用 其他情况都不允许
	var t9 int = 1;
	t9++;
	// ++t9;//不允许前置++或者--
	println(t9)


	//交换a与b 不通过中间变量
	a:=10;
	b:=20;
	a = a+b;
	b = a-b;
	a = a-b;
	fmt.Println(a)
	fmt.Println(b)

	// 控制台输入  has some problem
	// var name string
	// var age int
	// fmt.Println("请输入姓名")
	// // fmt.Scanln(&name)
	// fmt.Println("请输入年龄")
	// fmt.Scanln(&age)
	// fmt.Scanf("%v %v",name,age)
	// fmt.Printf("name = %v age = %d\n",name,age)

	//if
	var age = 10
	if age<10{
		fmt.Println("小于10岁")
	}else{
		fmt.Println("大于10岁")
	}

	//switch
	switch age{
		case 18:
			fmt.Println("我是18岁")
			fallthrough //默认只能穿透一层
		case 19:
			fmt.Println("我是19岁")
		default:
			fmt.Println("我不知道多少岁")
	}

	//for
	for i = 0;i<len(str);i++ {
		fmt.Printf("%c",str[i])
	}
	fmt.Println()

	//随机数
	rand.Seed(time.Now().Unix())
	var num = rand.Intn(100)
	fmt.Println(num)

	//函数调用
	fmt.Println(fb(3))

	//函数也可以是参数，自身也是一种参数类型
	//var funvar = fb

	fmt.Println(getFb(fb,3))

	//支持函数返回值自定义名称 ，这样可以写返回值
	var sum,sub = simpleOperation(12,8)
	fmt.Printf("sum = %v,sub = %v\n",sum,sub)


	//可变参数  函数多个参数
	var res = addSum(1,2,3,4)
	fmt.Println(res)

	//匿名函数的调用方式
	//方式1
	res = func(a int,b int) int{
		return a-b
	}(1,2)
	fmt.Println(res)
	//方式2
	var fun = func(a int,b int) int{
		return a-b
	}
	fmt.Println(fun(1,2))
	//方式3 全局匿名函数
	fmt.Println(funn(1,2))


	//闭包 闭包是类，函数是操作，n为字段
	var f func(int)int = AddUpper()//返回一个func(int)int类型的变量
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))

	
}