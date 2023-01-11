package ast

import (
	"fmt"

	"github.com/maniartech/uexl_go/types"
)

type IdentifierNode struct {
	BaseNode

	Name string `json:"name"`
}

func NewIdentifierNode(token string, offset, line, col int) (*IdentifierNode, error) {
	node := &IdentifierNode{
		BaseNode: BaseNode{
			Type:   NodeTypeIdentifier,
			Line:   line,
			Offset: offset,
			Column: col,
			Token:  token,
		},
		Name: token,
	}

	return node, nil
}

func (n IdentifierNode) Eval(m types.Map) (any, error) {
	// If m is a nil, return an error
	if m == nil {
		return nil, fmt.Errorf("cannot access identifier '%s' from nil map", n.Name)
	}

	return m.Get(n.Name)
}
