package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Name     string
	Complete bool
}

var tasks []Task

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nTo-Do List CLI")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Mark Task Complete")
		fmt.Println("4. Delete Completed Tasks")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			addTask(reader)
		case "2":
			viewTasks()
		case "3":
			markComplete(reader)
		case "4":
			deleteCompletedTasks()
		case "5":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func addTask(reader *bufio.Reader) {
	fmt.Print("Enter task name: ")
	taskName, _ := reader.ReadString('\n')
	taskName = strings.TrimSpace(taskName)
	tasks = append(tasks, Task{Name: taskName, Complete: false})
	fmt.Println("Task added!")
}

func viewTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	for i, task := range tasks {
		status := "Incomplete"
		if task.Complete {
			status = "Complete"
		}
		fmt.Printf("%d. %s [%s]\n", i+1, task.Name, status)
	}
}

func markComplete(reader *bufio.Reader) {
	if len(tasks) == 0 {
		fmt.Println("No tasks to mark as complete.")
		return
	}
	viewTasks()
	fmt.Print("Enter the task number to mark as complete: ")
	var taskNum int
	fmt.Scanln(&taskNum)
	if taskNum <= 0 || taskNum > len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}
	tasks[taskNum-1].Complete = true
	fmt.Println("Task marked as complete!")
}

func deleteCompletedTasks() {
	var incompleteTasks []Task
	for _, task := range tasks {
		if !task.Complete {
			incompleteTasks = append(incompleteTasks, task)
		}
	}
	tasks = incompleteTasks
	fmt.Println("Completed tasks deleted!")
}
