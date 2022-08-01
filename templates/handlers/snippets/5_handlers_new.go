package snippets

import (
	"api-builder/templates"
)

const HandlersNew templates.Template = `
func New({{ .Name.Short }} *{{ .Name.Singular }}.{{ .Name.CamelUpper }}) error {

	return {{ .Name.Short }}.New()
}
`
