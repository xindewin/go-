package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
)


//请求 json

//图片请求包
type Requestpicture struct {
	Title string `json:"title"`
	Isbn string `json:"isbn"`
	RecordId float64 `json:"recordId"`
}

//定义解析图片返回json字段
type ResponsePic struct {
	DuxiuImageUrl string `json:"duxiuImageUrl"`
	ECount float64 `json:"ECount"`
	PCount float64 `json:"PCount"`
}

//定义请求json结构体
type RequestBody2 struct {
	DocCode []string `json:"docCode"`
	SearchFieldContent string `json:"searchFieldContent"`
	SearchField string `json:"searchField"`
	MatchMode string `json:"matchMode"`
	ResourceType []string `json:"resourceType"`
	Subject []string `json:"subject"`
	Discode1 []string `json:"discode1"`
	Publisher []string `json:"publisher"`
	LibCode []string `json:"libCode"`
	LocationId []string `json:"locationId"`
	ECollectionIds []string `json:"eCollectionIds"`
	CurLocationId []string `json:"curLocationId"`
	CampusId []string `json:"campusId"`
	KindNo []string `json:"kindNo"`
	CollectionName []string `json:"collectionName"`
	Author []string `json:"author"`
	LangCode []string `json:"langCode"`
	CountryCode []string `json:"countryCode"`
	PublishBegin *string `json:"publishBegin"`
	PublishEnd *string `json:"publishEnd"`
	CoreInclude []string `json:"coreInclude"`
	DdType []string `json:"ddType"`
	VerifyStatus []string `json:"verifyStatus"`
	Group []string `json:"group"`
	SortField string `json:"sortField"`
	SortClause string `json:"sortClause"`
	Page int `json:"page"`
	Rows int`json:"rows"`
	OnlyOnShelf *string `json:"onlyOnShelf"`
	IndexSearch int `json:"indexSearch"`
	SearchItems *string `json:"searchItems"`

}

//定义解析json结构体字段

type (
	booklist struct {
		BookId float64 `json:"bookId"`
		Name string `json:"name"`
		Author string `json:"author"`
		Publisher string `json:"publisher"`
		Isbn string `json:"isbn"`
		Pcount float64 `json:"pcount"`
		Ecount float64 `json:"ecount"`
		SearchCode string `json:"searchCode"`
		Image string `json:"image"`
		//StatusNow string `json:"statusNow"`
		//Status bool `json:"status"`
	}

	data struct{
		All float64 `json:"all"`
		Booklist []booklist `json:"booklist"`
	}

	result struct {
		Status string `json:"status"`
		Msg string `json:"msg"`
		Data data `json:"data"`

	}

)



