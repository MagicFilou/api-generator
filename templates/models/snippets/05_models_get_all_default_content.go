package snippets

import (
	"api-builder/templates"
)

const ModelsGetAllDefaultContent templates.Template = `
func GetAllDefaultContent(preload bool) ({{ .Name.Plural }} []{{ .Name.CamelUpper }}, err error) {

	db, err := models.GetConDefaultContent()
	if err != nil {
  return {{ .Name.Plural }}, err
	}

  {{ if .Preload }}
	var result *gorm.DB

	if preload {
    result = db{{ .Preload }}.Find(&{{ .Name.Plural }})
	} else {
    result = db.Find(&{{ .Name.Plural }})
	}
  {{ else }}
    result := db.Find(&{{ .Name.Plural }})
  {{ end }}

	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
      return {{ .Name.Plural }}, result.Error 
		}
	}

  return {{ .Name.Plural }}, nil
}
`
