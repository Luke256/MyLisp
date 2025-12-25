package parser

import (
	"fmt"
	"strings"
)

type Exprer interface {
	exprNode()
}

// -------------------------------

type Number struct {
	Value int32
}

func (n *Number) exprNode() {}

func (n *Number) String() string {
	return fmt.Sprintf("%d", n.Value)
}

// -------------------------------

type String struct {
	Value string
}

func (s *String) exprNode() {}

func (s *String) String() string {
	return fmt.Sprintf("%q", s.Value)
}

// -------------------------------

type Symbol struct {
	Name string
}

func (s *Symbol) exprNode() {}

func (s *Symbol) String() string {
	return s.Name
}

// -------------------------------

type List struct {
	// Func Exprer
	// Args []Exprer
	Exprs []Exprer
}

func (c *List) exprNode() {}

func (c *List) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	for i, arg := range c.Exprs {
		if i > 0 {
			sb.WriteString(" ")
		}
		sb.WriteString(fmt.Sprintf("%v", arg))
	}
	sb.WriteString(")")
	return sb.String()
}

// -------------------------------

type Boolean struct {
	Value bool
}

func (b *Boolean) exprNode() {}

func (b *Boolean) String() string {
	if b.Value {
		return "#t"
	}
	return "#f"
}