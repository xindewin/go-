package main

import "fmt"

func main() {
	//var slice1 []int=make([]int,len)//利用make创建切片
	//数组与切片
	arr1:=[3] int{1,2,3}
	fmt.Println(arr1)
	var sli []int

	if sli==nil{
		fmt.Println("sli和nil相等")
	}else {
		fmt.Println("sli和nil不相等")
	}
	//数组
	//array:=[...]int{4:1}
	//切片
	//slice:=[]int{4:1}
	//nil切片
	/*var slice1 []int
	fmt.Printf("nil切片%p\n",&slice1)
	//空切片
	slice2:=[]int{}
	//数组
	var array [3]int
	fmt.Printf("空切片%p\n",&slice2)
	fmt.Println("over")
	fmt.Printf("没有初始化的数组%p",&array)
    */




    //初始化切片
	slice4:=make([]int,10,20)
	slice:=[]int{1,2,3,4,5}
	slice1:=slice[:]
	slice2:=slice[0:]
	slice3:=slice[:5]
	newslice:=slice[1:3]//左开右闭
	newslice[0]=10

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(newslice)
	fmt.Println(slice4)

	//append

}
