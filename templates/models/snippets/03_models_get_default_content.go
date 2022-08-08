package snippets

import (
	"api-builder/templates"
)

const ModelsGetDefaultContent templates.Template = `

func ({{ .Name.Short }} *{{ .Name.CamelUpper }}) GetDefaultContent() (bool, error) {

	db, err := models.GetConDefaultContent()
	if err != nil {
		return false, err
	}


	var result *gorm.DB
		
  result := db.Find(&{{ .Name.Short }})

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	if result.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}
`
