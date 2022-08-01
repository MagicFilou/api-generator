package snippets

import (
	"api-builder/templates"
)

const RoutesFunctionCheckError templates.Template = `
func checkError(err error) (int, gin.H) {

	switch {

	case strings.Contains(err.Error(), "out of bounds"):
		return cfg.CODE_BADREQUEST, gin.H{"status": cfg.STATUS_BADREQUEST, "error": err.Error()}

	case strings.Contains(err.Error(), "not found"):
		return cfg.CODE_EMPTY, gin.H{"status": cfg.STATUS_EMPTY, "error": err.Error()}

	default:
		return cfg.CODE_BADREQUEST, gin.H{"status": cfg.STATUS_BADREQUEST, "error": err.Error()}
	}
}
`
