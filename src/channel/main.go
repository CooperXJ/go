package main

import(
	"fmt"
	"time"
)

type Cat struct{
	Name string
	age int
}

func test(){
	for i:=0;i<10;i++{
		fmt.Printf("Hello World!~\n")
	}
}

func testError(){
	//解决协程中出现的panic现象，也就是可以防止该协程的错误影响到其他协程
	defer func(){
		if err:=recover();err!=nil{
			fmt.Printf("testError() 发生错误\n")
		}
	}()

	//此处的map没有make因此是一个错误
	var myMap map[int]string
	myMap[0] = "Hello"
}

func main(){
	var intchan chan int
	intchan = make(chan int,3)

	fmt.Printf("intchan的值为%v intchan的地址为%p\n",intchan,&intchan)

	intchan<- 10
	num:=11
	intchan<-num
	intchan<-55

	fmt.Printf("intchan的值为%v intchan的地址为%v\n",len(intchan),cap(intchan))

	num1:=<-intchan
	fmt.Println(num1)

	num2:=<-intchan
	fmt.Println(num2)

	num3:=<-intchan
	fmt.Println(num3)

	// num4:=<-intchan
	// fmt.Println(num4) //没有使用协程的情况下channel中的的元素取完了的话再去元素就会报错

	//定义一个存放任意类型的管道
	//因此结构体都继承了空接口
	allChan:=make(chan interface{},3)
	allChan<-"helo"
	allChan<-123
	allChan<-Cat{"小花猫",12}

	//退出“hello”和123
	<-allChan
	<-allChan

	newCat:=<-allChan
	fmt.Printf("%T,%v\n",newCat,newCat)
	//虽然可以答应出来cat类型，但是不能够直接访问newCat.Name 因为此时接受到的还是一个空接口类型，需要进行一个断言
	cat:=newCat.(Cat)
	fmt.Printf("%v,%v\n",cat.Name,cat.age);

	//当管道关闭的时候就不能再往里面存放数据了，但是依旧可以读取管道中的数据
	//如果未关闭管道就遍历其中的数据就会使得最终报错，也就是遍历到最后没有元素取了 所以在遍历的时候需要关闭管道
	intchan1:=make(chan int,10)
	intchan1<-1
	intchan1<-2


	//遍历之前必须要关闭管道
	close(intchan1)

	for v := range intchan1{
		fmt.Println(v)
	}


	//如果说有两个管道，只有一个在写并且写的元素个数大于管道的capacity，但是没有往外取的协程，那么会发生死锁
	//如果一个在读，一个在写，那么无论两者速度如果，都不会发生死锁

		 
	// 自动关闭管道 因为有的时候我们无法确定什么时候关闭管道，因此可以使用select来自动关闭管道  不能先关闭了管道再使用这句话
	// for{
	// 	select{
	// 		case v:=<-intchan :
	// 			fmt.Printf("还有值%v，别关我\n",v)
	// 		case v:=<-intchan1 :
	// 			fmt.Printf("----还有值%v，别关我\n",v)
	// 		default:
	// 			fmt.Printf("都没有值了，我全部关闭了")
	// 			return
	// 	}
	// }
	
	go test()
	go testError()

	time.Sleep(time.Second*5)
}
