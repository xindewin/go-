package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)
const(
	//正则表达式匹配图片
	repic=`/_upload[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)
func getData (url string) {
	//获取网站数据
	resp,err:=http.Get(url)
	if err!=nil{
		log.Fatal("") //
	}
	defer resp.Body.Close()
	//读取数据内容
	dataBytes,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		log.Fatal("ioutil.ReadAll error:",err)
	}
	str:=string(dataBytes)

	//输出结果
	fmt.Println(str)

	//写入文件
	//创建文件
	file,err:=os.Create("content.html")
	//写入数据
	file.WriteString(str)
    defer file.Close()
}
//获取图片数据
func GetStr(url string) string {
	//获取网站数据
	resp,err:=http.Get(url)
	if err!=nil{
		log.Fatal("") //
	}
	defer resp.Body.Close()
	//读取数据内容
	dataBytes,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		log.Fatal("ioutil.ReadAll error:",err)
	}
	str:=string(dataBytes)
	return str
}

func getPic(url string,n int)  {
	str:=GetStr(url)
	re:=regexp.MustCompile(repic)
    results:=re.FindAllStringSubmatch(str,n)

	for _,result:=range results{
		fmt.Println(result[0])
		filename:=getFileName(result[0])
		picurl:="http://www.cumt.edu.cn/"+result[0]
		downPic(picurl,filename)
	}

}

func getFileName(url string) (filename string) {
	//找到最后一个=的索引
	lastIndex:=strings.LastIndex(url,"/")
	//获取/后的字符串，这就是源文件名
	filename=url[lastIndex+1:]//左开右闭，从=的下一位开始取

	//把时间戳加在原来名字前，拼成一个新的名字
	prefix:=fmt.Sprintf("%d",time.Now().Unix())
	filename=prefix+"_"+filename

	return filename

}

func downPic(url string,filename string) {
	resp,err:=http.Get(url)
	if err!=nil{
		log.Fatal("httpGet_err",err)
	}
	defer resp.Body.Close()
	bytes,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		log.Fatal("ioutil.ReadAll_err",err)
	}
	//写入文件路径
	filepath:="./"+filename
	err=ioutil.WriteFile(filepath,bytes,0666)
	if err!=nil{
		log.Fatal("WriteFile_error",err)
	}

}

func main() {
    getPic("http://www.cumt.edu.cn/",-1)
}