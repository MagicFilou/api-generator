package snippets

import (
	"api-builder/templates"
)

const ModelsGet templates.Template = `

func ({{ .Name.Short }} *{{ .Name.CamelUpper }}) Get(preload bool) error {

	db, err := models.GetCon()
	if err != nil {
		return err
	}

  {{ if .Preload }}
	var result *gorm.DB

	if preload {
		result = db{{ .Preload }}.Find(&{{ .Name.Short }})
	} else {
		result = db.Find(&{{ .Name.Short }})
	}
  {{ else }}
    result := db.Find(&{{ .Name.Short }})
  {{ end }}

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
