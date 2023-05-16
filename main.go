package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	files, _ := filepath.Glob("source/*")

	for _, file := range files {
		fmt.Println("parse file: " + string(file))
		yamlResult, body := ParseFile(string(file))

		var data string
		for k, v := range yamlResult {
			fmt.Printf("key: %s \tvalue: %v\n", k, v)
			if v != nil {
				if k == "tags" {
					data += strings.Title(k) + ": "
					for _, tag := range v.([]interface{}) {
						data += tag.(string) + ", "
					}
					data = strings.TrimRight(data, ", ")
					data += "\n"
				} else if k != "published" {
					data += strings.Title(k) + ": " + fmt.Sprint(v) + "\n"
				}
			}
		}
		data += "\n" + string(body)
		outputFile := strings.Replace(string(file), "source", "result", 1)
		if _, err := os.Stat("result"); os.IsNotExist(err) {
			err := os.Mkdir("result", 0777)
			Check(err)
		}
		err := ioutil.WriteFile(outputFile, []byte(data), 0644)
		Check(err)
	}
}
