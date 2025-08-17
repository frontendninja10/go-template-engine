# Go Template Engine
A lightweight templating engine written in Go.
Supports variables, loops, conditionals, and custom functions.

## Features
- Render templates with `{{ variable }}` syntax
- Conditional blocks: `{{if ...}} ... {{else}} ... {{end}}`
- Loops: ``{{range ...}} ... {{end}}``
- File-based templates
- Extendable with custom functions

## Getting Started

### ⬇️ Install
```bash
go get github.com/frontendninja10/go-template-engine
```

### 🛠️ Usage
```go
package main

import (
    "fmt"
    "github.com/frontendninja10/go-template-engine/pkg/template"
)

func main() {
    tmpl := "Hello, {{ Name }}!"
    data := map[string]string{
        "Name": "Frontend Ninja",
    }

    output, _ := template.Render(tmpl, data)
    fmt.Println(output)
}
```

### Output
```
Hello, Frontend Ninja!
```

### 🛠️ CLI Usage
Render a template with data from JSON:
```bash
gtpl examples/hello.tpl examples/data.json
```

#### Example Template
```
Hello, {{ Name }}!

{{if IsAdmin}}
Welcome back, admin!
{{else}}
Welcome back, user!
{{end}}

Actions:
{{range Actions}}
- {{.}}
{{end}}
```

#### Example Data
```json
{
    "Name": "Frontend Ninja",
    "IsAdmin": true,
    "Actions": [
        "Create",
        "Read",
        "Update",
        "Delete"
    ]
}
```

#### Output
```
Hello, Frontend Ninja!

Welcome back, admin!

Actions:
- Create
- Read
- Update
- Delete
```

### ✅ To-Do
- [ ] Write tests
- [ ] Add more examples
- [ ] Add more documentation
- [ ] Build a playground frontend

### 🤝 Personal Notes
This section is for my personal notes, thought process, and decisions and is not part of the documentation.

#### Template package
After initializing the project, I proceed by creaing `pkg` and `pkg/template` directories, and a `render.go` file. This would contain logic to render output and would serve as a public API.

In the `render.go` file, I wrote the function signature of the `Render` function:
```go
func Render(tmpl string, data map[string]string) (string, error) {
    return "", nil
}
```

#### Examples
This folder would contain ready-to-run demos.

The `examples` directory contains a `hello.tpl` and `data.json` for now.








