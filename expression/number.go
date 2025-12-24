package expression

import "fmt"

type NumberExpr struct {
	Value int32
}

func (NumberExpr) Express() {}

func (n *NumberExpr) String() string {
	return fmt.Sprintf("%d", n.Value)
}