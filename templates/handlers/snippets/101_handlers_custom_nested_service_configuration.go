package snippets

import "api-builder/templates"

const HandlersCustomNestedServiceConfigurations templates.Template = `
func GetAllNested() ([]serviceconfiguration.ServiceConfiguration, error) {

	return serviceconfiguration.GetAllNested()
}
`
