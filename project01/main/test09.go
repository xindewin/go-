package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//map
	//使用make()分配内存
	scoreMap:= make(map[string]int,8)
	scoreMap["张三"]=90
	scoreMap["小明"]=100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n",scoreMap)
	//在声明时填充
	userInfo:=map[string]string{
		"username":"pprof.cn",
		"password":"xinxin",
	}
	fmt.Println(userInfo)
    //判断某个键是否存在
	v,ok:=scoreMap["小明"]
	if ok {
		fmt.Println(v)
		fmt.Println(ok)
	}else {
		fmt.Println("查无此人")
	}
	//map的遍历
	for k,v:=range scoreMap{
		fmt.Println(k,v)
	}
	for _,v:=range scoreMap{
		fmt.Println(v)
	}
	//使用delete()函数删除键值
	delete(scoreMap,"小明")
	for k,v:=range scoreMap{
		fmt.Println(k,v)
	}

	//按照指定顺序遍历map
	rand.Seed(time.Now().UnixNano())//初始随机化种子


}
