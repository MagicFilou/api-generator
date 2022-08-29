package snippets

import (
	"api-builder/templates"
)

const RoutesPatch templates.Template = `
  {{ .Name.CamelLower }}Routes.PATCH("",
      func(c *gin.Context) {

			var {{ .Name.Singular }} {{ .Name.Singular }}model.{{ .Name.CamelUpper }}

			if err := json.NewDecoder(c.Request.Body).Decode(&{{ .Name.Singular }}); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if {{ .Name.Singular }}.ID == 0 {
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

			err := {{ .Name.Singular }}handler.Patch(&{{ .Name.Singular }}, int(storageID.(int32)))
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
