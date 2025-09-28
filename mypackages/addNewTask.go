package mypackages

import (
	"fmt"
	"time"
)

func AddNewTask(TaskStorage map[int]Properties, count *int) {
	//Now read the descritpion
	fmt.Println("Enter the description")
	desc := UserInput()
	(TaskStorage)[*count] = Properties{ID: *count, Description: desc, Status: "in-progress", CreatedAt: (time.Now().Format("2006-01-02 15:04:05")), UpdatedAt: (time.Now().Format("2006-01-02 15:04:05"))}
	(*count)++

}
