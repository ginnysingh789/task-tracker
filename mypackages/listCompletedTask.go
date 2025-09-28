package mypackages

//List the completed Task
import (
	"encoding/json"
	"fmt"
)

func ListCompletedTask(TaskStorage map[int]Properties) {

	for key, value := range TaskStorage {
		if value.Status == "Completed" {
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
