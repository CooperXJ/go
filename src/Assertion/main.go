package main

import "fmt"

func findType(i interface{}) {
    switch x := i.(type) {
    case int:
        fmt.Println(x, "is int")
    case string:
        fmt.Println(x, "is string")
    case nil:
        fmt.Println(x, "is nil")
    default:
        fmt.Println(x, "not type matched")
    }
}

func main() {

	//断言判断类型  如果成功则断言成功，否则则失败，触发 panic
    var i interface{} = 10
    t1 := i.(int)
    fmt.Println(t1)


    fmt.Println("=====分隔线=====")

    t2,ok:= i.(string)
	fmt.Println(t2)
	fmt.Println(ok)

	var j interface{} // nil  也会断言失败
	test,isOk:= j.(interface{})
	fmt.Printf("%v,%v\n",test,isOk);

	findType(10) //float64
}