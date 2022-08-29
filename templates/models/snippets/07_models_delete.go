package snippets

import (
	"api-builder/templates"
)

const ModelsDelete templates.Template = `

  func ({{ .Name.Short }} {{ .Name.CamelUpper }}) Delete(conKey int) error {

	db, err := clients.GetCon(conKey)
	if err != nil {
		return err
	}

	result := db.Delete(&{{ .Name.Short }})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

`
