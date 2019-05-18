package j

import (
	"encoding/json"
	"io"
	// "fmt"
)

func Decode(r io.Reader) (j interface{}, err error) {
	err = json.NewDecoder(r).Decode(&j)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func Type (v interface{}) (t string) {
	switch v.(type) {
	case map[string]interface{}:
		t = "[obj]"
	case []interface{}:
		t = "[arr]"
	case string:
		t = "[str]"
	case float64:
		t = "[num]"
	case bool:
		t = "[bool]"
	case nil:
		t = "[null]"
	default:
		t = "[invalid]"
	}
	return t
}

// func IsEmpty(v interface{}) bool {
// 	switch val := v.(type) {
// 	case map[string]interface{}:
// 		return len(val) == 0
// 	case []interface{}:
// 		return len(val) == 0
// 	case string:
// 	    	return len(val) == 0
// 	case bool, float64:
// 		return false
// 	case nil:
// 		return true
// 	default:
// 		return false
// 	}
// }
