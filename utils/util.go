package utils

import (
	"fmt"
	"sort"
	"strings"
)

func Map2String(m map[string]interface{}) string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var msg strings.Builder
	for i, key := range keys {
		if i > 0 {
			msg.WriteString(" ")
		}
		msg.WriteString(fmt.Sprintf("%s=%v", key, m[key]))
	}
	return msg.String()
}
