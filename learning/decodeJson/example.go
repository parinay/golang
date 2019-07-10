package main

import (
	"encoding/json"
	"fmt"
)

var JSON = `{
	"name":"Mark Taylor",
	"jobtitle": "software Dev",
	"phone": "1829990",
	"email":"ancd@gmail.com"
}`

func main() {
	// var info map[string]interface{}
	// var info map[string]string

	fmt.Printf("Json type %T\n", JSON)

	info := make(map[string]string)

	err := json.Unmarshal([]byte(JSON), &info)

	if err != nil {
		fmt.Println("Incorrect JSON string")
		panic(err)
	} else {
		for key, value := range info {
			fmt.Println("index: ", key, "value: ", value)
		}
	}
}
