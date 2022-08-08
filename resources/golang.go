package resources

import (
	"strings"

	"api-builder/utils/translator"

	"github.com/iancoleman/strcase"
)

func (r Resource) ToGoStruct() string {

	b := strings.Builder{}

	if r.Description != "" {
		b.WriteString("// " + r.Description + "\n")
	}
	b.WriteString("type ")
	b.WriteString(r.Name.CamelUpper)
	b.WriteString(" struct {\n")

	if r.Storage.Config.DefaultFields {
		b.WriteString("models.SharedModel\n")
	}

	for _, field := range r.Storage.Fields {

		name := strcase.ToCamel(field.Name)

		b.WriteString(name + " ")
		b.WriteString(translator.ToGo(field.DataType) + " ")
		b.WriteString("`json:\"" + strcase.ToLowerCamel(field.Name) + "\"")

		for _, constraint := range field.Constraints {
			if constraint.Value == "not null" {
				b.WriteString(" binding:\"required\"")
			}
		}

		b.WriteString("`\n")

	}

	// if len(relatedResources) != 0 {

	// 	for _, r := range utils.DeduplicateNames(relatedResources) {
	// 		b.WriteString(strcase.ToCamel(r.PluralUnderscored) + " []" + r.Singular + "." + strcase.ToCamel(r.CamelUpper))
	// 		b.WriteString(" `json:\"" + strcase.ToLowerCamel(r.PluralUnderscored) + "\"`\n")
	// 	}
	// }

	b.WriteString("}")

	return b.String()
}
