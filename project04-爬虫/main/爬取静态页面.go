package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

const(
	RePic=`/_upload[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
	RePic2=`https?://xwzx[^"]+?(\.((html)|(htm)))`
)
//解析html
func findByEle(htmlDoc string) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(htmlDoc))
	if err != nil {
		log.Fatalln(err) // 输出并退出程序
	}
	dom.Find("[class~=ant-table-row]").Each(func(i int, s *goquery.Selection) {
		fmt.Println("i", i, "选中的文本", s.Text())
	})

	/*dom.Find("p").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})*/
	/*// 1.元素选择器
	dom.Find("div").Each(func(i int, s *goquery.Selection) {
		fmt.Println("i", i, "选中的文本", s.Text())
	})
	// 2.ID选择器
	dom.Find("#div1").Each(func(i int, s *goquery.Selection) {
		fmt.Println("i", i, "选中的文本", s.Text())
	})
	dom.Find("div#div1").Each(func(i int, s *goquery.Selection) {
		fmt.Println("i", i, "选中的文本", s.Text())
	})
	//3.class类选择器
	dom.Find(".c1").Each(func(i int, s *goquery.Selection) {
		fmt.Println("i", i, "选中的文本", s.Text())
	})
	dom.Find("div.c1").Each(func(i int, s *goquery.Selection) {
		fmt.Println("i", i, "选中的文本", s.Text())
	})*/
}

//获取返回内容的字节编码
func GetBytes(url string) []byte{
	resp,err:=http.Get(url)
	if err!=nil{
		log.Fatal("get_err",err)
	}
	defer resp.Body.Close()
	databytes,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		log.Fatal("reaall_err",err)
	}
	return databytes
}
//获取文件名
func getfileName(url string) (filename string) {
	//找到最后一个=的索引
	lastIndex:=strings.LastIndex(url,"/")
	//获取/后的字符串，这就是源文件名
	filename=url[lastIndex+1:]//左开右闭，从=的下一位开始取

	//把时间戳加在原来名字前，拼成一个新的名字
	prefix:=fmt.Sprintf("%d",time.Now().Unix())
	filename=prefix+"_"+filename

	return filename

}
//图片下载
func downPict(url string,filename string) {
	bytes:=GetBytes(url)
	//写入文件路径
	filepath:="./"+filename
	err:=ioutil.WriteFile(filepath,bytes,0666)
	if err!=nil{
		log.Fatal("WriteFile_error",err)
	}

}
//爬取网页
func getHtml(url string) {
	bytes:=GetBytes(url)
	str:=string(bytes)
	//创造文件
	file,err:=os.Create("content.html")
	if err!=nil{
		log.Fatal("creat file fail",err)
	}
	//写入文件
	file.WriteString(str)
	defer file.Close()
}
//爬图片
func GetPict (url string, n int) {
	bytes:=GetBytes(url)
	str:=string(bytes)
	re:=regexp.MustCompile(RePic)
	results:=re.FindAllStringSubmatch(str,n)

	for _,result:=range results{
		picurl:="http://www.cumt.edu.cn/"+result[0]
		filename:=getfileName(picurl)
		downPict(picurl,filename)
	}
}
//打印新闻
func printNews(url string) {
	bytes:=GetBytes(url)
	str:=string(bytes)
	findByEle(str)
}

//获取新闻内容
func GetNews(url string,n int) {
	bytes:=GetBytes(url)
	str:=string(bytes)
	re2:=regexp.MustCompile(RePic2)
	results:=re2.FindAllStringSubmatch(str,n)

	for _,result:=range results{
		fmt.Println(result)
		printNews(result[0])
	}
}

func main() {
	//getHtml("https://findcumt.libsp.com/#/searchListExternal/01/计算机组成原理/01")
	htmdoc:=string(GetBytes("https://findcumt.libsp.com/#/searchList/bookDetails/744526"))
	fmt.Println(htmdoc)
	findByEle(htmdoc)
	//GetNews("http://www.cumt.edu.cn/",-1)//-1输出全部
}
