package resources

import (
	"fmt"
	"io/ioutil"
	"strings"

	cfg "api-builder/configs"
	"api-builder/utils/name"

	"gopkg.in/yaml.v3"
)

type AllResources struct {
	Destination string
	Resources   map[string]Resource
}

func ResourceFromFiles(rootPath string, name string) (r Resource, err error) {

	// Get resource file
	r, err = ResourceFromFile(rootPath + "/" + name)
	if err != nil {

		return r, err
	}

	return r, nil
}

type Resource struct {
	name.Name   `yaml:"name,omitempty"`
	Description string `yaml:"description,omitempty"`
	Storage     `yaml:"storage,omitempty"`
}

func ResourceFromFile(path string) (r Resource, err error) {

	r, err = ReadResourceFile(path)
	if err != nil {
		return r, err
	}

	err = r.Name.Build()
	if err != nil {
		return r, err
	}

	if r.Storage.Config.Type == "" {
		r.Storage.Config.Type = "table"
	}

	if r.Storage.Config.DefaultFields == false {
		r.Storage.Config.DefaultFields = true
	}

	if r.Storage.Name == "" {
		r.Storage.Name = r.Name.PluralUnderscored
	}

	if r.Name.Short == "" {
		r.Name.Short = r.Name.PluralUnderscored[0:1]
	}

	for fieldIndex, field := range r.Fields {

		hasRelation := field.Relation.Resource != ""

		if field.Name == "display_id" {
			fmt.Printf("%+v\n", field)
		}
		if field.DataType == "" {
			if hasRelation {
				r.Fields[fieldIndex].DataType = "uid"
			} else {
				r.Fields[fieldIndex].DataType = "string"
			}
		}

		if hasRelation && field.Relation.Field == "" {
			r.Fields[fieldIndex].Relation.Field = "id"
		}

		if len(field.Constraints) != 0 {

			for constraintIndex, constraint := range field.Constraints {
				if constraint.Type == "" {
					r.Fields[fieldIndex].Constraints[constraintIndex].Type = "sql"
				}
			}
		}
	}

	return r, nil
}

func ReadResourceFile(path string) (r Resource, err error) {

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return r, err
	}

	err = yaml.Unmarshal(file, &r)
	if err != nil {
		return r, err
	}

	return r, nil
}

type Storage struct {
	Name   string        `yaml:"name,omitempty"`
	Config StorageConfig `yaml:"config,omitempty"`
	Fields []Field       `yaml:"fields,omitempty"`
}

type StorageConfig struct {
	Type          string `yaml:"type,omitempty" default:"table"`
	DefaultFields bool   `yaml:"default_fields,omitempty" default:"true"`
}

type Field struct {
	Name        string       `yaml:"name,omitempty"`
	DataType    string       `yaml:"data_type,omitempty"`
	Default     string       `yaml:"default,omitempty"`
	Example     string       `yaml:"example,omitempty"`
	Constraints []Constraint `yaml:"constraints"`
	Relation    Relation     `yaml:"relation,omitempty"`
}

func (f Field) hasRequired() bool {

	for _, constraint := range f.Constraints {
		if constraint.Value == "not null" {
			return true
		}
	}

	return false
}

type Constraint struct {
	Value string `yaml:"value"`
	Type  string `yaml:"type"`
}

type Relation struct {
	Resource   string `yaml:"resource,omitempty"`
	Field      string `yaml:"field,omitempty"`
	LocalField string `yaml:"local_field,omitempty"`
}

func (r Resource) hasRelations() bool {

	for _, field := range r.Storage.Fields {
		if field.Relation.Resource != "" {
			return true
		}
	}

	return false
}

func (r Resource) GetRelations() (relatedResources []name.Name, hasRelations bool) {

	if !r.hasRelations() {
		return relatedResources, false
	}

	for _, field := range r.Storage.Fields {

		if field.Relation.Resource != "" {

			relatedResources = append(relatedResources, name.Name{PluralUnderscored: field.Relation.Resource})
		}
	}

	return relatedResources, true
}

func (r Resource) RelationsAsImports() (relationImports []string) {

	relatedResources, ok := r.GetRelations()
	if ok {
		for _, relatedResource := range relatedResources {

			relationImports = append(relationImports, "\""+cfg.GetConfig().Repo.Service+"/models/"+strings.ToLower(strings.ReplaceAll(relatedResource.PluralUnderscored, "_", ""))+"\"")
		}
	}

	return relationImports
}

func (r Resource) WriteResourceFile(path string) (err error) {

	data, err := yaml.Marshal(&r)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, data, 0644)
}
