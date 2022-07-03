// 这个示例程序展示如何解码JSON字符串
package main

import (
	"encoding/json"
	"log"
)

var JSON = `{
	"name": "Gopher",
	"title": "programmer",
	"contact": {
		"home": "415.333.3333",
		"cell": "415.555.5555"
	}
}`

func main() {
	// 将JSON字符串反序列化到map变量
	var m map[string]interface{}
	err := json.Unmarshal([]byte(JSON), &m)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	log.Println("Name:", m["name"])
	log.Println("Title:", m["title"])
	log.Println("Contact")
	log.Println("H:", m["contact"].(map[string]interface{})["home"])
	log.Println("C:", m["contact"].(map[string]interface{})["cell"])
}

/*Output
2022/06/30 12:22:41 Name: Gopher
2022/06/30 12:22:41 Title: programmer
2022/06/30 12:22:41 Contact
2022/06/30 12:22:41 H: 415.333.3333
2022/06/30 12:22:41 C: 415.555.5555
*/
