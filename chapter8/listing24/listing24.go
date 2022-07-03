// 这个示例程序展示如何使用json包和NewDecoder函数
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type gResult struct {
	GsearchResultClass string `json:"GsearchResultClass"`
	UnescapedURL       string `json:"unescapedUrl"`
	URL                string `json:"url"`
	VisibleURL         string `json:"visibleUrl"`
	CacheURL           string `json:"cacheUrl"`
	Title              string `json:"title"`
	TitleNoFormatting  string `json:"titleNoFormatting"`
	Content            string `json:"content"`
}

type gResponse struct {
	ResponseData struct {
		Results []gResult `json:"results"`
	} `json:"responseData"`
}

func main() {
	// uri := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&rsz=8&q=golang"
	// resp, err := http.Get(uri)
	// if err != nil {
	// 	log.Println("ERROR:", err)
	// 	return
	// }
	// defer resp.Body.Close()

	// var gr gResponse
	// err = json.NewDecoder(resp.Body).Decode(&gr)
	// if err != nil {
	// 	log.Println("ERROR:", err)
	// 	return
	// }

	// fmt.Println(gr)

	test()
}

func test() {
	dat, err := os.ReadFile("data.json")
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	var gr gResponse
	err = json.Unmarshal(dat, &gr)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(gr)
}
