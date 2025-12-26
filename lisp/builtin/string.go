package builtin

import (
	"fmt"
	"strings"
	
	"github.com/Luke256/mylisp/lisp/value"
)

func Concat(args []value.Valuer) (value.Valuer, error) {
	var sb strings.Builder
	for _, arg := range args {
		strVal, ok := arg.(*value.String)
		if !ok {
			return nil, fmt.Errorf("concat expects string arguments, got %T", arg)
		}
		sb.WriteString(strVal.Value)
	}
	return &value.String{Value: sb.String()}, nil
}

func NumberToString(args []value.Valuer) (value.Valuer, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("number-to-string expects 1 argument, got %d", len(args))
	}
	numVal, ok := args[0].(*value.Number)
	if !ok {
		return nil, fmt.Errorf("number-to-string expects a number argument, got %T", args[0])
	}
	return &value.String{Value: fmt.Sprintf("%v", numVal.Value)}, nil
}