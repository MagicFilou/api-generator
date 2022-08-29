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
	"fmt"

	cfg "generator-gw/configs"
	l "generator-gw/logger"
	{{ .Name.Singular }}handler "generator-gw/handlers/{{ .ParentFolder }}/{{ .Name.Singular }}"
  {{ .Name.Singular }}model "generator-gw/models/{{ .ParentFolder }}/{{ .Name.Singular }}"
	"generator-gw/models"
	"generator-gw/util"

	"generator-gw/mw"
	
	"github.com/gin-gonic/gin"
)
`
