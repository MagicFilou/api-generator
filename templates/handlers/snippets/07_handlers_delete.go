package snippets

import (
	"api-builder/templates"
)

const HandlersDelete templates.Template = `
  func Delete(ID string) error {

	if ID == "0" {
		return fmt.Errorf("id out of bounds")
	}

  var {{ .Name.Short }} {{ .Name.Singular }}.{{ .Name.CamelUpper }}

  {{ .Name.Short }}.ID = ID

  return {{ .Name.Short }}.Delete()
  }
`
