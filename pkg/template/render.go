package template

import (
	"strings"
)

func Render(tmpl string, data map[string]string) (string, error) {
	newStartIndex := 0
	for {
		firstIndex := strings.Index(tmpl[newStartIndex:], "{{")
		if firstIndex == -1 {
			break
		}
		firstIndex += newStartIndex

		end := strings.Index(tmpl[firstIndex+2:], "}}")
		if end == -1 {
			break
		}
		lastIndex := end + firstIndex + 2

		key := strings.TrimSpace(tmpl[firstIndex+2:lastIndex])

		val, exists := data[key]
		toInsert := tmpl[firstIndex:lastIndex+2]
		if exists {
			toInsert = val
		}

		tmpl = tmpl[:firstIndex] + toInsert + tmpl[lastIndex+2:]
		newStartIndex = firstIndex + len(toInsert)
	}
	return tmpl, nil
}