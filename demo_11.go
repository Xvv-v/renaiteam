package main

import (
	"fmt"
	"go/ast"
)
type Command struct {
	Name    string    // 指令名称
	id     int      // 指令绑定的变量
}

func main()  {
	a := new(Command)
	b := &Command{}
	a.id=545
	a.Name="sfhiu"
	b.id=465
	b.Name="dsgfedr"
	c:=&Command{
		id:123,//初始化结构体键值对
	Name:"djhviuhi"}
	//匿名结构体
	ins :=&struct {
		id int
		name string
		TEl  int
	}{
		123,
		"jdghuis",
		13564}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(ins)//打印匿名结构体
}