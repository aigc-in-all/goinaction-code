// 这个示例程序展示如何序列化JSO字符串
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	m := make(map[string]interface{})
	m["name"] = "Gopher"
	m["title"] = "programmer"
	m["contact"] = map[string]interface{}{
		"home": "415.333.3333",
		"cell": "415.555.5555",
	}

	// 将这个map序列化到JSON字符串
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	fmt.Println(string(data))

	fmt.Println("------------------------")

	data, err = json.Marshal(m)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	fmt.Println(string(data))
}

/*Output
{
  "contact": {
    "cell": "415.555.5555",
    "home": "415.333.3333"
  },
  "name": "Gopher",
  "title": "programmer"
}
------------------------
{"contact":{"cell":"415.555.5555","home":"415.333.3333"},"name":"Gopher","title":"programmer"}
*/
