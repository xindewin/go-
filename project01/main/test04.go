package main

import (
	"fmt"
)

/*func main()  {
	var a int = 200
	if a < 200 {
		fmt.Println("too low")
	} else if a > 200 {
		fmt.Println("too high")
	} else if a==200 {
		fmt.Println("too suitable")

	}
}*/
//函数
func sum(x,y int) int{
	var c int
	c=x+y
	return c

}

//向函数传递数组
func getAversge(arr1 []int, size int) float32 {
	var i,sum int
	var result float32
	for i=0;i<size;i++{
		sum+=arr1[i]
	}
	result=float32(sum)/float32(size)
	return result

}

func main() {
	a:="abdc"
	var b int = 1000
	c:=2324
	switch b {
	case 1:
		fmt.Println("so ggod",c)
		break
	case 1000:
		fmt.Println("so well\t",a,sum(b,c))
	}

	//数组
	//var balance [10] int{1，2，3，4，5}
	//a1 := [...]int{12,123,13,13}
	str := [3]string{"a","b","c"}
	for i,s:=range str{
		fmt.Printf("第%d位是%q值\n",i,s)
	}

	//指针
	var d int
	fmt.Println(&d)

	var ip *int
	ip=&d

	fmt.Println(ip)
	fmt.Println(*ip)

	//空指针
	var ptr *int
	fmt.Println(ptr)   //nil

	//指针数组

	num1 := []int{1,2,3,4}
	var i int
	const MAX int = 4
	for i=0;i<MAX;i++ {
		fmt.Println(num1[i])
	}
	var num2 [MAX]*int

	for i=0;i<MAX;i++{
		num2[i]=&num1[i]
	}

	for i=0;i<MAX;i++ {
		fmt.Println(*num2[i])
	}

	fmt.Println(getAversge(num1,4))


	//二维数组（多维数组）
	var arr2 =[2][3]string{
		{"hh","aa","cc"},
		{"dd","ff","ee"}}
	for i:=0;i<2;i++{
		for j:=0;j<3;j++{
			fmt.Printf("arr2[%d][%d]=%q",i,j,arr2[i][j])
		}
	}



}




