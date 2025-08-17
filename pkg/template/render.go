package template

import (
	"regexp"
)

func Render(tmpl string, data map[string]string) (string, error) {
	pattern := `{{\s*([a-zA-Z]+)\s*}}`

	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}

	out := re.ReplaceAllStringFunc(tmpl, func(m string) string {
		sub := re.FindStringSubmatch(m)
		if len(sub) == 2 {
			key := sub[1]
			if val, ok := data[key]; ok {
				return val
			}
		}

		return m
	})

	return out, nil
}