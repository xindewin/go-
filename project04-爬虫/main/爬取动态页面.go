package main

import (
	"fmt"
	"strconv"
)

//获取页面
func GetPage(url string,num1 int,num2 int) {
	for i:=num1;i<=num2;i++{
		fmt.Printf("正在爬取第%d页",i)
		url=url+strconv.Itoa(i)
		PrintPage(url)

	}
}
//打印页面
func PrintPage(url string) {
	bytes:=GetBytes(url)
	str:=string(bytes)
	findByEle(str)
}



func main() {
	var start int
	var end int
	fmt.Println("输入爬取初始页面")
	fmt.Scan(&start)
	fmt.Println("输入截止爬取页面")
	fmt.Scan(&end)
	GetPage("https://www.zhipin.com/c100010000/?page=",start,end)
}