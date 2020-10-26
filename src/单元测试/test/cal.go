package test

//AddNumber 是一个测试
func addNumber(num int) (res int){
	res = 0
	for i:=1;i<=num;i++{
		res+=i;
	}

	return res;
}

func subNumber(num1 int,num2 int)(res int){
	return num1-num2;
}