package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"task-tracker/mypackages"
)

func main() {
	// Global Map which stores key and value pair
	TaskStorage := make(map[int]mypackages.Properties) // Initialize here

	var count int = 1

	reader := bufio.NewReader(os.Stdin)

	// Main input loop to manage tasks
	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1 -> Add new Task")
		fmt.Println("2 -> Update Task")
		fmt.Println("3 -> Delete Task")
		fmt.Println("4 -> List all the tasks")
		fmt.Println("5 -> List all the tasks that are done")
		fmt.Println("6 -> List all the tasks that are not done")
		fmt.Println("7 -> List all the tasks that are in progress")
		fmt.Println("8 -> Exit")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		userInput, err := strconv.Atoi(input)
		if err != nil || userInput < 1 || userInput > 8 { // Validate input
			fmt.Println("Invalid choice. Please enter a number between 1 and 8.")
			continue
		}

		switch userInput {
		case 1:
			mypackages.AddNewTask(TaskStorage, &count)
		case 2:
			mypackages.UpdateTask(TaskStorage)
		case 3:
			mypackages.DeleteTask(TaskStorage)
		case 4:
			mypackages.ListAllTheTask(TaskStorage)
		case 5:
			mypackages.ListCompletedTask(TaskStorage)
		case 6:
			mypackages.ListUnCompletedTask(TaskStorage)
		case 7:
			mypackages.ListInProgressTask(TaskStorage)
		case 8:
			fmt.Println("Exiting program.")
			return // Exit the program gracefully
		}
	}
}
