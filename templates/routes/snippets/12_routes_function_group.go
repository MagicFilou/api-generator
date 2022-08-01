package snippets

import (
	"api-builder/templates"
)

const RoutesFunctionGroup templates.Template = `
func {{ .Name.CamelUpper }}Group(r *gin.Engine) {
`
