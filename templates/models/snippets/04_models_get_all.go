package snippets

import (
	"api-builder/templates"
)

const ModelsGetAll templates.Template = `
func GetAll() ({{ .Name.Plural }} []{{ .Name.CamelUpper }}, err error) {

	db, err := models.GetCon()
	if err != nil {
  return {{ .Name.Plural }}, err
	}

	var result *gorm.DB

	result = db.Find(&{{ .Name.Plural }})

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
  return {{ .Name.Plural }}, fmt.Errorf("not found")
		}
  return {{ .Name.Plural }}, err
	}

	if result.RowsAffected == 0 {
  return {{ .Name.Plural }}, fmt.Errorf("not found")
	}

  return {{ .Name.Plural }}, nil
}
`
