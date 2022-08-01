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
  {{ range .Imports }}
    {{ . }}
  {{ end }}

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)
`
