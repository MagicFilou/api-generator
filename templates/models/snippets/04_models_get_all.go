package snippets

import (
	"api-builder/templates"
)

const ModelsGetAll templates.Template = `
func GetAll(wds []models.WhereData) ({{ .Name.Plural }} []{{ .Name.CamelUpper }}, err error) {

	db, err := clients.GetCon()
	if err != nil {
  return {{ .Name.Plural }}, err
	}

	var result *gorm.DB

	query, args := models.BuildWHereQuery(wds)
	if len(query) > 0 {
		result = db.Where(query, args...).Find(&{{ .Name.Plural }})
	} else {
	result = db.Find(&{{ .Name.Plural }})
	}

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
