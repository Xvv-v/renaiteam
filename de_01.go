package main

import "fmt"

func main()  {

	var a,b,c int
	var sum int = 0;
	for a = 1; a<=4; a++{
		for b = 1;b <= 4;b++{
			for c =1;c <= 4;c++{
				if a!=b&&b!=c&&a!=c{
					sum++;
				}
			}
		}
	}
	fmt.Printf("sum=%d",sum);
}
