package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
)
	
func main(){
	file,_:=os.Open("test.txt")
	fmt.Println(file)

	defer file.Close()

	reader := bufio.NewReader(file)
	for{
		//此处需要注意一下，这里的读取到最后一个行的时候如果该行有内容的话err会直接接收到EOF符号，也就会导致读取不出来，因此需要卸载判断结尾之前
		str,err:= reader.ReadString('\n')
		fmt.Print(str)
		if err == io.EOF{
			fmt.Println()
			break
		}
	}

	fmt.Println("已经读取到文章结尾了！")
}