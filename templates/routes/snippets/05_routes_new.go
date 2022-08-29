package snippets

import (
	"api-builder/templates"
)

const RoutesNew templates.Template = `
  {{ .Name.CamelLower }}Routes.POST("",
      func(c *gin.Context) {

			// redisStore.Delete("")

			var {{ .Name.Singular }} {{ .Name.Singular }}model.{{ .Name.CamelUpper }}

			if err := c.ShouldBindJSON(&{{ .Name.Singular }}); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			storageID, ok := c.Get(mw.KEY_STORAGE)
				if !ok {
					c.AbortWithStatusJSON(checkError(fmt.Errorf("No storage ID")))
					return
				}

			err := {{ .Name.Singular }}handler.New(&{{ .Name.Singular }}, int(storageID.(int32)))
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
