package mypackages

import (
	"encoding/json"
	"fmt"
)

func Getjson(TaskStorage map[int]Properties) (string, error) {
	jsonData, err := json.MarshalIndent(TaskStorage, "", " ")
	if err != nil {

		return "", fmt.Errorf("error in converting map to the json fomat")
	}
	data := string(jsonData)

	return data, nil

}
