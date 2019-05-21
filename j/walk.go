package j

import (
	// "log"
)

type Do = func(n Node)

func Walk(j interface{}, do Do) {
	n := Node{
		Branch: branch{j},
		Path: path{[]interface{}{}, defaultSep},
	}
	switch j.(type) {
	case map[string]interface{}:
		do(n)
		walkObj(n, do)
	case []interface{}:
		do(n)
		walkArray(n, do)
	}
}

func walkObj(r Node, do Do) {
	for k, v := range r.Branch.value.(map[string]interface{}) {
		n := Node{
			Branch: branch{v},
			Path: path{append(r.Path.list, k), defaultSep},
		}
		switch v.(type) {
		case map[string]interface{}:
			do(n)
			walkObj(n, do)
		case []interface{}:
			do(n)
			walkArray(n, do)
		case string, float64, bool, nil:
			do(n)
		}
	}
}

func walkArray(r Node, do Do) {
	for i, v := range r.Branch.value.([]interface{}) {
		n := Node{
			Branch: branch{v},
			Path: path{append(r.Path.list, i), defaultSep},
		}
		switch v.(type) {
		case map[string]interface{}:
			do(n)
			walkObj(n, do)
		case []interface{}:
			do(n)
			walkArray(n, do)
		case string, float64, bool, nil:
			do(n)
		}
	}
}
