package walk

import (
	"strconv"
	"strings"
)

type Node struct {
	Base Base
	Val interface{}
	Type string
	Level int32
	IsLeaf bool
}

type Base []interface{}

func (b Base) Path(keep bool) string {
	var k string
	buf := make([]string, 0, len(b))
	for _, v := range b {
		switch val := v.(type) {
		case string:
			k = val
		case int:
			i := `i`
			if keep {
				i = strconv.Itoa(val)
			}
			k = `[`+ i +`]`
		}
		buf = append(buf, k)
	}
	return `/`+ strings.Join(buf, `/`)
}
