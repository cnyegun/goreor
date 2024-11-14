package def

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	Id       int
	Priority string // {"Urgent", "High", "Low"}
	Due      string // "DD-MM-yyyy"
	Added    string // "DD-MM-yyyy"
	Content  string
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
	fmt.Printf("Cannot find the label:%s\n", name)
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
			l[i].Name = new_name
			return l
		}
	}
	fmt.Printf("Cannot find the label:%s\n", name)
	return l
}

// []Label String -> []Label
// Create a task inside the given Label name String
func CreateTask(l []Label, name string) []Label {
	var newTask Task
	for i := range len(l) {
		if l[i].Name == name {
			newTask = getTaskInfo()
			newTask.Id = len(l[i].Tasks) + 1
			l[i].Tasks = append(l[i].Tasks, newTask)
			return l
		}
	}
	fmt.Printf("Cannot find the label:%s\n", name)
	return l
}

// Task -> Task
// Get user input for Task property
func getTaskInfo() Task {
	var newTask Task
	//input priority
	for !isPriority(newTask.Priority) {
		fmt.Printf("Priority is one of {\"Urgent\", \"High\",\"Low\"}\n")
		fmt.Printf("Priority: ")
		fmt.Scan(&newTask.Priority)
	}
	//input due
	for !isValidDate(newTask.Due) {
		fmt.Printf("Due date must be in format DD-MM-yyyy\n")
		fmt.Printf("Due: ")
		fmt.Scan(&newTask.Due)
	}
	//time added
	newTask.Added = time.Now().Format("02-01-2006")
	//content
	fmt.Print("Content: ")
	reader := bufio.NewReader(os.Stdin)
	newTask.Content, _ = reader.ReadString('\n')
	return newTask
}

// String -> Boolean
// False if String is not one of Urgent, High, Low
func isPriority(s string) bool {
	validPriority := []string{
		"Urgent",
		"High",
		"Low",
	}
	for i := range len(validPriority) {
		if s == validPriority[i] {
			return true
		}
	}
	return false
}

// String -> Boolean
// Verify if the string is in format DD-MM-YYYY
func isValidDate(date string) bool {
	layout := "02-01-2006"
	_, err := time.Parse(layout, date)
	return err == nil
}

// save data to a JSON file
func SaveData(labels []Label, filename string) error {
	data, err := json.MarshalIndent(labels, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling data: %v", err)
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}
	return nil
}

// load data from a JSON file
func LoadData(filename string) ([]Label, error) {
	// return empty slice if file doesn't exist
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return []Label{}, nil
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	// Unmarshal the JSON data into a slice of Label structs
	var labels []Label
	err = json.Unmarshal(data, &labels)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling data: %v", err)
	}

	return labels, nil
}
