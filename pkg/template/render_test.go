package template

import (
	"testing"
)

func TestRenderConditionals(t *testing.T) {
	tests := []struct{
		name     string
		template string
		data     map[string]string
		expected string
	}{
		{
			name:     "Test Case 1: key present and “true”",
			template: "{{ if IsAdmin }}Welcome Admin!{{ else }}Welcome User!{{ end }}",
			data:     map[string]string{"IsAdmin": "true"},
			expected: "Welcome Admin!",
		},
		{
			name:     "Test Case 2: key present and “false”",
			template: "{{ if IsAdmin }}Welcome Admin!{{ else }}Welcome User!{{ end }}",
			data:     map[string]string{"IsAdmin": "false"},
			expected: "Welcome User!",
		},
		{
			name:     "Test Case 3: key NOT present",
			template: "{{ if IsAdmin }}Welcome Admin!{{ else }}Welcome User!{{ end }}",
			data:     map[string]string{},
			expected: "Welcome User!",
		},
		{
			name:     "Test Case 4: More complex content inside branches",
			template: "{{ if IsAdmin }}Hello, {{Name}}!{{ else }}Please log in.{{ end }}",
			data:     map[string]string{"IsAdmin": "true", "Name": "Bo"},
			expected: "Hello, {{Name}}!",
		},
		{
			name:     "Test Case 5: Multiple branches",
			template: `{{ if IsAdmin }}Welcome Admin!{{ else }}Welcome User!{{ end }} Some content...{{ if Verified }}Enjoy full features.{{ else }}Please verify your email.{{ end }}`,
			data : map[string]string{"IsAdmin": "false", "Verified": "true"},
			expected: "Welcome User! Some content...Enjoy full features.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := renderConditionals(tt.template, tt.data)
			if err != nil {
				t.Fatalf("Render returned an unexpected error: %v", err)
			}
			if res != tt.expected {
				t.Errorf("Expected: %q\n Got: %q", tt.expected, res)
			}
		})
	}
}

func TestRender(t *testing.T) {
	tests := []struct {
		name     string
		template string
		data     map[string]string
		expected string
	}{
		{
			name:     "Simple replacement",
			template: "Hello {{ Name }}!",
			data:     map[string]string{"Name": "Alice"},
			expected: "Hello Alice!",
		},
		{
			name:     "Multiple replacements",
			template: "Hello {{ Name }}, welcome to {{ Site }}!",
			data:     map[string]string{"Name": "Bob", "Site": "chelsea.com"},
			expected: "Hello Bob, welcome to chelsea.com!",
		},
		{
			name:     "Key not found",
			template: "Hello {{ Unknown }}!",
			data:     map[string]string{"Name": "Charlie"},
			expected: "Hello {{ Unknown }}!",
		},
		{
			name:     "Empty template",
			template: "",
			data:     map[string]string{"Name": "Dave"},
			expected: "",
		},
		{
			name:     "No placeholders",
			template: "This is a plain string.",
			data:     map[string]string{"Name": "Eve"},
			expected: "This is a plain string.",
		},
		{
			name:     "Whitespace in placeholder",
			template: "Hello {{   User   }}!",
			data:     map[string]string{"User": "Frank"},
			expected: "Hello Frank!",
		},
		{
			name:     "Placeholder at start and end",
			template: "{{Greeting}}, {{Name}}!",
			data:     map[string]string{"Greeting": "Hi", "Name": "Grace"},
			expected: "Hi, Grace!",
		},
        {
            name:     "Malformed template missing closing bracket",
            template: "Hello {{ Name!",
            data:     map[string]string{"Name": "Heidi"},
            expected: "Hello {{ Name!",
        },
        {
            name:     "Malformed template missing opening bracket",
            template: "Hello Name }}!",
            data:     map[string]string{"Name": "Ivan"},
            expected: "Hello Name }}!",
        },
        {
            name:     "Multiple same placeholders",
            template: "{{Name}} and {{Name}} again!",
            data:     map[string]string{"Name": "Julia"},
            expected: "Julia and Julia again!",
        },
		{
			name:     "Placeholder content inside branches",
			template: "{{ if IsAdmin }}Hello, {{Name}}!{{ else }}Please log in.{{ end }}",
			data:     map[string]string{"IsAdmin": "true", "Name": "Bo"},
			expected: "Hello, Bo!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := Render(tt.template, tt.data)
			if res != tt.expected {
				t.Errorf("Render() got = %q, want %q", res, tt.expected)
			}
		})
	}
}