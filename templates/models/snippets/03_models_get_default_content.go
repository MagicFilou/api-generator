package snippets

import (
	"api-builder/templates"
)

const ModelsGetDefaultContent templates.Template = `

func ({{ .Name.Short }} *{{ .Name.CamelUpper }}) GetDefaultContent(preload bool) (bool, error) {

	db, err := models.GetConDefaultContent()
	if err != nil {
		return false, err
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
