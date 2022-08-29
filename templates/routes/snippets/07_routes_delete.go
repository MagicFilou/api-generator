package snippets

import (
	"api-builder/templates"
)

const RoutesDelete templates.Template = `
  {{ .Name.CamelLower }}Routes.DELETE("/:id",
      func(c *gin.Context) {

			ID := c.Param("id")
			intID, err := strconv.Atoi(ID)

			if err != nil {
				c.AbortWithStatusJSON(422, gin.H{
					"status": "failure",
					"error":  "missing ID",
				})
				return
			}

			storageID, ok := c.Get(mw.KEY_STORAGE)
				if !ok {
					c.AbortWithStatusJSON(checkError(fmt.Errorf("No storage ID")))
					return
				}

			err = {{ .Name.Singular }}handler.Delete(int32(intID), int(storageID.(int32)))

			if err != nil {
				c.AbortWithStatusJSON(checkError(err))
				return
			}

			c.JSON(cfg.CODE_SUCCESS, gin.H{
				"status": cfg.STATUS_SUCCESS,
			})
		})
`
