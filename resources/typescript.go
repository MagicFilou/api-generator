package resources

import (
	"io/ioutil"
	"strings"

	"api-builder/utils/translator"

	"github.com/iancoleman/strcase"
)

func (r Resource) WriteTypescriptFiles(path string) error {

	err := ioutil.WriteFile(path+"/"+r.Name.CamelUpper+".ts", []byte(r.ToTypescriptInterface()), 0644)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path+"/"+r.Name.CamelUpper+"_dto.ts", []byte(r.ToTypescriptInterfaceDto()), 0644)
}

//Relation disabled -> apply to all typescript types
//func (r Resource) ToTypescriptInterface(relatedResources []name.Name) string {
func (r Resource) ToTypescriptInterface() string {

	b := strings.Builder{}

	b.WriteString("export interface " + r.Name.CamelUpper + " {\n")

	if r.Storage.Config.DefaultFields {

		for _, field := range DefaultFields {
			b.WriteString("          " + strcase.ToLowerCamel(field.Name) + ": ")
			b.WriteString(translator.ToTypescript(field.DataType))
			if !field.hasRequired() {
				b.WriteString(" | null")
			}
			b.WriteString("\n")
		}
	}

	for _, field := range r.Storage.Fields {
		b.WriteString("          " + strcase.ToLowerCamel(field.Name) + ": ")
		b.WriteString(translator.ToTypescript(field.DataType))
		if !field.hasRequired() {
			b.WriteString(" | null")
		}
		b.WriteString("\n")
	}

	//Relations disabled
	// if len(relatedResources) != 0 {

	// 	for _, r := range utils.DeduplicateNames(relatedResources) {

	// 		b.WriteString("          " + strcase.ToLowerCamel(r.PluralUnderscored) + "?: " + strcase.ToCamel(r.SingularUnderscored) + "[]\n")
	// 	}
	// }

	b.WriteString("        }\n")

	return b.String()
}

func (r Resource) ToTypescriptInterfaceDto() string {

	b := strings.Builder{}

	b.WriteString("export interface " + r.Name.CamelUpper + "Dto {\n")

	for _, field := range r.Storage.Fields {
		b.WriteString("          " + strcase.ToLowerCamel(field.Name) + ": ")
		b.WriteString(translator.ToTypescript(field.DataType))
		if !field.hasRequired() {
			b.WriteString(" | null")
		}
		b.WriteString("\n")
	}

	b.WriteString("        }\n")

	return b.String()
}

func (r Resource) ToTypescriptFullFile() string {

	b := strings.Builder{}

	// for index, r := range utils.DeduplicateNames(relatedResources) {
	// 	if index != 0 {
	// 		b.WriteString("        ")
	// 	}
	// 	b.WriteString("import { " + strcase.ToCamel(r.SingularUnderscored) + " } from './" + strcase.ToCamel(r.SingularUnderscored) + "'\n")
	// }

	b.WriteString("\n")

	// b.WriteString("          " + r.ToTypescriptInterface(relatedResources))
	b.WriteString("        " + r.ToTypescriptInterface())

	b.WriteString("\n")

	// b.WriteString("          " + r.ToTypescriptInterfaceDto(relatedResources))
	b.WriteString("        " + r.ToTypescriptInterfaceDto())

	return b.String()
}
