package builtin

import (
	"fmt"
	"strings"

	"github.com/Luke256/mylisp/lisp/value"
)

func Println(args []value.Valuer) (value.Valuer, error) {
	var sb strings.Builder
	for i, arg := range args {
		if i > 0 {
			sb.WriteString(" ")
		}
		sb.WriteString(arg.String())
	}
	fmt.Println(sb.String())
	return &value.Unit{}, nil
}