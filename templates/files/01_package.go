package files

import (
	"api-builder/templates"
)

const Package templates.Template = `
  {{ .Templates.Package }} 
  {{ .Templates.Imports }}
  {{ .Templates.Struct }}
  {{ range .Templates.Functions }}
    {{ .Start }}
    {{ range .Content }}
      {{ range .Templates }}
        {{ . }}
      {{ end }}
      {{ range .Custom }}
        {{ . }}
      {{ end }}
    {{ end }}
    {{ .End }}
  {{ end }}
`
