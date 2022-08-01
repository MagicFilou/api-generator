package snippets

import (
	"api-builder/templates"
)

const RoutesDelete templates.Template = `
  {{ .Name.CamelLower }}Routes.DELETE("/:id",
      func(c *gin.Context) {

			ID := c.Param("id")

			err := {{ .Name.Singular }}handler.Delete(ID)
			if err != nil {
				c.AbortWithStatusJSON(checkError(err))
				return
			}

			c.JSON(cfg.CODE_SUCCESS, gin.H{
				"status": cfg.STATUS_SUCCESS,
			})
		})
`
