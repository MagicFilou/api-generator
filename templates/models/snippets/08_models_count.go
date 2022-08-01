package snippets

import (
	"api-builder/templates"
)

const ModelsCount templates.Template = `

func Count(distinct, group string) (ca []models.CountData, err error) {

	db, err := models.GetCon()
	if err != nil {
		return ca, err
	}

	var result *gorm.DB
	var count int64

	switch models.CountType(distinct, group) {
	case "group_distinct":

		result = db.Model(&{{ .Name.CamelUpper }}{}).Select("count(distinct(" + distinct + ")) as count, " + group + " as group_by").Group(group).Scan(&ca)
	case "group":

		result = db.Model(&{{ .Name.CamelUpper }}{}).Select("count(id) as count, " + group + " as group_by").Group(group).Scan(&ca)
	default:

		result = db.Find(&{{ .Name.CamelUpper }}{}).Count(&count)
	}

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return ca, fmt.Errorf("not found")
		}
		return ca, err
	}

	if len(ca) == 0 {
		ca = append(ca, models.CountData{
			Count:   int(count),
			GroupBy: "all",
		})
	}

	if result.RowsAffected == 0 {
		return ca, fmt.Errorf("not found")
	}

	return ca, nil
}
`
