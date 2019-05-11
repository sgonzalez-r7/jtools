package jtools

import (
	"fmt"
)

func IsEmpty(v interface{}) bool {
	switch val := v.(type) {
	case map[string]interface{}:
		return len(val) == 0
	case []interface{}:
		return len(val) == 0
	case string:
	    	return len(val) == 0
	case bool, float64:
		return false
	case nil:
		return true
	default:
		return false
	}
}

func Type (v interface{}) string {
	switch v.(type) {
	case map[string]interface{}:
		return "[obj]"
	case []interface{}:
		return "[arr]"
	case string:
		return "[str]"
	case float64:
		return "[num]"
	case bool:
		return "[bool]"
	case nil:
		return "[null]"
	default:
		return fmt.Sprintf("%T", v)
	}
}
