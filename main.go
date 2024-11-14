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
	labels, _ := def.LoadData("data.json")

	var cmd string
	var opt string
	var opt2 string
	def.PrintHelp()
	for {
		def.SaveData(labels, "data.json")
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
		case "/?":
			def.PrintHelp()
		default:
			fmt.Print("Invalid syntax '/?' for help\n")
		}
	}
}
