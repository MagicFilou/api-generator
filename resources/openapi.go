package resources

import (
	"strings"

	"api-builder/utils/translator"

	"github.com/iancoleman/strcase"
)

//Relations disabled
//func (r Resource) ToOpenAPIComponent(relatedResources []name.Name, withID bool) string {
func (r Resource) ToOpenAPIComponent(withID bool) string {

	b := strings.Builder{}

	b.WriteString("type: object\n")
	b.WriteString("tags:\n")
	b.WriteString("  - " + r.CamelLower + "\n")

	var hasRequired bool

	//only have required for withID (aka patch for now)
	if withID {

		b.WriteString("required:\n")
		b.WriteString("  - id\n")
	} else {

		for _, field := range r.Storage.Fields {
			for _, constraint := range field.Constraints {
				if constraint.Value == "not null" {
					if !hasRequired {
						b.WriteString("required:\n")
					}
					hasRequired = true
					b.WriteString("  - " + strcase.ToLowerCamel(field.Name) + "\n")
				}
			}
		}
	}

	b.WriteString("properties:\n")
	if withID {
		b.WriteString("  id:\n")
		b.WriteString("    $ref: " + translator.ToOpenAPI("uid") + "\n")
	}

	for _, field := range r.Storage.Fields {

		datatype := translator.ToOpenAPI(field.DataType)
		b.WriteString("  " + strcase.ToLowerCamel(field.Name) + ":\n")
		if strings.Contains(datatype, "components") {
			b.WriteString("    $ref: " + translator.ToOpenAPI(field.DataType) + "\n")
		} else {
			b.WriteString("    type: " + translator.ToOpenAPI(field.DataType) + "\n")
		}
		if field.Default != "" {
			b.WriteString("    example: \"" + field.Default + "\"\n")
		} else if field.Example != "" {
			b.WriteString("    example: \"" + field.Example + "\"\n")
		}
	}

	//Relations disabled
	// for _, relatedResource := range utils.DeduplicateNames(relatedResources) {
	// 	b.WriteString("  " + strcase.ToLowerCamel(relatedResource.PluralUnderscored) + ":\n")
	// 	b.WriteString("    type: array\n")
	// 	b.WriteString("    items:\n")
	// 	b.WriteString("      allOf:\n")
	// 	b.WriteString("        - $ref: \"#/components/schemas/SharedModel\"\n")
	// 	b.WriteString("        - $ref: \"../schemas/" + relatedResource.CamelLower + ".yaml\"\n")
	// }

	return b.String()
}

func (r Resource) ToOpenAPIPathsRef() string {

	b := strings.Builder{}

	b.WriteString("\n  /" + r.Plural + ":")
	b.WriteString("\n    $ref: \"./paths/" + r.CamelLower + ".yaml\"")

	b.WriteString("\n  /" + r.Plural + "/{id}:")
	b.WriteString("\n    $ref: \"./paths/" + r.CamelLower + "_id.yaml\"")

	return b.String()
}

func (r Resource) ToOpenAPITagDescription() string {

	b := strings.Builder{}

	b.WriteString("- name: " + r.CamelLower + "\n")
	b.WriteString("    description: " + r.Description + "\n")

	return b.String()
}
