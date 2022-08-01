package files

import (
	"api-builder/templates"
)

const SharedPackage templates.Template = `
package {{ .Name.Singular }}
  `
