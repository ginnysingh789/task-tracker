package mypackages

import (
	"fmt"
	"strconv"
	"time"
)

func UpdateTask(TaskStorage map[int]Properties) {
	fmt.Println("Enter the Task ID")
	getId := UserInput()
	userID, _ := strconv.Atoi(getId)
	//check task exist
	task, exist := TaskStorage[userID]
	if !exist {
		fmt.Println("There is no task with this ID")
		return
	}

	fmt.Println("Enter the new Task status ... 1->Cancel the Task  2->Marks Task as Completed")
	getValue := UserInput()
	value, err := strconv.Atoi(getValue)
	if err != nil || (value != 1 && value != 2) { // Combine error checking and validation
		fmt.Println("Enter the correct Value ...")
		return
	}
	if value == 1 {

		task.Status = "canceled"
	} else {
		task.Status = "Completed"
	}
	task.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	TaskStorage[userID] = task // Update the task in the storage
	fmt.Println("Task status updated successfully ...")

}

//Udpate the time task status and also ensure that task exist
