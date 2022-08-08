package snippets

import (
	"api-builder/templates"
)

const ModelsGetAllDefaultContent templates.Template = `
func GetAllDefaultContent() ({{ .Name.Plural }} []{{ .Name.CamelUpper }}, err error) {

	db, err := models.GetConDefaultContent()
	if err != nil {
  return {{ .Name.Plural }}, err
	}

	var result *gorm.DB

	result = db.Find(&{{ .Name.Plural }})

	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
      return {{ .Name.Plural }}, result.Error 
		}
	}

  return {{ .Name.Plural }}, nil
}
`
