package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func PrintBody(r *http.Response) {
	defer func() {_=r.Body.Close()}()
	content,err:=ioutil.ReadAll(r.Body)
	if err!=nil {
		panic(err)
	}
	fmt.Printf("%s",content)
}

func requestByParams() {
	request,err:=http.NewRequest(http.MethodGet,"http://httpbin.org/get",nil)
	if err != nil{
		panic(err)
	}

	params:=make(url.Values)//组织get请求参数
	params.Add("name","xinxin")
	params.Add("age","19")

	fmt.Println(params.Encode())
	request.URL.RawQuery=params.Encode()
	r,err:=http.DefaultClient.Do(request)
	if err!=nil {
		panic(err)
	}
	PrintBody(r)

}
//定制请求头
func requestByHead()  {
	request,err:=http.NewRequest(http.MethodPost,"https://findcumt.libsp.com",nil)
	if err!=nil{
		panic(err)
	}
	request.Header.Add("user-agent","huohu")
	r,err:=http.DefaultClient.Do(request)
	if err!=nil {
		panic(err)
	}
	PrintBody(r)


}

func main() {
	//设置请求参数
	//定制请求头
	requestByHead()

}
