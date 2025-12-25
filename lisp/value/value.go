package value

import (
	"fmt"

	"github.com/Luke256/mylisp/parser"
)

type Valuer interface {
	String() string
}

type SFunction func(args []Valuer) (Valuer, error)

// Unit : () -------------------------------
type Unit struct {
}

func (u *Unit) String() string {
	return "()"
}

// Number : number -------------------------------

type Number struct {
	Value int32
}

func (n *Number) String() string {
	return fmt.Sprintf("%d", n.Value)
}

// String : string -------------------------------

type String struct {
	Value string
}

func (s *String) String() string {
	return fmt.Sprintf("%q", s.Value)
}

// Function : callable function -------------------------------

type Function struct {
	Expr parser.Exprer
	Args []string
}

func (f *Function) String() string {
	return "<function>"
}

// BuiltinFunction : built-in function -------------------------------

type BuiltinFunction struct {
	Func func(args []Valuer) (Valuer, error)
}

func (bf *BuiltinFunction) String() string {
	return "<builtin-function>"
}

// KeyWord : special keyword -------------------------------

type KeyWord struct {
	Name string
}

func (kw *KeyWord) String() string {
	return fmt.Sprintf("<keyword:%s>", kw.Name)
}

// List : list

type List struct {
	A Valuer
	B Valuer
}

func (l *List) String() string {
	switch l.B.(type) {
	case *Unit:
		return fmt.Sprintf("(%s)", l.A.String())
	case *List:
		var b string = l.B.String()
		return fmt.Sprintf("(%s %s)", l.A.String(), b[1:len(b)-1])
	default:
		return fmt.Sprintf("(%s . %s)", l.A.String(), l.B.String())
	}
}