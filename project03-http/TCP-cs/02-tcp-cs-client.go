package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	con,err:=net.Dial("tcp","192.168.75.1:8000")
	if err!=nil{
		log.Fatal("net.Dial error",err)
	}
	defer  con.Close()
	fmt.Println("客户端与服务器连接成功")

	//主动写数据给服务器
	buf:=make([]byte,4096)
	con.Write([]byte("重言师傅nb！！！"))
	//读取服务器回发数据
	n,err:=con.Read(buf)
	if err != nil {
		log.Fatal("Write error",err)
	}
	fmt.Println("服务器回发",string(buf[:n]))


}