package def

import (
	"fmt"
)

type Task struct {
	// Priority is one of {"Urgent", "High","Low"}
	Priority string
	// Due and Added is date in format "DD-MM-yyyy"
	Due     string
	Added   string
	Content string
}

type Label struct {
	Name string
	Task []Task
}

func PrintTask(t Task) {
	fmt.Println("@Task INFO:")
	fmt.Println(" >Priority:", t.Priority)
	fmt.Println(" >Added   :", t.Added)
	fmt.Println(" >Due     :", t.Due)
	fmt.Println(" >Content :", t.Content)
}
