package snippets

import (
	"api-builder/templates"
)

const ModelsPatch templates.Template = `

func ({{ .Name.Short }} *{{ .Name.CamelUpper }}) Patch() error {

	db, err := clients.GetCon()
	if err != nil {
		return err
	}

	result := db.Clauses(clause.Returning{}).Updates(&{{ .Name.Short }})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
`
