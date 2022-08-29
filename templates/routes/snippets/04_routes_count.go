package snippets

import (
	"api-builder/templates"
)

const RoutesCount templates.Template = `
  {{ .Name.CamelLower }}Routes.GET("/count",
      func(c *gin.Context) {

			distinct := c.DefaultQuery("distinct", "")
			group := c.DefaultQuery("group", "")

			storageID, ok := c.Get(mw.KEY_STORAGE)
				if !ok {
					c.AbortWithStatusJSON(checkError(fmt.Errorf("No storage ID")))
					return
				}

			counts, err := {{ .Name.Singular }}handler.Count(distinct, group, int(storageID.(int32)))
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
