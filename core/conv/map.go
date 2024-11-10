package conv

import (
	"encoding/json"

	"example.com/test/core/types"
)

func ToMap(value any) map[string]interface{} {
	if value == nil {
		return map[string]interface{}{}
	}

	if types.IsStruct(value) {
		bytes, err := json.Marshal(&value)
		if err != nil {
			return map[string]interface{}{}
		}
		var result map[string]interface{}
		json.Unmarshal(bytes, &result)
		return result
	}

	return value.(map[string]interface{})
}
