package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)


func ToJsonFromFile(directory string) []map[string]interface{} {
	data, err := ioutil.ReadFile(directory)
	data = bytes.Replace(data, []byte("\\'"), []byte(`'`), -1)
	data = bytes.Replace(data, []byte("\t"), []byte(` `), -1)
	data = bytes.Replace(data, []byte("\r"), []byte(` `), -1)
	data = bytes.Replace(data, []byte("\n"), []byte(` `), -1)

	categories := make([]map[string]interface{}, 0)
	if err != nil {
		fmt.Println("in file: ", directory, err)
		return categories
	}

	err = json.Unmarshal(data, &categories)
	if err != nil{
		fmt.Println("error unmarshalling ", directory, err)
	}

	return categories
}