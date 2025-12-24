package expression

import "fmt"

type StringExpr struct {
	Value string
}

func (StringExpr) Express() {}

func (s *StringExpr) String() string {
	return fmt.Sprintf("%q", s.Value)
}