package main

import (
	"fmt"
	"os"

	"github.com/smartystreets/training-cohort/calcy-lib/calc"
)

func main() {
	handle := NewHandler(os.Stdout, calc.Addition{})
	result, err := handle.Handle(os.Args[1:])
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", result)
}
