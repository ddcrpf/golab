package main

import (
	"fmt"

	"github.com/ddcrpf/golab/programs"
)

func main() {
	programList := map[int]func(){
		1: programs.Program1,
		2: programs.Program2,
		3: programs.Program3,
		4: programs.Program4,
		5: programs.Program5,
		6: programs.Program6,
		7: programs.Program7,
		8: programs.Program8,
		// Add the rest...
	}

	fmt.Println("Choose a program to execute (1-8):")
	var choice int
	fmt.Scanln(&choice)

	if program, exists := programList[choice]; exists {
		program()
	} else {
		fmt.Println("Invalid choice")
	}
}
