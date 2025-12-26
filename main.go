package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Luke256/mylisp/lisp"
)

func main() {
	box := lisp.NewBox()
	
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("mylisp> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		result, err := box.Eval(line)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Println(result.String())
	}
}
