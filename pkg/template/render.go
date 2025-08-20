package template

import (
	"strings"
)


func renderConditionals(tmpl string, data map[string]string) (string, error) {
	open := "{{ if "
	close := "}}"
	for {
		ifBlockStart := strings.Index(tmpl, open)
		if ifBlockStart == -1 {
			break
		}

		keyStart := ifBlockStart + len(open)
		ifBlockEndIndex := strings.Index(tmpl[keyStart:], close)
		if ifBlockEndIndex == -1 {
			break
		}
		ifBlockEndIndex = ifBlockEndIndex + keyStart
		key := strings.TrimSpace(tmpl[keyStart:ifBlockEndIndex])

		trueBlock := ifBlockEndIndex + len(close)
		

		elseMarker := "{{ else }}"
		endMarker := "{{ end }}"

		elseCondIndex := strings.Index(tmpl[trueBlock:], elseMarker)
		if elseCondIndex == -1 {
			break
		}
		elseCondIndex = elseCondIndex + trueBlock

		elseBlock := elseCondIndex + len(elseMarker)

		endCondIndex := strings.Index(tmpl[elseBlock:], endMarker)
		if endCondIndex == -1 {
			break
		}
		endCondIndex = endCondIndex + elseBlock

		trueBranch := tmpl[trueBlock:elseCondIndex]
		falseBranch := tmpl[elseBlock:endCondIndex]

		var replacement string
		if val, exists := data[key]; exists && val == "true" {
			replacement = trueBranch
		} else {
			replacement = falseBranch
		}

		tmpl = tmpl[:ifBlockStart] + replacement + tmpl[endCondIndex+len(endMarker):]
	}
	return tmpl, nil
}

func Render(tmpl string, data map[string]string) (string, error) {
	processed, err := renderConditionals(tmpl, data)
	if err != nil {
		return "", err
	}


	newStartIndex := 0
	for {
		firstIndex := strings.Index(processed[newStartIndex:], "{{")
		if firstIndex == -1 {
			break
		}
		firstIndex += newStartIndex

		end := strings.Index(processed[firstIndex+2:], "}}")
		if end == -1 {
			break
		}
		lastIndex := end + firstIndex + 2

		key := strings.TrimSpace(processed[firstIndex+2:lastIndex])

		val, exists := data[key]
		toInsert := processed[firstIndex:lastIndex+2]
		if exists {
			toInsert = val
		}

		processed = processed[:firstIndex] + toInsert + processed[lastIndex+2:]
		newStartIndex = firstIndex + len(toInsert)
	}
	return processed, nil
}