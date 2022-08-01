package snippets

import (
	"api-builder/templates"
)

const RoutesGet templates.Template = `
  {{ .Name.CamelLower }}Routes.GET("/:id",
			// cache.CacheByRequestPath(redisStore, 2*time.Minute),
			func(c *gin.Context) {

				ID := c.Param("id")
				preload, _ := strconv.ParseBool(c.DefaultQuery("preload", "false"))

  {{ .Name.Singular }}, err := {{ .Name.Singular }}handler.Get(ID, preload)
				if err != nil {
					c.AbortWithStatusJSON(checkError(err))
					return
				}

				c.JSON(cfg.CODE_SUCCESS, gin.H{
					"status": cfg.STATUS_SUCCESS,
					"data":   {{ .Name.Singular }},
				})
			})
`