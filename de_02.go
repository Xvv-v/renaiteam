package main
import "fmt"
func main()  {
	var a,b int
	fmt.Printf("请输入两个数：");
	fmt.Scanf("%d %d",&a,&b);
	swap(&a,&b);
	fmt.Printf("交换后的a=%d,b=%d",a,b);
	
}
func swap(a *int,b *int)  {
	var temp int
	fmt.Printf("a=%d\nb=%d",a,b);
	temp = *a;
	*a = *b;
	*b = temp;
	fmt.Printf("a=%d\nb=%d",a,b);
	fmt.Printf("a=%d\nb=%d",*a,*b);
}