package main

import "fmt"

func main(){
	//显示转换
	//var a int32 =23123 //会报错int32不能给int64
	var a int64=23123
	var b int64 =12312
	//var c int8//会报错，超出范围
	var c int64
	c=a+b
	fmt.Println(c)
	//go中存放的是字符的码值
	var str1 byte ='s'
	var str2 int ='北' //byte存放的字符ascii值有限，超出之后可放在int里
	fmt.Println(str1)
	fmt.Printf("%c\n",str1)//%c输出字符
	fmt.Printf("%c",str2)

	var num rune = 32412 //rune就类似int32
	fmt.Printf("num的数据类型是%T，num的值是%v",num,num)

}
