package main

import (
	"fmt"
	"github.com/Luke256/mylisp/lisp"
)

func main() {
	var input = `(println "Hello, World!")`

	box := lisp.NewBox()
	
	result, err := box.Eval(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)
}
