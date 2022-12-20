package parser

import (
	"strings"
	"testing"

	"github.com/maniartech/uexl_go/ast"
)

func TestPipe(t *testing.T) {
	got, _ := ParseReader("", strings.NewReader("[1, 2, 3]|map:(1 + 2)"))
	gotNode := got.(ast.PipeNode)

	gotNodeType := gotNode.Type
	if gotNodeType != "pipe" {
		t.Errorf("Token: Expected pipe, got %v", gotNodeType)
	}

	gotNodeToken := gotNode.Token
	if gotNodeToken != "[1, 2, 3]|map:(1 + 2)" {
		t.Errorf("Token: Expected [1, 2, 3]|map:(1 + 2), got %v", gotNodeToken)
	}

	gotNodePipeType := gotNode.PipeType
	if gotNodePipeType[0] != "map" {
		t.Errorf("Pipe Type: Expected map, got %v", gotNodePipeType)
	}

	gotNodeExpressions1 := gotNode.Expressions[0].(ast.ArrayNode)
	gotNodeExpressionsToken1 := gotNodeExpressions1.Token
	if gotNodeExpressionsToken1 != "[1, 2, 3]" {
		t.Errorf("Token: Expected [1, 2, 3], got %v", gotNodeExpressionsToken1)
	}

	gotNodeExpressions1_1 := gotNodeExpressions1.Value[0].(ast.NumberNode).Token
	gotNodeExpressions1_1Value := gotNodeExpressions1.Value[0].(ast.NumberNode).Value
	if gotNodeExpressions1_1 != "1" {
		t.Errorf("Token: Expected 1, got %v", gotNodeExpressions1_1)
	}
	if gotNodeExpressions1_1Value != 1 {
		t.Errorf("Value: Expected 1, got %v", gotNodeExpressions1_1Value)
	}
	gotNodeExpressions1_2 := gotNodeExpressions1.Value[1].(ast.NumberNode).Token
	gotNodeExpressions1_2Value := gotNodeExpressions1.Value[1].(ast.NumberNode).Value
	if gotNodeExpressions1_2 != "2" {
		t.Errorf("Token: Expected 2, got %v", gotNodeExpressions1_1)
	}
	if gotNodeExpressions1_2Value != 2 {
		t.Errorf("Value: Expected 2, got %v", gotNodeExpressions1_1Value)
	}

	gotNodeExpressions1_3 := gotNodeExpressions1.Value[2].(ast.NumberNode).Token
	gotNodeExpressions1_3Value := gotNodeExpressions1.Value[2].(ast.NumberNode).Value
	if gotNodeExpressions1_3 != "3" {
		t.Errorf("Token: Expected 3, got %v", gotNodeExpressions1_1)
	}
	if gotNodeExpressions1_3Value != 3 {
		t.Errorf("Value: Expected 3, got %v", gotNodeExpressions1_1Value)
	}

	gotNodeExpressions2 := gotNode.Expressions[1].(ast.ExpressionNode)
	gotNodeExpressions2Token := gotNodeExpressions2.Token
	op := gotNodeExpressions2.Operator
	if op != "+" {
		t.Errorf("Operator: Expected +, got %v", op)
	}
	if gotNodeExpressions2Token != "1 + 2" {
		t.Errorf("Token: Expected 1 + 2, got %v", gotNodeExpressions2Token)
	}

	gotNodeExpressions2_1 := gotNodeExpressions2.Left.(ast.NumberNode).Token
	gotNodeExpressions2_1Value := gotNodeExpressions2.Left.(ast.NumberNode).Value
	if gotNodeExpressions2_1 != "1" {
		t.Errorf("Token: Expected 1, got %v", gotNodeExpressions2_1)
	}
	if gotNodeExpressions2_1Value != 1 {
		t.Errorf("Value: Expected 1, got %v", gotNodeExpressions2_1Value)
	}

	gotNodeExpressions2_2 := gotNodeExpressions2.Right.(ast.NumberNode).Token
	gotNodeExpressions2_2Value := gotNodeExpressions2.Right.(ast.NumberNode).Value
	if gotNodeExpressions2_2 != "2" {
		t.Errorf("Token: Expected 2, got %v", gotNodeExpressions2_2)
	}
	if gotNodeExpressions2_2Value != 2 {
		t.Errorf("Value: Expected 2, got %v", gotNodeExpressions2_2Value)
	}

	got, _ = ParseReader("", strings.NewReader("(25 * 4)|:'sum'"))
	gotNode = got.(ast.PipeNode)
	gotNodeType = gotNode.Type
	if gotNodeType != "pipe" {
		t.Errorf("Token: Expected pipe, got %v", gotNodeType)
	}

	gotNodeToken = gotNode.Token
	if gotNodeToken != "(25 * 4)|:'sum'" {
		t.Errorf("Token: Expected (25 * 4)|:'sum', got %v", gotNodeToken)
	}

	gotNodePipeType = gotNode.PipeType
	if gotNodePipeType[0] != "pipe" {
		t.Errorf("Pipe Type: Expected pipe, got %v", gotNodePipeType)
	}

	gotNodeExpressions1Exp := gotNode.Expressions[0].(ast.ExpressionNode)
	gotNodeExpressionsToken1 = gotNodeExpressions1Exp.Token
	op = gotNodeExpressions1Exp.Operator
	if op != "*" {
		t.Errorf("Operator: Expected *, got %v", op)
	}
	if gotNodeExpressionsToken1 != "25 * 4" {
		t.Errorf("Token: Expected 25 * 4, got %v", gotNodeExpressionsToken1)
	}

	num1 := gotNodeExpressions1Exp.Left.(ast.NumberNode).Value
	num1Token := gotNodeExpressions1Exp.Left.(ast.NumberNode).Token
	if num1Token != "25" {
		t.Errorf("Token: Expected 25, got %v", num1Token)
	}
	if num1 != 25 {
		t.Errorf("Value: Expected 25, got %v", num1)
	}
	num2 := gotNodeExpressions1Exp.Right.(ast.NumberNode).Value
	num2Token := gotNodeExpressions1Exp.Right.(ast.NumberNode).Token
	if num2Token != "4" {
		t.Errorf("Token: Expected 4, got %v", num1Token)
	}
	if num2 != 4 {
		t.Errorf("Value: Expected 4, got %v", num1)
	}

	got, _ = ParseReader("", strings.NewReader("{'name': 'abc', 'age': 30}|filter:true"))
	gotNode = got.(ast.PipeNode)
	gotNodeType = gotNode.Type
	if gotNodeType != "pipe" {
		t.Errorf("Token: Expected pipe, got %v", gotNodeType)
	}

	gotNodeToken = gotNode.Token
	if gotNodeToken != "{'name': 'abc', 'age': 30}|filter:true" {
		t.Errorf("Token: Expected {'name': 'abc', 'age': 30}|filter:true, got %v", gotNodeToken)
	}

	gotNodePipeType = gotNode.PipeType
	if gotNodePipeType[0] != "filter" {
		t.Errorf("Pipe Type: Expected filter, got %v", gotNodePipeType)
	}

	gotNodeExpressions1Obj := gotNode.Expressions[0].(ast.ObjectNode)
	gotNodeExpressionsToken1 = gotNodeExpressions1Obj.Token
	if gotNodeExpressionsToken1 != "{'name': 'abc', 'age': 30}" {
		t.Errorf("Token: Expected {'name': 'abc', 'age': 30}, got %v", gotNodeExpressionsToken1)
	}

	keys := gotNodeExpressions1Obj.Items
	key1 := keys[0].Key
	if key1 != "name" {
		t.Errorf("Key: Expected name, got %v", key1)
	}
	value1 := keys[0].Value.(ast.StringNode).Value
	if value1 != "abc" {
		t.Errorf("Value: Expected abc, got %v", value1)
	}

	key2 := keys[1].Key
	if key2 != "age" {
		t.Errorf("Key: Expected age, got %v", key2)
	}
	value2 := keys[1].Value.(ast.NumberNode).Value
	if value2 != 30 {
		t.Errorf("Value: Expected 30, got %v", value2)
	}

	got, _ = ParseReader("", strings.NewReader("[1, 3, 5]|map:(0 / 1)|empty:null"))
	gotNode = got.(ast.PipeNode)

	gotNodeType = gotNode.Type
	if gotNodeType != "pipe" {
		t.Errorf("Token: Expected pipe, got %v", gotNodeType)
	}

	gotNodeToken = gotNode.Token
	if gotNodeToken != "[1, 3, 5]|map:(0 / 1)|empty:null" {
		t.Errorf("Token: Expected [1, 3, 5]|map:(0 / 1)|empty:null, got %v", gotNodeToken)
	}

	gotNodePipeType = gotNode.PipeType
	if gotNodePipeType[0] != "map" {
		t.Errorf("Pipe Type 1: Expected map, got %v", gotNodePipeType[0])
	}
	if gotNodePipeType[1] != "empty" {
		t.Errorf("Pipe Type 2: Expected empty, got %v", gotNodePipeType[1])
	}

	gotNodeExpressions1 = gotNode.Expressions[0].(ast.ArrayNode)
	gotNodeExpressionsToken1 = gotNodeExpressions1.Token
	if gotNodeExpressionsToken1 != "[1, 3, 5]" {
		t.Errorf("Token: Expected [1, 3, 5], got %v", gotNodeExpressionsToken1)
	}

	gotNodeExpressions1_1 = gotNodeExpressions1.Value[0].(ast.NumberNode).Token
	gotNodeExpressions1_1Value = gotNodeExpressions1.Value[0].(ast.NumberNode).Value
	if gotNodeExpressions1_1 != "1" {
		t.Errorf("Token: Expected 1, got %v", gotNodeExpressions1_1)
	}
	if gotNodeExpressions1_1Value != 1 {
		t.Errorf("Value: Expected 1, got %v", gotNodeExpressions1_1Value)
	}
	gotNodeExpressions1_2 = gotNodeExpressions1.Value[1].(ast.NumberNode).Token
	gotNodeExpressions1_2Value = gotNodeExpressions1.Value[1].(ast.NumberNode).Value
	if gotNodeExpressions1_2 != "3" {
		t.Errorf("Token: Expected 3, got %v", gotNodeExpressions1_1)
	}
	if gotNodeExpressions1_2Value != 3 {
		t.Errorf("Value: Expected 3, got %v", gotNodeExpressions1_1Value)
	}

	gotNodeExpressions1_3 = gotNodeExpressions1.Value[2].(ast.NumberNode).Token
	gotNodeExpressions1_3Value = gotNodeExpressions1.Value[2].(ast.NumberNode).Value
	if gotNodeExpressions1_3 != "5" {
		t.Errorf("Token: Expected 5, got %v", gotNodeExpressions1_1)
	}
	if gotNodeExpressions1_3Value != 5 {
		t.Errorf("Value: Expected 5, got %v", gotNodeExpressions1_1Value)
	}

	gotNodeExpressions2 = gotNode.Expressions[1].(ast.ExpressionNode)
	gotNodeExpressions2Token = gotNodeExpressions2.Token
	op = gotNodeExpressions2.Operator
	if op != "/" {
		t.Errorf("Operator: Expected /, got %v", op)
	}
	if gotNodeExpressions2Token != "0 / 1" {
		t.Errorf("Token: Expected 0 / 1, got %v", gotNodeExpressions2Token)
	}

	gotNodeExpressions2_1 = gotNodeExpressions2.Left.(ast.NumberNode).Token
	gotNodeExpressions2_1Value = gotNodeExpressions2.Left.(ast.NumberNode).Value
	if gotNodeExpressions2_1 != "0" {
		t.Errorf("Token: Expected 0, got %v", gotNodeExpressions2_1)
	}
	if gotNodeExpressions2_1Value != 0 {
		t.Errorf("Value: Expected 0, got %v", gotNodeExpressions2_1Value)
	}

	gotNodeExpressions2_2 = gotNodeExpressions2.Right.(ast.NumberNode).Token
	gotNodeExpressions2_2Value = gotNodeExpressions2.Right.(ast.NumberNode).Value
	if gotNodeExpressions2_2 != "1" {
		t.Errorf("Token: Expected 1, got %v", gotNodeExpressions2_2)
	}
	if gotNodeExpressions2_2Value != 1 {
		t.Errorf("Value: Expected 1, got %v", gotNodeExpressions2_2Value)
	}

	gotNodeExpressions3 := gotNode.Expressions[2].(ast.NullNode)
	gotNodeExpressions3Token := gotNodeExpressions3.Token
	if gotNodeExpressions3Token != "null" {
		t.Errorf("Token: Expected null, got %v", gotNodeExpressions3Token)
	}

	got, _ = ParseReader("", strings.NewReader("'a'|x:'b'|:'c'|y:'d'"))
	gotNode = got.(ast.PipeNode)

	gotNodeType = gotNode.Type
	if gotNodeType != "pipe" {
		t.Errorf("Token: Expected pipe, got %v", gotNodeType)
	}

	gotNodeToken = gotNode.Token
	if gotNodeToken != "'a'|x:'b'|:'c'|y:'d'" {
		t.Errorf("Token: Expected 'a'|x:'b'|:'c'|y:'d', got %v", gotNodeToken)
	}

	gotNodePipeType = gotNode.PipeType
	if gotNodePipeType[0] != "x" {
		t.Errorf("Pipe Type 1: Expected x, got %v", gotNodePipeType[0])
	}
	if gotNodePipeType[1] != "pipe" {
		t.Errorf("Pipe Type 2: Expected pipe, got %v", gotNodePipeType[1])
	}
	if gotNodePipeType[2] != "y" {
		t.Errorf("Pipe Type 3: Expected y, got %v", gotNodePipeType[2])
	}

	strings := gotNode.Expressions
	str1 := strings[0].(ast.StringNode).Value
	str1Token := strings[0].(ast.StringNode).Token
	if str1Token != "'a'" {
		t.Errorf("Token: Expected a, got %v", str1Token)
	}
	if str1 != "a" {
		t.Errorf("Value: Expected a, got %v", str1)
	}
	str2 := strings[1].(ast.StringNode).Value
	str2Token := strings[1].(ast.StringNode).Token
	if str2Token != "'b'" {
		t.Errorf("Token: Expected b, got %v", str2Token)
	}
	if str2 != "b" {
		t.Errorf("Value: Expected b, got %v", str2)
	}
	str3 := strings[2].(ast.StringNode).Value
	str3Token := strings[2].(ast.StringNode).Token
	if str3Token != "'c'" {
		t.Errorf("Token: Expected c, got %v", str3Token)
	}
	if str3 != "c" {
		t.Errorf("Value: Expected c, got %v", str3)
	}
	str4 := strings[3].(ast.StringNode).Value
	str4Token := strings[3].(ast.StringNode).Token
	if str4Token != "'d'" {
		t.Errorf("Token: Expected d, got %v", str4Token)
	}
	if str4 != "d" {
		t.Errorf("Value: Expected d, got %v", str4)
	}
}
