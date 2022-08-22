package snippets

import "api-builder/templates"

const RoutesFunctionCollectGroups templates.Template = `
package routes

import (
  {{ range .Imports }}
    {{ . }}
  {{ end }}

	"github.com/gin-gonic/gin"
)

func CollectGroups(r *gin.Engine) {

	documentation(r)

  {{ range .Templates.Functions }}
    {{ .Start }}
  {{ end }}
}

func documentation(r *gin.Engine) {

	//r.StaticFile("/docs/{{ .ParentFolder }}", "./" + cfg.GetConfig().Documentation.Dir + "/{{ .ParentFolder }}/" + cfg.GetConfig().Documentation.HTML)
  docRoutes.StaticFile("/{{ .ParentFolder }}", "./"+cfg.GetConfig().Documentation.Dir+"/{{ .ParentFolder }}/"+cfg.GetConfig().Documentation.HTML)
}
`
