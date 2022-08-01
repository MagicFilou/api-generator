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
	// "time"

	cfg "generator-gw/configs"
	l "generator-gw/logger"
	{{ .Name.Singular }}handler "generator-gw/handlers/{{ .Name.Singular }}"
  {{ .Name.Singular }}model "generator-gw/models/{{ .Name.Singular }}"

	mw "git.wult.io/wult/libmiddleware/middleware/ginmw"

	// "github.com/chenyahui/gin-cache"
	// "github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
	// "github.com/go-redis/redis/v8"
)
`
