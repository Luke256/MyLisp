package main

import (
	"fmt"
	"github.com/Luke256/mylisp/tokenizer"
	"github.com/Luke256/mylisp/parser"
)

type Hoge interface {
	int | string
}

type Fuga[T Hoge] struct {
	Values []T
}

func main() {
	var input = `(define (square x) (* x x))
(print (square 5))`;

	tokens, err := tokenizer.Tokenize(input)
	if err != nil {
		fmt.Println("Error tokenizing input:", err)
		return
	}


	exprs, err := parser.Parse(tokens)
	if err != nil {
		fmt.Println("Error parsing tokens:", err)
		return
	}

	for _, expr := range exprs {
		fmt.Println(expr)
	}
}