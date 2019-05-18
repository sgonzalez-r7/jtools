package j

import (
	// "fmt"
)

type Node struct {
	Branch  branch
	Path    path
}

func (n Node) IsLeaf() bool {
	switch n.Branch.value.(type) {
	case string, float64, bool, nil:
		return true
	}
	return false
}


func (n Node) Type() string {
	return Type(n.Branch.value)
}

// func (n Node) String() string {

// }

// type Node struct {
// 	Base Base
// 	Val interface{}
// 	Type string
// 	Level int32
// 	IsLeaf bool
// }
// type Base []interface{}
// func (b Base) Path(sep string) string {
// 	buf := make([]string, 0, len(b))
// 	for _, v := range b {
// 		switch val := v.(type) {
// 		case string:
// 			buf = append(buf, val)
// 		case int:
// 			buf = append(buf, `[`+ strconv.Itoa(val) + `]`)
// 		}
// 	}
// 	return strings.Join(buf, sep)
// }

// func (b Base) Squash(sep string) string {
// 	buf := make([]string, 0, len(b))
// 	for _, v := range b {
// 		switch val := v.(type) {
// 		case string:
// 			buf = append(buf, val)
// 		case int:
// 			buf = append(buf, `[i]`)
// 		}
// 	}
// 	return strings.Join(buf, sep)
// }
