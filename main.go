package main

import "os"

func main() {
	switch os.Args[1] {
	case "parent":
		parent()
	case "child":
		child()
	case "parse":
		parse()
	default:
		panic("help")
	}
}
