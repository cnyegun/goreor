package main

import (
	def "goreor/packages"
)

func main() {
	mytask := def.Task{
		Priority: "High",
		Due:      "24-11-2024",
		Added:    "13-11-2024",
		Content:  "Go to LIDL buy meat",
	}
	def.PrintTask(mytask)
}

func home_menu_loop() {
	for {

	}
}
