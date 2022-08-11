package snippets

import (
	"api-builder/templates"
)

const ModelsImports templates.Template = `
import (
	"errors"
	"fmt"
  "time"

	"generator-gw/models"
	"generator-gw/clients"
  {{ range .Imports }}
    {{ . }}
  {{ end }}

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)
`
