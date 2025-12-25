package main

import (
	"fmt"
	"github.com/Luke256/mylisp/lisp"
)

func main() {
	var input = `
(define x (list 1 2 3 4 5))
(car (cdr (cdr x)))
`

	box := lisp.NewBox()
	
	result, err := box.Eval(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)
}
