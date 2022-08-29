package snippets

import (
	"api-builder/templates"
)

const HandlersPatch templates.Template = `
func Patch({{ .Name.Short }} *{{ .Name.Singular }}.{{ .Name.CamelUpper }}, conKey int) error {

	return {{ .Name.Short }}.Patch(conKey)
}
`
