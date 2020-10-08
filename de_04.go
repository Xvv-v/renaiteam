package main

import "fmt"


func main()  {
	var str string = "欢迎来到go的世界！"
	ptr :=&str

	fmt.Print("打印str的地址n",&str)
	fmt.Print("打印ptr的存储内容地址：\n",ptr)
	fmt.Print("打印ptr存储的地址里内容：\n",*ptr)
	*ptr = "欢迎退出！"
	fmt.Print("打印ptr存储的地址里内容：\n",*ptr)
	fmt.Print("打印str", str)

}

