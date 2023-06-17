package main

import (
	"fmt"
	"gromise/gromise"
	"gromise/process"
)

func main() {

	funcs := []gromise.ProcessFunc{
		func() (interface{}, error) { return process.ForInt(0) },
		func() (interface{}, error) { return process.ForInt(1) },
		func() (interface{}, error) { return process.ForInt(2) },
		func() (interface{}, error) { return process.ForInt(3) },
		func() (interface{}, error) { return process.ForInt(4) },
		func() (interface{}, error) { return process.ForString("test") },
	}

	results, err := gromise.All(funcs)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, result := range results {
		switch v := result.(type) {
		case int:
			fmt.Println("Processed int:", v)
		case string:
			fmt.Println("Processed string:", v)
		default:
			fmt.Println("Unknown type")
		}
	}

	fmt.Println("All goroutines complete")
}
