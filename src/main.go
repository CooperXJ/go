package main

//import "fmt"

import (
	"fmt"
	"unsafe"
)

func main() {

	var i = 10
	i = 11

	var c byte = 255

	fmt.Println(i, c)

	//输出查看变量的类型
	fmt.Printf("i =  %T ", i)

	fmt.Println()

	//输出查看变量占用的字节数
	fmt.Printf("i 的字节数为 %d ",unsafe.Sizeof(i))

}
