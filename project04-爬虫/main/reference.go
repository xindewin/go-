package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const (
	rePic=`https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

// 获取网页数据，且把数据转成 字符串
func getStr(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("http.Get error : ", err)
	}
	defer resp.Body.Close()

	// 去读数据内容为 bytes
	dataBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("ioutil.ReadAll error : ", err)
	}

	// 字节数组 转换成 字符串
	str := string(dataBytes)
	return str
}

// 获取图片数据
func GetPic(url string, n int) {

	str := getStr(url)

	// 过滤 图片
	re := regexp.MustCompile(rePic)

	// 匹配多少次， -1 默认是全部
	results := re.FindAllStringSubmatch(str, n)

	// 输出结果
	for _, result := range results {
		fmt.Println(result[0])
		// 获取具体的图片名字
		fileName := GetFileName(result[0])
		fmt.Println(fileName)
		// 下载图片
		DownloadPic(result[0], fileName)
	}
}

// 获取到 文件的 名字
func GetFileName(url string) (filename string) {
	// 找到最后一个 = 的索引
	lastIndex := strings.LastIndex(url, "/")
	// 获取 / 后的字符串 ，这就是源文件名
	filename = url[lastIndex+1:]

	// 把时间戳 加 在原来名字前，拼一个新的名字
	prefix := fmt.Sprintf("%d",time.Now().Unix())
	filename = prefix + "_" + filename

	return filename
}

func DownloadPic(url string, filename string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("http.Get error : ", err)
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("ioutil.ReadAll error : ", err)
	}

	// 文件存放的路径
	filepath:= "./" + filename

	// 写文件 并设置文件权限
	err = ioutil.WriteFile(filepath, bytes, 0666)
	if err != nil {
		log.Fatal("wirte failed !!", err)
	} else {
		log.Println("ioutil.WriteFile successfully , filename = ", filename)
	}
}

func main() {
	// 简单设置l og 参数
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	GetPic("https://zhidao.baidu.com/question/689879375452903724.html", 1)
}