package j

import "fmt"

type path struct {
	list []interface{}
	sep  string
}

func (p path) Value() []interface{} {
	return p.list
}

func (p path) String() (s string) {
	for _, index := range p.list {
		switch i := index.(type) {
		case string:
			s = s + p.sep + i
		case int:
			s = s + p.sep + fmt.Sprintf("%d", i)
		}
	}
	return s
}

func (p path) Join(sep string) path {
	p.sep = sep
	return p
}