func GetPost(BookName string,page int,row int){
	//构建请求并获得返回
	request:=RequestBody2{
		DocCode: []string{},
		SearchFieldContent: BookName,
		SearchField: "keyWord",
		MatchMode: "2",
		ResourceType: []string{},
		Subject: []string{},
		Discode1: []string{},
		Publisher: []string{},
		LibCode: []string{},
		LocationId: []string{},
		ECollectionIds: []string{},
		CurLocationId: []string{},
		CampusId: []string{},
		KindNo: []string{},
		CollectionName: []string{},
		Author: []string{},
		LangCode: []string{},
		CountryCode: []string{},
		CoreInclude:[]string{},
		DdType: []string{},
		VerifyStatus: []string{},
		Group:[]string{},
		SortField: "relevance",
		SortClause: "asc",
		Page: 1,
		Rows: 10,
		IndexSearch: 1,
	}

	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(request)

	url := "https://findcumt.libsp.com/find/unify/search"
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

	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))
	doc:=gjson.ParseBytes(body).Get("data.searchResult").Array()


	var magic []string=make([]string,0)

   //图片请求
   for _,j:=range doc{
	   requestpic:=Requestpicture{
		   Title: j.Get("title").String(),
		   Isbn: j.Get("isbn").String(),
		   RecordId: j.Get("recordId").Num,
	   }

	   requestBody2 := new(bytes.Buffer)
	   json.NewEncoder(requestBody2).Encode(requestpic)

	   urlpic := "https://findcumt.libsp.com/find/unify/getPItemAndOnShelfCountAndDuxiuImageUrl"
	   req2, err := http.NewRequest("POST", urlpic, requestBody2)

	   req2.Header.Set("Host","findcumt.libsp.com")
	   req2.Header.Set("Connection","keep-alive")
	   req2.Header.Set("Accept","application/json, text/plain, */*")
	   req2.Header.Set("Accept-Language","zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	   req2.Header.Set("SAccept-Encoding","gzip, deflate")
	   req2.Header.Set("Content-Type", "application/json;charset=utf-8")
	   req2.Header.Set("mappingPath","")
	   req2.Header.Set("groupCode","200069")
	   //req2.Header.Set("sec-ch-ua","\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"96\",\"Google Chrome\";v=\"96\"")
	   //req2.Header.Set("sec-ch-ua-mobile","?1")
	   req2.Header.Set("x-lang","CHI")
	   req2.Header.Set("sec-ch-ua-platform","Android")
	   req2.Header.Set("Origin","https://findcumt.libsp.com")
	   req2.Header.Set("Referer","https://findcumt.libsp.com/")
	   req2.Header.Set("Cookie","SameSite=None")
	   req2.Header.Set("Sec-Fetch-Dest","empty")
	   req2.Header.Set("Sec-Fetch-Mode","cors")
	   req2.Header.Set("Sec-Fetch-Site","same-origin")
	   //fmt.Println(req)

	   client2 := &http.Client{}
	   resp2, err := client2.Do(req2)

	   if err != nil {
		   panic(err)
	   }
	   defer resp2.Body.Close()

	   body2,_:=ioutil.ReadAll(resp2.Body)
	  // fmt.Println(string(body2))
	  // doc2:=gjson.ParseBytes(body2).Get("data").Array()

	   magic=append(magic,gjson.ParseBytes(body2).Get("data.duxiuImageUrl").String())

   }
   //fmt.Println(magic)
   /*BookId:     i.Get("recordId").Num,
   Name:       i.Get("title").String(),
   Author:     i.Get("author").String(),
   Publisher:  i.Get("publisher").String(),
   Isbn:       i.Get("isbn").String(),
   Pcount:     i.Get("physicalCount").Num,
   Ecount:     i.Get("groupECount").Num,
   SearchCode: i.Get("callNoOne").String(),*/
	/*requestpic:=Requestpicture{
		Title: "计算机组成原理",
		Isbn: "978-7-5636-3726-3",
		RecordId: 464605,
	}*/
//{"title":"计算机组成原理","isbn":"978-7-5636-3726-3","recordId":464605}

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)


		//fmt.Println(i)
        //fmt.Printf("当前%s\n",i.Get("title").String())
		r:=result{
			Status: resp.Status,
			Msg:    gjson.ParseBytes(body).Get("message").String(),
			Data: data{
				All: gjson.ParseBytes(body).Get("data.numFound").Num,
				Booklist: []booklist{

				},
			},
		}


		for k,i:=range doc {
			fmt.Println(magic[k])
			book := booklist{
				BookId:     i.Get("recordId").Num,
				Name:       i.Get("title").String(),
				Author:     i.Get("author").String(),
				Publisher:  i.Get("publisher").String(),
				Isbn:       i.Get("isbn").String(),
				Pcount:     i.Get("physicalCount").Num,
				Ecount:     i.Get("groupECount").Num,
				SearchCode: i.Get("callNoOne").String(),
				Image: magic[k],
			}
			r.Data.Booklist = append(r.Data.Booklist, book)
			fmt.Println(k)
		}
			var content []byte
			content, err = json.Marshal(r)
			if err != nil {
				log.Fatal("Marshal error", err)
			}


		fmt.Println(string(content))





}




func main() {
	GetPost("计算机组成原理",1,10)
}
