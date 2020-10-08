package main

import (
	"strings"
	"fmt"
)


func makeSuffix(suffix string) func(string) string{
	return func(name string) string{
		if !strings.HasSuffix(name,suffix){
			return name+suffix
		}
			return name
	}
}

func main() {
	var f = makeSuffix(".avi")
	fmt.Println(f("bird.avi"))
	fmt.Println(f("girl.jpg"))
	fmt.Println(f("boy"))
}

