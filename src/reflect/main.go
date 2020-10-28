package main

import(
	"fmt"
	"reflect"
)

// Student 测试类
type Student struct{
	Name string `json:"name"`
	Age int	`json:"age"`
}


func (stu *Student) SetName(name string){
	stu.Name = name;
}

func (stu Student) Add(val1,val2 int)int{
	return val1+val2
}

func (stu Student) Show(){
	fmt.Printf("name = %v,age = %v\n",stu.Name,stu.Age)
}

func testRelfect(b interface{}){
	val1:=reflect.ValueOf(b)
	fmt.Printf("%v Type=%T\n",val1,val1) //&{Cooper 12} Type=reflect.Value

	val2:= reflect.TypeOf(b)
	fmt.Printf("%v Type=%T\n",val2,val2)//*main.Student Type=*reflect.rtype

	//先转换为接口类型，在对该接口进行类型断言使其成为正确的类型  如果是一般的类型只可以直接使用go提供的方法，例如对于int类型的可以直接使用val1.Int()
	val3:=val1.Interface()
	studentTemp,ok:=val3.(*Student)//此处传入的是一个指针
	fmt.Printf("%v\n",ok)
	fmt.Printf("name = %v age = %v\n",studentTemp.Name,studentTemp.Age)

	//Elem的作用相当于是将指针转化为指针指向的变量在对其进行修改（基于该对象本身的修改）
	val1.Elem().Field(0).SetString("Aaron");

}

func main(){
	student1:=Student{"Cooper",12}
	testRelfect(&student1)
	fmt.Println(student1)

	student1.SetName("Cooper")
	student1.Show()

	stuVal:=reflect.ValueOf(student1)
	stuType:=reflect.TypeOf(student1)

	//获取struct的中的Filed以及对于的Tag（也就是之前将字段转换为小写的json注解）
	for i:=0;i<stuVal.NumField();i++{
		fmt.Printf("%v ",stuVal.Field(i))
		fmt.Printf("%v\n",stuType.Field(i).Tag.Get("json"))
	}

	
	var params []reflect.Value
	params = append(params,reflect.ValueOf(10))
	params = append(params,reflect.ValueOf(20))
	fmt.Printf("Type = %v\n",stuVal.NumMethod())
	res:=stuVal.MethodByName("Add").Call(params) //方法名需要大写  否则reflect包无法访问这些方法
	fmt.Println(res[0].Int())
}