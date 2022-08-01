package snippets

import "api-builder/templates"

const RoutesCustomNestedServiceConfigurations templates.Template = `
		serviceConfigurationRoutes.GET("/nested",
			// cache.CacheByRequestPath(redisStore, 2*time.Minute),
			func(c *gin.Context) {

				serviceconfigurations, err := serviceconfigurationhandler.GetAllNested()
				if err != nil {
					c.AbortWithStatusJSON(checkError(err))
					return
				}

				c.JSON(cfg.CODE_SUCCESS, gin.H{
					"status": cfg.STATUS_SUCCESS,
					"len":    len(serviceconfigurations),
					"data":   serviceconfigurations,
				})
			})
`
