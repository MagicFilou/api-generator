package snippets

import "api-builder/templates"

const ModelsCustomNestedServiceConfigurations templates.Template = `
func GetAllNested() (serviceconfigurations []ServiceConfiguration, err error) {

	db, err := models.GetCon()
	if err != nil {
		return serviceconfigurations, err
	}

	var result *gorm.DB

	result = db.Preload("Integrations.Datasets.Fields.Identifiers").Preload(clause.Associations).Find(&serviceconfigurations)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return serviceconfigurations, fmt.Errorf("not found")
		}
		return serviceconfigurations, err
	}

	if result.RowsAffected == 0 {
		return serviceconfigurations, fmt.Errorf("not found")
	}

	return serviceconfigurations, nil
}
`
