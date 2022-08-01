package snippets

import (
	"api-builder/templates"
)

const RoutesDeclarationGroup templates.Template = `
  {{ .Name.CamelLower }}Routes := r.Group("/{{ .Name.Plural }}", mw.AuthMW(l.GetLogger()))
	{
`
