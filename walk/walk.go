package walk

import (
	"fmt"
	jt "github.com/sgonzalez-r7/jtools"
)

type Do = func(n Node)

func Json(j interface{}, do Do) {
	n := Node{
		Base: Base{},
		Val: j,
		Type: jt.Type(j),
		Level: 0,
		IsLeaf: false,
	}

	switch j.(type) {
	case map[string]interface{}:
		do(n)
		walkObj(n, do)
	case []interface{}:
		do(n)
		walkArray(n, do)
	default:
		fmt.Printf("Not valid JSON: %T", j)
	}
}

func walkObj(p Node, do Do) {
	for k, v := range p.Val.(map[string]interface{}) {
		n := Node{
			Base: append(p.Base, k),
			Val: v,
			Type: jt.Type(v),
			Level: p.Level + 1,
			IsLeaf: false,
		}

		switch v.(type) {
		case map[string]interface{}:
			do(n)
			walkObj(n, do)
		case []interface{}:
			do(n)
			walkArray(n, do)
		case string, float64, bool, nil:
			n.IsLeaf = true		
			do(n)
		default:
			fmt.Printf("Not valid JSON: %T", v)
		}
	}
}

func walkArray(p Node, do Do) {
	for i, v := range p.Val.([]interface{}) {
		n := Node{
			Base: append(p.Base, i),
			Val: v,
			Type: jt.Type(v),
			Level: p.Level + 1,
			IsLeaf: false,
		}

		switch v.(type) {
		case map[string]interface{}:
			do(n)
			walkObj(n, do)
		case []interface{}:
			do(n)
			walkArray(n, do)
		case string, float64, bool, nil:
			n.IsLeaf = true		
			do(n)
		default:
			fmt.Printf("Not valid JSON: %T", v)
		}
	}
}
