package snippets

import (
	"api-builder/templates"
)

const RoutesCount templates.Template = `
  {{ .Name.CamelLower }}Routes.GET("/count",
      func(c *gin.Context) {

			distinct := c.DefaultQuery("distinct", "")
			group := c.DefaultQuery("group", "")

			counts, err := {{ .Name.Singular }}handler.Count(distinct, group)
			if err != nil {
				c.AbortWithStatusJSON(checkError(err))
				return
			}

			c.JSON(cfg.CODE_SUCCESS, gin.H{
				"status": cfg.STATUS_SUCCESS,
				"data": counts,
			})
		})

  `
