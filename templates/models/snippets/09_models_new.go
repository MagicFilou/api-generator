package snippets

import (
	"api-builder/templates"
)

const ModelsNew templates.Template = `
func ({{ .Name.Short }} *{{ .Name.CamelUpper }}) New() error {

	db, err := clients.GetCon()
	if err != nil {
		return err
	}

  {{ .Name.Short }}.Created = int(time.Now().Unix())

	result := db.Create(&{{ .Name.Short }})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
`
