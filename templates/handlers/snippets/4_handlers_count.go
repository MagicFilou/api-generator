package snippets

import (
	"api-builder/templates"
)

const HandlersCount templates.Template = `
func Count(distinct, group string) ([]models.CountData, error) {

	return {{ .Name.Singular }}.Count(distinct, group)
}
`
