package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func responseBody(r *http.Response) {
	content,_:=ioutil.ReadAll(r.Body)
	fmt.Printf("%s",content)
}
func status(r *http.Response) {
    fmt.Println(r.StatusCode) //状态码
	fmt.Println(r.Status)   //描述
}
func header(r *http.Response) {
	//get不区分大小写
	fmt.Println(r.Header.Get("Content-type"))
	//fmt.Println(r.Header["Content-type"])
}

func encoding(r *http.Response) {
	//获取相应的编码信息
  //content-type 中会提供编码
	//html head meta 获取编码
	//可以通过网页的头部猜测网页编码信息

}


func main() {
	r,err:=http.Get("http://httpbin.org/get")
	if err!=nil{
		panic(err)
	}
	defer func() {_=r.Body.Close()}()

	//responseBody(r)
	//status(r)
    header(r)


}
