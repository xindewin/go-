package main

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	"regexp"
	"strings"
)

const(
	repic1=`[A-Z]+[0-9]+\/[A-Z]+-[0-9]+`
    repic2=`[\p{Han}]+-[\p{Han}]+`
    repic3=`C[0-9]+`
)

//获取渲染后网页
func ChromedpGecontent(url string) string {
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
	defer cancel()
	var response string
	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.OuterHTML("[class~=ant-table-row]", &response),
	})
	if err != nil {
		log.Println(err)
		return ""
	}
	//创造文件
	file, err := os.Create("content2.html")
	if err != nil {
		log.Fatal("creat file fail", err)
	}
	//写入文件
	file.WriteString(response)
	defer file.Close()
	return response
}

//解析html
func findByEle1(htmlDoc string) {
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

func main() {
	var ResultBookId=make(map[string]string,3)


	str:=ChromedpGecontent("https://findcumt.libsp.com/#/searchList/bookDetails/744526")
	//fmt.Println(str)

	re:=regexp.MustCompile(repic1)
	results:=re.FindAllStringSubmatch(str,1)
	ResultBookId["索书号"]=results[0][0]
	//fmt.Println(results[0][0])

	re=regexp.MustCompile(repic2)
	results=re.FindAllStringSubmatch(str,1)
	ResultBookId["馆藏地点"]=results[0][0]
	//fmt.Println(results[0][0])
	re=regexp.MustCompile(repic3)
	results=re.FindAllStringSubmatch(str,-1)
	ResultBookId["条码号"]=results[0][0]

	fmt.Println(ResultBookId)





}
