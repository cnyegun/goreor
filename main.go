package main

import (
	def "goreor/packages"
)

func main() {
	mytask := def.Task{
		Priority: "Urgent",
		Due:      "24-11-2024",
		Added:    "13-11-2024",
		Content:  "Learn docker & install arch",
	}
	def.PrintTask(mytask)
}
