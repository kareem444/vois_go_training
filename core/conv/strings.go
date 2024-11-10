package conv

import "fmt"

func ToString(value any) string {
	if value == nil {
		return ""
	}

	return fmt.Sprintf("%v", value)
}
