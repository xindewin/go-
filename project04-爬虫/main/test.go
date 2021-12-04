package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)


type RequestBody struct {
	Q string `json:"q"`
	Limit int `json:"limit"`
    Offset int `json:"offset"`
}

func testPost(q string,limit int,offset int) {

	request := RequestBody{
		Q: q,
		Limit: limit,
		Offset: offset,

	}
	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(request)

	url := "https://findcumt.libsp.com/find/chaoxing/searchList"
	req, err := http.NewRequest("POST", url, requestBody)

	req.Header.Set("Host","findcumt.libsp.com")
	req.Header.Set("Accept","application/json, text/plain, */*")
	req.Header.Set("Accept-Language","zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	req.Header.Set("SAccept-Encoding","gzip, deflate")
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.Header.Set("mappingPath","")
	req.Header.Set("groupCode","200069")
	req.Header.Set("x-lang","CHI")
	req.Header.Set("Origin","https://findcumt.libsp.com")
	req.Header.Set("Referer","https://findcumt.libsp.com/")
	req.Header.Set("Cookie","SameSite=None")
	req.Header.Set("Sec-Fetch-Dest","empty")
	req.Header.Set("Sec-Fetch-Mode","cors")
	req.Header.Set("Sec-Fetch-Site","same-origin")
	//fmt.Println(req)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()


	fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}








//获取渲染后网页
/*func ChromedpGecontent(url string) string {
	ctx,cancel:=chromedp.NewContext(context.Background(),chromedp.WithLogf(log.Printf))
	defer cancel()
	var response string
	err:=chromedp.Run(ctx,chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.OuterHTML("#app",&response),
	})
	if err!=nil{
		log.Println(err)
		return ""
	}
	//创造文件
	file,err:=os.Create("content2.html")
	if err!=nil{
		log.Fatal("creat file fail",err)
	}
	//写入文件
	file.WriteString(response)
	defer file.Close()
	return response
}*/


func main(){
    testPost("计算机组成原理",10,10)

   //str:=ChromedpGecontent("https://findcumt.libsp.com/src/components/common/js/jquery.min.js#/searchList")

	//log.Fatal()
	//fmt.Println(str)
	/*res,err:=http.Get("https://findcumt.libsp.com/#/searchListExternal/01/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BB%84%E6%88%90%E5%8E%9F%E7%90%86/01")
	if err!=nil{
		log.Fatal("get error",err)
	}
	defer res.Body.Close()
	databyte,err:=ioutil.ReadAll(res.Body)
	if err!=nil{
		log.Fatal("readall error",err)
	}
	fmt.Println(string(databyte))
	doc:=gjson.ParseBytes(databyte)
	fmt.Println(doc.Array())*/




/*//定义解析字段
type ResultForJson struct{

}
*/

}
