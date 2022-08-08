package snippets

import (
	"api-builder/templates"
)

const HandlersGetAll templates.Template = `
func GetAll() ([]{{ .Name.Singular }}.{{ .Name.CamelUpper }}, error) {

	return {{ .Name.Singular }}.GetAll()
}
`
