package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	//指定服务器，通信协议，ip地址，port。创建一个用于监听的socket
	listner,err:=net.Listen("tcp","192.168.75.1:8000")
	if err!=nil{
		log.Fatal("net.listen error",err)
	}
    defer listner.Close()
	fmt.Println("服务器等待客户端建立连接。。。")
   //阻塞监听客户端连接请求,成功建立连接，返回用于通信的socket
	con,err:=listner.Accept()
	if err != nil {
		log.Fatal("listener.Accept error",err)
	}
	defer con.Close()
	fmt.Println("服务器与客户端连接成功")
	//读取客户端发送数据
	buf:= make([]byte,4096)
	n,err:=con.Read(buf)  //返回值为int
	if err!=nil{
		log.Fatal("con.Read error",err)
	}
    //读到什么写什么
	con.Write(buf[:n])
	//处理数据--打印
	fmt.Println("服务器读到数据",string(buf[:n]))


}
