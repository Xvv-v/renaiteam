package main

import "fmt"


func main()  {
	const(
		a int = iota
		c
		b
	)
	//a在iota置0
	//每访问一个常量iota都会加1
	type Weekday int//Weekday类型声明

	const (
		Sunday Weekday = iota//0
		Monday//1
		Tuesday//2
		Wednesday//3
		Thursday//4
		Friday//5
		Saturday//6

	)

	fmt.Println(a,"   ",b,"	",c)
	fmt.Print(Sunday,"  ",Monday,"  ", Tuesday,"   ", Wednesday,"   ", Thursday,"   ",
		Friday,"   ",Saturday)

}
