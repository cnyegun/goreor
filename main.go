package main

import (
	// "fmt"
	"fmt"
	def "goreor/packages"
)

func main() {
	home_menu_loop()
}

func home_menu_loop() {
	labels := []def.Label{
		{
			Name: "Work",
			Tasks: []def.Task{
				{Id: 1, Priority: "Urgent", Due: "15-11-2024", Added: "01-11-2024", Content: "Finish report"},
				{Id: 2, Priority: "High", Due: "20-11-2024", Added: "05-11-2024", Content: "Prepare presentation"},
			},
		},
		{
			Name: "Personal",
			Tasks: []def.Task{
				{Id: 3, Priority: "Low", Due: "25-12-2024", Added: "10-11-2024", Content: "Buy gifts"},
				{Id: 4, Priority: "High", Due: "22-11-2024", Added: "12-11-2024", Content: "Book tickets"},
			},
		},
		{
			Name: "Fitness",
			Tasks: []def.Task{
				{Id: 5, Priority: "Urgent", Due: "15-11-2024", Added: "10-11-2024", Content: "Renew gym membership"},
				{Id: 6, Priority: "Low", Due: "20-11-2024", Added: "05-11-2024", Content: "Buy new running shoes"},
			},
		},
		{
			Name: "Finance",
			Tasks: []def.Task{
				{Id: 7, Priority: "Low", Due: "25-12-2024", Added: "20-6-2024", Content: "Pay student loan"},
				{Id: 8, Priority: "High", Due: "02-01-2025", Added: "21-6-2024", Content: "Buy more BTC"},
			},
		},
	}

	var cmd string
	var opt string
	var opt2 string
	def.PrintHelp()
	for {
		opt = ""
		cmd = ""
		opt2 = ""

		fmt.Print("=go.reor> ")
		fmt.Scanln(&cmd, &opt, &opt2)

		switch cmd {
		case "ls":
			def.ListLabels(labels)
		case "see":
			def.SeeLabel(labels, opt)
		case "cl":
			labels = def.NewLabel(labels, opt)
		case "rm":
			labels = def.RemoveLabel(labels, opt)
		case "cn":
			labels = def.ChangeName(labels, opt, opt2)
		case "ct":
			def.CreateTask(labels, opt)
		default:
			fmt.Print("Invalid syntax '/?' for help\n")
		}
	}
}
