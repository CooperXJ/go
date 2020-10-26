package main

import(
	"fmt"
	_"math"
	"time"
)

func putNum(intchan chan int,count int){

	for i:=2;i<count;i++{
		intchan<-i
	}

	close(intchan)

	fmt.Println("数据存放完成！")
}

//判断素数
func primeNum(intchan chan int,primeChan chan int,exitChan chan bool){

	var flag bool

	for {
		//这个千万不能忘记添加了！！！！，每次都需要重置一下flag
		flag = true

		num,ok := <-intchan

		if !ok{
			fmt.Println("取不到数据了...")
			break
		}
		for i:=2;i<num;i++ {
			if num%i==0 { 
				flag = false;
				break;
			}
		}

		if flag{
			primeChan<-num
		}
	}

	fmt.Println("取不到数据了，我退出了...")
	exitChan<-true
}


func main(){
	intChan:=make(chan int,1000)
	primeChan:=make(chan int,2000)
	exitChan:=make(chan bool,4)

	startTime := time.Now().Unix()
	var endTime int64

	go putNum(intChan,200000)

	for i:=0;i<4;i++{
		go primeNum(intChan,primeChan,exitChan)
	}

	go func(){
		for i:=0;i<4;i++{
			<-exitChan
		}
		
		close(primeChan)
		close(exitChan)

		endTime = time.Now().Unix()
		fmt.Printf("time: %v\n",endTime-startTime)
	}()

	

	//只要管道不关闭就会一直阻塞
	for v:=range primeChan{
		fmt.Println(v)
	}

	//只要有人在取这个管道里面的东西就不会自动退出，会阻塞在这里
	// for {
	// 	res,ok:=<-primeChan
	// 	if !ok{
	// 		break;
	// 	}
	// 	fmt.Println(res)
	// }

}	