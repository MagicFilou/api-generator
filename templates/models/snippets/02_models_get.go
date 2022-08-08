package snippets

import (
	"api-builder/templates"
)

const ModelsGet templates.Template = `

func ({{ .Name.Short }} *{{ .Name.CamelUpper }}) Get() error {

	db, err := models.GetCon()
	if err != nil {
		return err
	}

	var result *gorm.DB

	result = db.Find(&{{ .Name.Short }})

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("not found")
		}
		return err
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("not found")
	}

	return nil
}
`
