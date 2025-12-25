package value

import (
	"fmt"

	"github.com/Luke256/mylisp/parser"
)

type Valuer interface {
	String() string
	Equal(other Valuer) bool
}

type SFunction func(args []Valuer) (Valuer, error)

// Unit : () -------------------------------
type Unit struct {
}

func (u *Unit) String() string {
	return "()"
}

func (u *Unit) Equal(other Valuer) bool {
	_, ok := other.(*Unit)
	return ok
}

// Number : number -------------------------------

type Number struct {
	Value int32
}

func (n *Number) String() string {
	return fmt.Sprintf("%d", n.Value)
}

func (n *Number) Equal(other Valuer) bool {
	otherNum, ok := other.(*Number)
	if !ok {
		return false
	}
	return n.Value == otherNum.Value
}

// String : string -------------------------------

type String struct {
	Value string
}

func (s *String) String() string {
	return fmt.Sprintf("%q", s.Value)
}

func (s *String) Equal(other Valuer) bool {
	otherStr, ok := other.(*String)
	if !ok {
		return false
	}
	return s.Value == otherStr.Value
}

// Function : callable function -------------------------------

type Function struct {
	Expr parser.Exprer
	Args []string
}

func (f *Function) String() string {
	return "<function>"
}

func (f *Function) Equal(other Valuer) bool {
	return false
}

// BuiltinFunction : built-in function -------------------------------

type BuiltinFunction struct {
	Func func(args []Valuer) (Valuer, error)
}

func (bf *BuiltinFunction) String() string {
	return "<builtin-function>"
}

func (bf *BuiltinFunction) Equal(other Valuer) bool {
	return false
}

// KeyWord : special keyword -------------------------------

type KeyWord struct {
	Name string
}

func (kw *KeyWord) String() string {
	return fmt.Sprintf("<keyword:%s>", kw.Name)
}

func (kw *KeyWord) Equal(other Valuer) bool {
	otherKw, ok := other.(*KeyWord)
	if !ok {
		return false
	}
	return kw.Name == otherKw.Name
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

func (l *List) Equal(other Valuer) bool {
	otherList, ok := other.(*List)
	if !ok {
		return false
	}
	return l.A.Equal(otherList.A) && l.B.Equal(otherList.B)
}

// Bool : boolean -------------------------------

type Boolean struct {
	Value bool
}

func (b *Boolean) String() string {
	if b.Value {
		return "#t"
	}
	return "#f"
}

func (b *Boolean) Equal(other Valuer) bool {
	otherBool, ok := other.(*Boolean)
	if !ok {
		return false
	}
	return b.Value == otherBool.Value
}