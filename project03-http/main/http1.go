package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func get() {
	r,err:=http.Get("https://findcumt.libsp.com")
	if err != nil{
		panic(err)
	}
	defer func(){_=r.Body.Close()}()

	content,err:= ioutil.ReadAll(r.Body)
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%s" ,content)

}

func post() {
	r,err:=http.Post("https://findcumt.libsp.com/find/chaoxing/searchList","",nil)
	if err != nil {
		panic(err)
	}
	defer func() {_=r.Body.Close()}()
	content,err:=ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s",content)
}

func put() {
	request,err:= http.NewRequest(http.MethodPut,"http://httpbin.org/put",nil)
	if err !=nil{
		panic(err)   //一个底层的方法实际上http.Get()是设置的快捷模式
	}
	//fmt.Println(request)
	r,err:=http.DefaultClient.Do(request)
	if err !=nil{
		panic(err)
	}

	defer func() {_=r.Body.Close()}()

	content,err:=ioutil.ReadAll(r.Body)
	if err !=nil{
		panic(err)
	}

	fmt.Printf("%s",content)
}

func del() {
	request,err:= http.NewRequest(http.MethodDelete,"http://httpbin.org/delete",nil)
	if err !=nil{
		panic(err)   //一个底层的方法实际上http.Get()是设置的快捷模式
	}
	//fmt.Println(request)
	r,err:=http.DefaultClient.Do(request)
	if err !=nil{
		panic(err)
	}

	defer func() {_=r.Body.Close()}()

	content,err:=ioutil.ReadAll(r.Body)
	if err !=nil{
		panic(err)
	}

	fmt.Printf("%s",content)
}

func main() {
	post()
}