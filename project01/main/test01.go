package main //文件必须属于一个包，把test.go归属到main
import "fmt"  //引入一个包 fmt

//定义全局变量
var (
    n1=123
	n2=2313
	n3="chongyan"
)

func main()  {
	//输出hello goland!
	//fmt.Println("hell\"234234\"o goland!")
	//fmt.Println("hello\t go\rworld!")
	//fmt.Println("hello   world!distroy")
	//fmt.Println("hello \ngo!")
/*
    //变量使用
	//1.指定变量类型，若没有赋值，使用默认值
	var num int 
	fmt.Println("num =",num)

	//2.根据赋值推导
	var num2= 123.34
	fmt.Println("num2 =",num2)

	//利用name:=XXX
	num3:="string"
	fmt.Println("num3 =",num3)

	fmt.Println("n1",num,"n2",num2,"n3",num3)
*/

    //一次声明多个变量
	//var n1,n2,n3=123,'s',"qwddqwd"
	//fmt.Println("n1",n1,"n2",n2,"n3",n3)

	//n1,n2,n3:=123,'s',"qwddqwd"
	fmt.Println("n1",n1,"n2",n2,"n3",n3)

	var a []string=make([]string,0)
	a=append(a,"aaaa")
	a=append(a,"bbbb")

	fmt.Println(len(a))
fmt.Println(a)
	fmt.Printf("%v",a)
} 