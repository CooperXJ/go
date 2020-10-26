package main

import(
	"fmt"
	"runtime"
	"sync"
	"time"
)

var(
	myMap = make(map[int]int,10)

	//锁
	lock sync.Mutex
)

func test(n int){
	res:=1
	for i:=1;i<=n;i++{
		res*=i
	}
	lock.Lock()
	myMap[n] = res;
	lock.Unlock()
}

func main(){
	//获取CPU个数
	cpuNum:=runtime.NumCPU()  
	fmt.Println(cpuNum)

	//设置使用CPU的个数
	runtime.GOMAXPROCS(cpuNum-1)

	for i :=0;i<200;i++{
		go test(i)
	}

	time.Sleep(time.Second*5)

	lock.Lock()
	fmt.Println(myMap)
	lock.Unlock()
}