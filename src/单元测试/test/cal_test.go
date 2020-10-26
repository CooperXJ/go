package test

import(
	"testing"
)

//TestAddNumber 测试
func TestAddNumber(t *testing.T){
	res := addNumber(10)
	if res!=55{
		t.Fatalf("执行错误！")
	}

	t.Logf("执行正确")
}

func TestSubNumber(t *testing.T){
	res:=subNumber(1,2)
	if res!=-2{
		t.Fatalf("执行错误")
	}

	t.Logf("执行成功！")
}