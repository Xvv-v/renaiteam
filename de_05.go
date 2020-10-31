package main

import (
	"container/list"
	"fmt"
)

func main()  {
	//1) 通过 container/list 包的 New() 函数初始化 list
	list_name := list.New()
	//2) 通过 var 关键字声明初始化 list
	var list_name_01 list.List
	//在列表中插入元素
	list_name.PushBack("你好！")
	list_name_01.PushBack("世界！")
	e := list_name_01.PushBack("我喜欢你！")//添加句柄删除
	list_name_01.Remove(e)
	for i:= list_name_01.Front(); i != nil ;i = i.Next(){
		fmt.Print(i.Value)
	}


}