package snippets

import (
	"api-builder/templates"
)

const HandlersDelete templates.Template = `
  func Delete(ID int32, conKey int) error {

	if ID == 0 {
		return fmt.Errorf("id out of bounds")
	}

  var {{ .Name.Short }} {{ .Name.Singular }}.{{ .Name.CamelUpper }}

  {{ .Name.Short }}.ID = ID

  return {{ .Name.Short }}.Delete(conKey)
  }
`
