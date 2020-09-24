package main

import (
	"fmt"
	"log"
	"os"

	"task_2/calc"
)

func main() {
	if len(os.Args) == 0 {
		log.Println("Empty input")
		os.Exit(1)
	}

	expr := os.Args[1]
	res, err := calc.Calculate(expr)
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}
	fmt.Println(res)
}
