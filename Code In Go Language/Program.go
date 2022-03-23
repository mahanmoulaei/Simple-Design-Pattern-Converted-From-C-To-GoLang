package main

import (
	"fmt"
)

func main() {
	var input string
	
	chain := new(Fan)
	chain.Initialize()

	for {
		input = getLine()
		if input == "0" {
			break
		} else {
			chain.pull()
			//fmt.Println("currentState:", reflect.TypeOf(chain.currentState)) //for testing
		}

	}
}

func getLine() string {
	var temp string
	fmt.Println("\nPress Enter (type 0 to exit)")
	fmt.Scanln(&temp)
	return temp
}
