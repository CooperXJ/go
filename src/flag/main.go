package main

import(
	"fmt"
	"flag"
)

func main(){
	var user string
	var password string
	var port int
	var host string

	flag.StringVar(&user,"u","","用户名默认为空")
	flag.StringVar(&password,"p","","密码默认为空")
	flag.StringVar(&host,"h","localhost","IP地址为空")
	flag.IntVar(&port,"port",3306,"端口默认为3306")

	//必须调用该方法
	flag.Parse()

	fmt.Printf("user=%v,password=%v,host=%v,port=%v\n",user,password,host,port)

	//测试指令 ./test.exe -u Cooper -p 123  | ./test.exe -help
}