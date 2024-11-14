package def

import (
	"fmt"
)

type Task struct {
	Id int
	// Priority is one of {"Urgent", "High","Low"}
	Priority string
	// Due and Added is date in format "DD-MM-yyyy"
	Due     string
	Added   string
	Content string
}

type Label struct {
	Name  string
	Tasks []Task
}

func printTask(t Task) {
	fmt.Println("@Task ID   :", t.Id)
	fmt.Println(" - Priority:", t.Priority)
	fmt.Println(" - Added   :", t.Added)
	fmt.Println(" - Due     :", t.Due)
	fmt.Println(" - Content :", t.Content)
	fmt.Println()
}

func PrintHelp() {
	fmt.Println(" _____  __ _  ___   _ __ ___  ___  _ __")
	fmt.Println("|_____|/ _` |/ _ \\ | '__/ _ \\/ _ \\| '__|")
	fmt.Println("|_____| (_| | (_) || | |  __/ (_) | |")
	fmt.Println("       \\__, |\\___(_)_|  \\___|\\___/|_|")
	fmt.Println("       |___/                           ")
	fmt.Println("============Available Commands============")
	fmt.Println("> ls")
	fmt.Println(" ?list all the labels")
	fmt.Println()
	fmt.Println("> see <label_name>")
	fmt.Println(" ?list all the tasks in that label")
	fmt.Println()
	fmt.Println("> cl <label_name>")
	fmt.Println(" ?create new label")
	fmt.Println()
	fmt.Println("> rm <label_name>")
	fmt.Println(" ?remove a label")
	fmt.Println()
	fmt.Println("> cn <label_name> <new_name>")
	fmt.Println(" ?change a label's name")
	fmt.Println()
	fmt.Println("> ct <label_name>")
	fmt.Println(" ?create a task in that label")
	fmt.Println("==========================================")

}

// []Label -> Stdout
// Print the name of all label in stdout
func ListLabels(l []Label) {
	length := len(l)
	for i := 0; i < length; i++ {
		fmt.Printf("[%s] ", l[i].Name)
	}
	fmt.Println()
}

// []Label String -> Stdout
// Search for label name and print all the label's task in stdout if matched
func SeeLabel(l []Label, name string) {
	length := len(l)
	for i := 0; i < length; i++ {
		if l[i].Name == name {
			printTasks(l[i])
			return
		}
	}
	fmt.Printf("Can't find label:%s\n", name)
}

func printTasks(l Label) {
	length := len(l.Tasks)
	for i := 0; i < length; i++ {
		printTask(l.Tasks[i])
	}
}

// []Label String -> []Label
// Create a new Label and put in []Label
func NewLabel(l []Label, name string) []Label {
	for i := 0; i < len(l); i++ {
		if l[i].Name == name {
			fmt.Printf("Label:%s existed\n", name)
			return l
		}
	}
	newLabel := Label{
		Name:  name,
		Tasks: []Task{},
	}
	return append(l, newLabel)
}

// []Label String -> []Label
// Remove a Label from []Label
func RemoveLabel(l []Label, name string) []Label {
	length := len(l)
	for i := 1; i < length; i++ {
		if l[i].Name == name {
			return append(l[:i], l[(i+1):]...)
		}
	}
	fmt.Printf("Cannot find the label:%s\n", name)
	return l
}

// []Label String String -> []Label
// Search for the matching label and change the name of it
func ChangeName(l []Label, name string, new_name string) []Label {
	length := len(l)
	for i := 1; i < length; i++ {
		if l[i].Name == name {
			return append(l[:i], l[(i+1):]...)
		}
	}
	fmt.Printf("Cannot find the label:%s\n", name)
	return l
}

// []Label String -> []Label
// Create a task inside the given Label name String
func CreateTask(l []Label, name string) {
}
