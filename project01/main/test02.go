package main

//import "fmt"
import (
	"fmt"
)

func main(){

	//var c byte = 123   //byte int8
	//var x uint8 = 1<<1 | 1<<5
	//int 默认int32/64
	//fmt.Println(x)
	//fmt.Println(c)
	//fmt.Printf 格式化输出
	//fmt.Printf("x的类型是%T\nx占用的字节数是%d\n",x,unsafe.Sizeof(x))
	//fmt.Printf("%08b\n",x)

	//float 浮点型 符号位+指数位+尾数位
	//float 默认
	var f1 float32 =1293.123123  //输出1293.1232  注意精度 损失
	var f2 float64=1293.123123123123123
	fmt.Println(f1)
	fmt.Println(f2)

	//字符类型
	//go的字符串是由字节构成
	var c1 byte = 'a'
	fmt.Printf("字符输出%c，对应码值%d\n",c1,c1)
	//fmt.Println("\n")
	//字符串统一使用utf-8，不会出现中文乱码
	var c2 string = "我爱中国！"
    var c3 = "flying"
	//c3[0] = 'a'  字符串的值不可改变
	c3="xinxinxinxin"
	fmt.Println(c2,"\n",c3)

	//多行输出 反引号
	var c4 =`
	sdasdadwdnw
	jffdkskjfwe
	asdjasdjkadw
	`

	fmt.Println(c4)

	//字符串拼接

}
