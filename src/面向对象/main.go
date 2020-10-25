package main

import (
	"mygo/src/model"
	"encoding/json"
	"fmt"
)

//Cat 是一个例子
type Cat struct {
	name       string
	age        int
	mapAddress map[int]string
}

//Point 点
type Point struct {
	x int
	y int
}

//Rect 矩阵
type Rect struct {
	left, right *Point
}

//Person 序列化demo
type Person struct {
	Name string `json:"name"` //后面是tag 是为了方便json字符串将首字母大写改成小写 这里属性值不得不首字母大写  因为不首字母大写其他包就访问不了这个属性
	Age  int    `json:"age"`
}

//方法
func (p Person) test() {
	p.Name = "jack"
	fmt.Println("名字是: ", p.Name)
}

func (p *Person) testPoint() {
	p.Name = "Harry"
	fmt.Println("hello,this is a point test ", (*p).Name)
}

func (p *Person) String() string {
	str := fmt.Sprintf("name = [%v],age = [%v]", p.Name, p.Age)
	return str
}

type integer int

func (i *integer) printInt() {
	fmt.Println(*i)
}


//接口实例测试
type usb interface{
	start()
	stop()
}

type camera struct{
	
}

func (c *camera) start() {
	fmt.Println("相机打开...")
}

func (c *camera) stop() {
	fmt.Println("相机关闭...")
}


type phone struct{
	
}

func (c *phone) start() {
	fmt.Println("手机打开...")
}

func (c *phone) stop() {
	fmt.Println("手机关闭...")
}

type computer struct{

}

func (c *computer) work(usb usb) {
	usb.start()
	usb.stop()
}


func main() {
	var cat Cat
	cat.name = "Tom"
	cat.age = 1
	cat.mapAddress = make(map[int]string)
	cat.mapAddress[0] = "Home"
	cat.mapAddress[1] = "Wild"

	fmt.Println(cat)

	//复制过来的catCopy与本身cat没有联系
	catCopy := cat
	catCopy.age = 2
	fmt.Println(cat)
	fmt.Println(catCopy)

	//结构体初始化的几种方式
	//new方法产生的对象其实返回的是一个指针类型，按道理来说应该写成 （*cat1.name = "T"但是go底层做了优化，可以直接写成下面的方式
	var cat1 *Cat = new(Cat)
	cat1.name = "T"
	cat1.age = 20
	fmt.Println(cat1)

	var cat2 *Cat = &Cat{}
	cat2.name = "OO"

	//在结构体中所有的元素在内存中都是连续存在的，即使是这里的指针作为结构体内的元素指向他们的指针也是连续的，但是指针本身有可能并不是连续的，因为这是操作系统随机分配的
	var rect01 = Rect{&Point{1, 2}, &Point{3, 4}}
	fmt.Println(rect01)

	fmt.Println(*rect01.left, " ", *rect01.right)
	fmt.Println(&rect01.left, " ", &rect01.right) //结构体中的元素，也就是指向指针的指针的元素

	//结构体之间如果想要相互转换，必须结构体中的所有字段都相同（包括名字、类型、个数）

	//结构体序列化
	perosn := Person{"Cooper", 23}
	josnstr, _ := json.Marshal(perosn)
	fmt.Println(string(josnstr))

	//调用方法
	perosn.test()
	perosn.testPoint()
	(&perosn).testPoint()

	//不一定非要是结构体才可以绑定方法，普通的变量类型也可以绑定方法(但是必须要给变量取一个别名，这样的话go会认为这是一个用户自定义的类型)，例如
	var testInt integer = 10
	testInt.printInt()

	//结构体的toString方法
	//只要对结构体的String方法进行重写在fmt.Println(person)的时候就会默认调用编写的String方法
	perosn1 := Person{"Cooper", 22}
	fmt.Println(perosn1)
	fmt.Println(&perosn1) //此处一定要添加&，否则会和原来输出的一样

	//通过变量指针调用方法，依然是值传递
	(&perosn).test() //虽然这里使用指针调用的方法，但是原来的person依旧没有改变，不信看下面p.Name还是原来的  此处可以得出一个结论就是无论外面调用者的身份是什么（指针调用还是变量调用），始终唯一的结果都和方法的参数有关，参数为值那么就是值引用，参数为指针，那么就是地址引用
	fmt.Println(perosn)


	//工厂模式  不直接在包里面对struct首字母进行大写，而是选择调用方法返回指针
	mycar:=model.GetNewCar("Audio",1)
	fmt.Println((*mycar).GetAge())
	fmt.Println(*mycar)

	//继承
	pupil:=model.Pupil{}
	pupil.SetName("Cooper")
	pupil.SetScore(10)
	pupil.ShowInfo()

	graduate:=model.Graduate{}
	graduate.SetName("Aaron")
	graduate.SetScore(100)
	graduate.ShowInfo()


	//接口实例
	computer:=computer{}
	camera:=camera{}
	phone:=phone{}

	computer.work(&camera)
	computer.work(&phone)
}
