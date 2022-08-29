package snippets

import (
	"api-builder/templates"
)

const HandlersGetAll templates.Template = `
func GetAll(wds []models.WhereData, conKey int) ([]{{ .Name.Singular }}.{{ .Name.CamelUpper }}, error) {

	return {{ .Name.Singular }}.GetAll(wds, conKey)
}
`
