package main

import "fmt"

type dongwu interface {
	call(a string)
	run(a string)
}
type cat struct {

}

func (* cat) call(a string)  {
	fmt.Print(a)
}
type dog struct {
	cat
}

func (*dog)run(a string)  {
	fmt.Print(a)
}
func main()  {
	n :=new(dog)
	var m dongwu
	m = n
	m.call("喵喵叫......")
	m.run("跑的很快......")
}