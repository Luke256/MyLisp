package main

import (
	"fmt"
	"github.com/Luke256/mylisp/lisp"
)

func main() {
	var input = `
(define a 10)
(define b 10)
(if (= a b) (println "a and b are equal") (println "a and b are not equal"))
`

	box := lisp.NewBox()
	
	result, err := box.Eval(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)
}
