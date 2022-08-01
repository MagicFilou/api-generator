package snippets

import (
	"api-builder/templates"
)

const RoutesGetAll templates.Template = `
  {{ .Name.CamelLower }}Routes.GET("",
			// cache.CacheByRequestPath(redisStore, 2*time.Minute),
			func(c *gin.Context) {

				preload, _ := strconv.ParseBool(c.DefaultQuery("preload", "false"))

  {{ .Name.Plural }}, err := {{ .Name.Singular }}handler.GetAll(preload)
				if err != nil {
					c.AbortWithStatusJSON(checkError(err))
					return
				}

				c.JSON(cfg.CODE_SUCCESS, gin.H{
					"status": cfg.STATUS_SUCCESS,
					"len":    len({{ .Name.Plural }}),
					"data":   {{ .Name.Plural }},
				})
			})
`
