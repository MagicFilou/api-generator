package snippets

import (
	"api-builder/templates"
)

const RoutesDeclarationGroup templates.Template = `
  {{ .Name.CamelLower }}Routes := r.Group("/{{ .ParentFolder }}/{{ .Name.Plural }}", mw.AuthMW("{{ .SecretType }}", l.GetLogger()))
	{
`
