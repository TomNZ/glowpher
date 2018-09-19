package variable

type String interface {
	Param
	Value() string
}

type StringLiteral string

func (s StringLiteral) Value() string {
	return string(s)
}
