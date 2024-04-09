package main

import (
	"flag"
	"fmt"
	"os"
)

type Task struct {
	Title string
	Done  bool
}

var tasks = []Task{}

func addTask(title string) {
	tasks = append(tasks, Task{Title: title, Done: false})
	fmt.Println("Added task:", title)
}

func listTasks() {
	for i, task := range tasks {
		status := "Not Done"
		if task.Done {
			status = "Done"
		}
		fmt.Printf("%d. %s [%s]\n", i+1, task.Title, status)
	}
}

func markDone(index int) {
	if index <= 0 || index > len(tasks) {
		fmt.Println("Invalid task number:", index)
		return
	}
	tasks[index-1].Done = true
	fmt.Println("Marked task as done:", tasks[index-1].Title)
}

func main() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addTitle := addCmd.String("title", "", "Title of the task")

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	doneCmd := flag.NewFlagSet("done", flag.ExitOnError)
	doneIndex := doneCmd.Int("index", 0, "Index of the task to mark as done")

	if len(os.Args) < 2 {
		fmt.Println("expected 'add', 'list' or 'done' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		if *addTitle == "" {
			fmt.Println("title flag is required for the add command")
			os.Exit(1)
		}
		addTask(*addTitle)
	case "list":
		listCmd.Parse(os.Args[2:])
		listTasks()
	case "done":
		doneCmd.Parse(os.Args[2:])
		if *doneIndex == 0 {
			fmt.Println("index flag is required for the done command")
			os.Exit(1)
		}
		markDone(*doneIndex)
	default:
		fmt.Println("expected 'add', 'list' or 'done' subcommands")
		os.Exit(1)
	}
}
