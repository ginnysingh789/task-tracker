package mypackages

//List the completed Task
import (
	"encoding/json"
	"fmt"
)

func ListInProgressTask(TaskStorage map[int]Properties) {

	for key, value := range TaskStorage {
		if value.Status == "in-progress" {
			jsonData, err := json.MarshalIndent(TaskStorage[key], "", " ")
			if err != nil {
				fmt.Println("Error in converting the map into the json format")
				return
			}
			data := string(jsonData)
			fmt.Println(data)

		}
	}

}
