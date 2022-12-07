package ast

type OperatorType uint8

const (

	// ArithmeticOperator evaluates the expression by applying arithmatic operations
	// on the oparands.
	ArithmeticOperator OperatorType = iota // + - * / //

	// ComparisonOperator evaluates the expression by comparting the operands.
	ComparisonOperator // = <> < > <= >=

	// LogicalOperator evaluates the expression by applying logical operations
	LogicalOperator // AND OR

	// BitwiseOperator evaluates the expression by applying bitwise operations
	BitwiseOperator // & | ^ << >>

)

func (ot OperatorType) String() string {
	switch ot {
	case ArithmeticOperator:
		return "ArithmeticOperator"

	case ComparisonOperator:
		return "ComparisonOperator"

	case LogicalOperator:
		return "LogicalOperator"

	case BitwiseOperator:
		return "BitwiseOperator"

	default:
		return ErrInvalidOperator
	}
}

func getOperatorType(op string) OperatorType {
	switch op {
	case "+", "-", "*", "/", "//":
		return ArithmeticOperator

	case "=", "<>", "<", ">", "<=", ">=":
		return ComparisonOperator

	case "AND", "OR":
		return LogicalOperator

	case "&", "|", "^", "<<", ">>":
		return BitwiseOperator

	default:
		return 0
	}
}
