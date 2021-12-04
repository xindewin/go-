package main

import (
	"fmt"
	"net"
	"os"
)

//错误处理函数封装
func errFunc(err error, info string) {
	if err!=nil{
		fmt.Println(info,err)
		//return //返回当前函数调用
		//runtime.Goexit()  //结束当前go程
		os.Exit(1)    //结束当前进程
	}
}

func main() {
	listen,err:=net.Listen("tcp","127.0.0.1:8000")
    errFunc(err,"listen")
    defer listen.Close()

	fmt.Println("等待服务器进行连接")
	con,err:=listen.Accept()
	errFunc(err,"Accept")
	defer con.Close()

	buf:=make([]byte,4096)

	n,err:=con.Read(buf)
    errFunc(err,"Read")
	fmt.Println("客户端请求",string(buf[:n]))
}
