package model

import "fmt"
//HeroName 这是一个结构体访问的测试
var HeroName string = "宋江"

//Car this is factory method demo
type car struct {
	Name string
	age  int
}

//Student is a test
type Student struct {
	Name string
	score int
}


func (stu *Student) ShowInfo(){
	fmt.Printf("name = %v,score = %v\n",stu.Name,stu.score)
}

func (stu *Student) SetScore(score int) {
	stu.score = score
}

func (stu *Student) SetName(name string){
	stu.Name = name
}

//Pupil is a test
type Pupil struct{
	Student
}

//Graduate is a test
type Graduate struct{
	Student
}

//GetNewCar 获取一个Car对象
func GetNewCar(name string, age int) *car {
	return &car{
		Name: name,
		age:  age,
	}
}

//访问私有字段的方法
func (c *car) GetAge() int {
	return c.age
}


