package snippets

import (
	"api-builder/templates"
)

const RoutesImports templates.Template = `
import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	cfg "generator-gw/configs"
	l "generator-gw/logger"
	{{ .Name.Singular }}handler "generator-gw/handlers/{{ .Name.Singular }}"
  {{ .Name.Singular }}model "generator-gw/models/{{ .Name.Singular }}"

	"generator-gw/mw"
	
	"github.com/gin-gonic/gin"
)
`
