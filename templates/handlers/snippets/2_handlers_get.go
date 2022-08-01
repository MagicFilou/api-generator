package snippets

import (
	"api-builder/templates"
)

const HandlersGet templates.Template = `
func Get(ID string, preload bool) ({{ .Name.Short }} {{ .Name.Singular }}.{{ .Name.CamelUpper }}, err error) {

	if ID == "0" {
		return {{ .Name.Short }}, fmt.Errorf("id out of bounds")
	}

  {{ .Name.Short }}.ID = ID

  err = {{ .Name.Short }}.Get(preload)

	return {{ .Name.Short }}, err
}
`
