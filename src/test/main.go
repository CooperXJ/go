package main

import (
	"errors"
	"strings"
	"fmt"
)


func makeSuffix(suffix string) func(string) string{
	return func(name string) string{
		if !strings.HasSuffix(name,suffix){
			return name+suffix
		}
			return name
	}
}

func testDefer(n1 int,n2 int){
	defer fmt.Println("这是一个defer语句 ",n1)
	defer fmt.Println("这是一个defer语句 ",n2)
	fmt.Println(n1+n2)
}

//异常处理
 func testError(){
	defer func() {
		err:=recover()
		if err != nil{
			fmt.Println("err = ",err)

		}
	}()
	num1 := 10
	num2 := 0
	fmt.Println(num1/num2)
 }

 func testDIYError(name string) (err error){
	if name =="config.ini"{
		return nil
	}else{
		return errors.New("读取文件错误")
	}
 }

 func testDIYError2(){
	err:=testDIYError("config.ini1")
	if err!=nil {
		panic(err)
	}
	fmt.Println("继续执行...")
 }


func main() {
	var f = makeSuffix(".avi")
	fmt.Println(f("bird.avi"))
	fmt.Println(f("girl.jpg"))
	fmt.Println(f("boy"))

	testDefer(1,2)

	num2 := new(int)
	fmt.Printf("num2的类型为%T  num2的值为%v num2的地址为%v  num2这个指针指向的值为%v\n",num2,num2,&num2,*num2)


	testError()

	fmt.Println("这是异常之后的内容...")

	testDIYError2()
}

