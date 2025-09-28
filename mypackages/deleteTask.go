package mypackages

import (
	"fmt"
	"strconv"
)

func DeleteTask(TaskStorage map[int]Properties) {
	fmt.Println("Enter the id ")
	getId := UserInput()
	userID, _ := strconv.Atoi(getId)
	var flag bool = false
	for _, value := range TaskStorage {
		if value.ID == userID {
			flag = true
		}
	}
	if !flag {
		fmt.Println("Entry not found in the database...")
		return
	}
	delete(TaskStorage, userID)
	fmt.Println("Task is deleted ...")

}
