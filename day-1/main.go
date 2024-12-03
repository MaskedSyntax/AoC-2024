package main

import (
	"fmt"
	"os"

	"github.com/maskedsyntax/AoC-2024/day-1/part1"
	"github.com/maskedsyntax/AoC-2024/day-1/part2"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [part1|part2]")
	}

	part := os.Args[1]
	switch part {
	case "part1":
		part1.Run()
	case "part2":
		part2.Run()
	default:
		fmt.Println("Invalid Argument, Use 'part1' or 'part2'")
	}

}
