package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"nova/parser"
)

func main() {
	content, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println("Error when opening file: ", err)
		return
	}

	var data map[string]interface{}
	json.Unmarshal(content, &data)

	p := parser.NewParser()
	p.Parse(data)

	b, _ := json.MarshalIndent(data, "", "    ")
	fmt.Println(string(b))
}
