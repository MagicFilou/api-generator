package snippets

import (
	"api-builder/templates"
)

const HandlersCount templates.Template = `
func Count(distinct, group string, conKey int) ([]models.CountData, error) {

	return {{ .Name.Singular }}.Count(distinct, group, conKey)
}
`

//postgresql://tnbxhwrsdtjtwzk:ifpangghhhenhhf@wult-dev.crmfjao48mj7.eu-central-1.rds.amazonaws.com:5432/com_200_llcdpfggelqmmdx
//wult-dev.crmfjao48mj7.eu-central-1.rds.amazonaws.com	5432	tnbxhwrsdtjtwzk	ifpangghhhenhhf	com_200_llcdpfggelqmmdx	postgresql
