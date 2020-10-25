package main

import (
	"fmt"
	"os"
	"io"
	"io/ioutil"
	"bufio"
)

func copyFile(srcFile string,destFile string) (written int64,err error) {
	file1,_:=os.OpenFile(srcFile,os.O_RDONLY,1)
	reader:=bufio.NewReader(file1)
	defer file1.Close()

	file2,_:=os.OpenFile(destFile,os.O_WRONLY|os.O_CREATE,777)
	writer:=bufio.NewWriter(file2)
	defer file2.Close()

	return io.Copy(writer,reader)
}

func main(){
	file,error:=os.OpenFile("test.txt",os.O_RDWR|os.O_CREATE,1)
	if error!=nil {
		fmt.Println("打开文件失败")
	}

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

	//一次性读取整个文件，不适合大文件
	readerTotal,_ := ioutil.ReadFile("test.txt")
	fmt.Printf("%v",readerTotal)
	fmt.Printf("%v",string(readerTotal))
	fmt.Println()

	//写文件
	str := "say hi\n"
	writer := bufio.NewWriter(file)
	writer.WriteString(str)

	//刷新缓冲区  一定要记得
	writer.Flush()

	copyFile("test.txt","copy_file.txt")
}