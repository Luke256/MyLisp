package expression

import (
	"strings"
)

type ListExpr struct {
	Elements []Exprer
}

func (ListExpr) Express() {}

func (l *ListExpr) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	for i, elem := range l.Elements {
		if i > 0 {
			sb.WriteString(" ")
		}
		sb.WriteString(elem.(interface{ String() string }).String())
	}
	sb.WriteString(")")
	return sb.String()
}
