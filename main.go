package main

import (
	"fmt"
	"github.com/Luke256/mylisp/lisp"
)

func main() {
	var input = `
(define (fact n) (if (= n 0) 1 (* n (fact (- n 1)))))
(define (fib n) (if (or (= n 0) (= n 1)) n (+ (fib (- n 1)) (fib (- n 2)))))
(fib 35)
`

	box := lisp.NewBox()
	
	result, err := box.Eval(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)
}
