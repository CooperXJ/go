package main

//import "fmt"

import (
	"mygo/src/model"
	"fmt"
	"unsafe"
	"strconv"
	"time"
	"math/rand"
	"sort"
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

func changeArr(arr *[3]int){
	(*arr)[0] = 10;
	(*arr)[1] = 20;
	(*arr)[2] = 30;
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

	//数组  数组是值类型，他直接指向的是数据空间而不是地址
	var arr [3]int;
	fmt.Println(arr)
	fmt.Println(&arr)
	fmt.Println("arr[0]的地址为 ",&arr[0])
	fmt.Println("arr[1]的地址为 ",&arr[1])


	//数组遍历
	var arr1  = [...]int{1,2,3}
	fmt.Println(arr1)
	for _,v := range arr1{
		fmt.Printf("v=%v",v)
		fmt.Println()
	}

	fmt.Printf("%p\n",&arr1)
	fmt.Printf("%T\n",arr1)

	changeArr(&arr1)
	fmt.Println(arr1)

	//切片
	slice:=arr1[0:2]
	fmt.Println(slice)
	fmt.Printf("%p\n",&slice)//&slice是存储&arr1的地址
	fmt.Println(&arr1[0])
	fmt.Println(&slice[0])

	slice = arr1[:]
	fmt.Println(slice)
	slice = arr1[:3]
	fmt.Println(slice)
	slice = arr1[0:]
	fmt.Println(slice)

	var slice1 = make([]int,5,10)
	slice1[1] = 10
	slice1[2]  =20
	fmt.Println(slice1)

	var slice2 []int = []int{1,2,3}
	fmt.Printf("%T\n",slice2)
	fmt.Println(slice2)

	//切片追加
	slice3 := []int{2,4,5}
	slice2 = append(slice2,slice3...)
	fmt.Println(slice2)

	slice3 = append(slice3,100,200)
	fmt.Println(slice3)

	var slice4 = make([]int,10,10);
	slice5:=[]int{1,2,3,4,5}
	copy(slice4,slice5)
	fmt.Println(slice4)

	//string和切片
	strSlice := "Hello"
	//slice6:=strSlice[:]
	// slice6[0] = 'h'//这样是不可以的，因为此时的切片指向的是strSlice处的切片，因此不能够进行一个赋值操作，此处的赋值操作就相当于是对字符串进行赋值操作

	//如果需要对字符串进行修改的话
	//byte只能处理英文字母，不能处理带汉字的字符串
	slice7 := []byte(strSlice)
	slice7[0] = 'h'
	strSlice = string(slice7)
	fmt.Println(strSlice)
	//带有中文字符，需要按照字符处理
	slice8 := []rune(strSlice)
	slice8[0] = '薛'
	strSlice = string(slice8)
	fmt.Println(strSlice)


	//map
	var mapTest map[int]string
	mapTest = make(map[int]string,10)//必须要make才能为map分配空间，否则不会自动分配空间
	mapTest[0] = "Cooper"
	mapTest[1] = "Test"
	mapTest[2] = "Marry"

	fmt.Println(mapTest)

	mapTest2:=map[int]string{0:"shanghai",1:"hangzhou",2:"shenzhen"}
	fmt.Println(mapTest2)

	//删除map中的某个值
	delete(mapTest2,0)
	//因为go中没有clean清除所有map里面值的操作，因此直接将需要清空的变量make一下重新分配空间即可,原来的直接由gc回收即可
	mapTest = make(map[int]string)
	mapTest[0] = "Hello"
	fmt.Println(mapTest)

	//查找
	val,ok := mapTest2[1]
	if ok{
		fmt.Println("找到了该值 ",val)
	}else{
		fmt.Println("未找到该值")
	}

	//map切片  需要注意一下，需要make两次，因为切片需要make一次，然后map也需要make一次
	var newMap []map[int]string
	newMap = make([]map[int]string,2)

	if newMap[0]==nil{
		newMap[0] = make(map[int]string,2)
		newMap[0][0] = "Hi"
		newMap[0][1] = "Cooper"
	}
	if newMap[1]==nil{
		newMap[1] = make(map[int]string,2)
		newMap[1][0] = "Hello"
		newMap[1][1] = "Mary"
	}
	
	//添加元素
	var newMapadd map[int]string
	newMapadd = make(map[int]string,2)
	newMapadd[0] = "How"
	newMapadd[1] = "are"
	newMap = append(newMap,newMapadd)
	fmt.Println(newMap)

	//map排序
	mapSort := map[int]string {
		10:"Beijing",
		1:"New York",
		6:"Canada",
		3:"Shanghai",
	}

	var keys []int
	for k := range mapSort{
		keys = append(keys,k)
	}

	sort.Ints(keys)

	for _,i := range keys{
		fmt.Println(mapSort[i])
	}

}