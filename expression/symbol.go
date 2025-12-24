package expression

type SymbolExpr struct {
	Name string
}

func (SymbolExpr) Express() {}

func (s *SymbolExpr) String() string {
	return s.Name
}