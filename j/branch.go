package j

import "fmt"

type branch struct {
	value interface{}
}

func (b branch) Value() interface{} {
	return b.value
}

func (b branch) String() string {
	switch v := b.value.(type) {
	case map[string]interface{}:
		return "[obj]"
	case []interface{}:
		return "[arr]"
	case string:
	    	return v
	case float64:
		return fmt.Sprintf("%g", v)
	case bool:
		return fmt.Sprintf("%v", v)
	case nil:
		return "[null]"
	}
	return "[invalid]"
}
