package snippets

import (
	"api-builder/templates"
)

const HandlersImports templates.Template = `
import (
	"fmt"
	"generator-gw/models"
	"generator-gw/models/{{ .Name.Singular }}"
)
`
