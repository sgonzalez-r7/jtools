package j

import (
	"fmt"
	"regexp"
	"strconv"
)

const defaultSep = " "

var reSlice *regexp.Regexp = regexp.MustCompile(`(\d*):(\d*)`)

type path struct {
	list []interface{}
	sep  string
}

func (p path) Value() []interface{} {
	return p.list
}

func (p path) Slice(s string) path {
	field := reSlice.FindStringSubmatch(s)
	si, sj := field[1], field[2]
	i, _ := strconv.Atoi(si)
	j, _ := strconv.Atoi(sj)

	list := make([]interface{}, 0)

	switch {
	case len(si) > 0 && len(sj) > 0:
		list = append(list, p.list[i:j]...)
	case len(si) > 0 && len(sj) == 0:
		list = append(list, p.list[i:]...)
	case len(si) == 0 && len(sj) > 0:
		list = append(list, p.list[:j]...)
	case len(si) == 0 && len(sj) == 0:
		list = append(list, p.list[:]...)
	}

	return path{
		list: list,
		sep: p.sep,
	}
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
