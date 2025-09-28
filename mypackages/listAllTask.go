package mypackages

import "fmt"

func ListAllTheTask(TaskStorage map[int]Properties) {
	tasks, err := Getjson(TaskStorage)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tasks)

}
