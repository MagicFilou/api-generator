package snippets

import (
	"api-builder/templates"
)

const HandlersNew templates.Template = `
func New({{ .Name.Short }} *{{ .Name.Singular }}.{{ .Name.CamelUpper }}, conKey int) error {

	return {{ .Name.Short }}.New(conKey)
}
`
