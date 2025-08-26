package main

import (
	"fmt"

	"github.com/containerd/console"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func center(s string, w int) string {

	return fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(s))/2, s))

}

func main() {
	winSize, err := console.Current().Size()

	if err != nil {
		fmt.Println("Error geting console width")
	}

	width := int(winSize.Width)

	fmt.Println(center(Cyan+"Library Management System"+Reset, width))
	fmt.Println(center(Magenta+"Select one of the options..."+Reset, width))
	fmt.Println(center(Yellow+"1. Add Book "+Reset, width))
	fmt.Println(center(Yellow+"2. Delete Book "+Reset, width))
	fmt.Println(center(Yellow+"3. Find Book By Id "+Reset, width))
	fmt.Println(center(Yellow+"4. View All Books "+Reset, width))

}
