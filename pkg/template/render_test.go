package template // Or whatever package your Render function is in

import (
	"testing"
)

// You might need to add other imports if your Render function uses them,
// like "errors" if you had specific error returns earlier that you removed.

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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Render(tt.template, tt.data)
			if err != nil {
				t.Fatalf("Render returned an unexpected error: %v", err)
			}
			if got != tt.expected {
				t.Errorf("Render() got = %q, want %q", got, tt.expected)
			}
		})
	}
}