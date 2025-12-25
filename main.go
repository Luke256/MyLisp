package main

import (
	"fmt"
	"github.com/Luke256/mylisp/lisp"
)

func main() {
	var input = `
(let ((x 10) (y (+ x 20)))
  (+ x y))
`

	box := lisp.NewBox()
	
	result, err := box.Eval(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)
}
