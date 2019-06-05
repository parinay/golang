package main

import (
	"encoding/json"
	"fmt"
)

var JSON = `{
	"name":"Mark Taylor",
	"jobtitle": "software Dev",
	"phone": {
		"home":"12-456-78",
		"office":"890-123-456"
	},
	"email":"ancd@gmail.com"
}`

func main() {
	// var info map[string]interface{}
	// var info map[string]string

	fmt.Printf("Json type %T\n", JSON)

	info := make(map[string]interface{})

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
