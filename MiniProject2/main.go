package main

import (
	hogwarts "GOMODULE/MiniProject2/HogwartsTrainModule"
	hogwartsStructs "GOMODULE/MiniProject2/HogwartsUsingStructs"
	howzatt "GOMODULE/MiniProject2/Howzatt"
	"fmt"
)

func main() {
	hogwarts.HarryPotter()

	a := "\n\nSolution for problem 2 in mini project-2\n"
	fmt.Println(a)

	howzatt.Wankhade()

	//Using structs
	fmt.Println("\n\n\n\n\n\nThis is using structs")
	hogwartsStructs.HarryPotter()
}
