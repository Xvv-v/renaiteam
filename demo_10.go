package main

import "fmt"


func main()  {
	//匿名函数
	var a,b int
	fmt.Scanln(&a,&b)
	fmt.Println("请输入两个数：")
	f:= func(a,b int) (int){
		return a+b
	}(a,b)
	fmt.Println(f)
}
