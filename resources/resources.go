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
	//Relations   map[string][]name.Name
	//Priorities  map[string]int
	Resources map[string]Resource
}

//Relations diabled for now 010822
// func (allR AllResources) GatherRelations() (relations map[string][]name.Name) {

// 	relations = make(map[string][]name.Name)

// 	for _, r := range allR.Resources {

// 		relatedResources, ok := r.Resource.GetRelations()
// 		if !ok {
// 			continue
// 		}

// 		for _, relatedName := range relatedResources {
// 			relations[relatedName.PluralUnderscored] = append(relations[relatedName.PluralUnderscored], r.Name)
// 		}
// 	}

// 	return relations
// }

//Not sure What priorities are for but let's ignore it for now
// // NOTE: this is not the way to do it. I'm sorry. Frankly, im embarrassed.
// func (allR *AllResources) GatherPriorities() {

// 	allR.Priorities = make(map[string]int)

// 	for _, r := range allR.Resources {

// 		allR.Priorities[r.Resource.SingularUnderscored] = 1

// 		relatedResources, ok := r.Resource.GetRelations()
// 		if ok {
// 			allR.Priorities[r.Resource.SingularUnderscored]++
// 		} else {
// 			continue
// 		}

// 		for _, relatedResource := range relatedResources {
// 			value, ok := allR.Priorities[relatedResource.SingularUnderscored]
// 			if ok {
// 				allR.Priorities[r.Resource.SingularUnderscored] = value + 1
// 				continue
// 			} else {
// 				allR.Priorities[relatedResource.SingularUnderscored] = allR.Priorities[r.SingularUnderscored] + 1
// 			}
// 		}
// 	}
// }

// type ResourceWithState struct {
// 	Resource
// 	New     bool
// 	Deleted bool
// 	State   Resource
// }

// func ResourceWithStateFromFiles(rootPath string, name string) (r ResourceWithState, err error) {

// 	// Get resource file
// 	r.Resource, err = ResourceFromFile(rootPath + "/" + name + ".yaml")
// 	if err != nil {
// 		// If this file does not exist, the resource has been marked for deletion
// 		if strings.Contains(err.Error(), "no such file or directory") {
// 			r.Deleted = true
// 		} else {
// 			return r, err
// 		}
// 	}

// 	// Get state file
// 	r.State, err = ResourceFromFile(rootPath + "/" + name + "_state.yaml")
// 	if err != nil {
// 		// If this file does not exist, the resource is new
// 		if strings.Contains(err.Error(), "no such file or directory") {
// 			r.New = true
// 		} else {
// 			return r, err
// 		}
// 	}

// 	return r, nil
// }

func ResourceFromFiles(rootPath string, name string) (r Resource, err error) {

	// Get resource file
	r, err = ResourceFromFile(rootPath + "/" + name + ".yaml")
	if err != nil {
		//To do with states
		// If this file does not exist, the resource has been marked for deletion
		// if strings.Contains(err.Error(), "no such file or directory") {
		// 	r.Deleted = true
		// } else {
		return r, err
		//}
	}

	return r, nil
}

// func (r ResourceWithState) calculateChanges() (newFields map[string]Field, removedFields map[string]Field, changed bool) {

// 	newFields = make(map[string]Field)
// 	removedFields = make(map[string]Field)

// 	if !reflect.DeepEqual(r.Resource.Fields, r.State.Fields) {

// 		resourceFields := make(map[string]Field)
// 		stateFields := make(map[string]Field)

// 		for _, resourceField := range r.Resource.Fields {

// 			resourceFields[resourceField.Name] = resourceField
// 		}

// 		for _, stateField := range r.State.Fields {

// 			stateFields[stateField.Name] = stateField
// 		}

// 		for key, value := range stateFields {
// 			_, ok := resourceFields[key]
// 			if !ok {
// 				changed = true
// 				removedFields[key] = value
// 			}
// 		}

// 		for key, value := range resourceFields {
// 			_, ok := stateFields[key]
// 			if !ok {
// 				changed = true
// 				newFields[key] = value
// 			}
// 		}
// 	}

// 	if changed {

// 		for key, field := range newFields {
// 			fmt.Println("FIELD ADDED: ", key, field)
// 		}
// 		for key, field := range removedFields {
// 			fmt.Println("FIELD REMOVED: ", key, field)
// 		}
// 	}

// 	return newFields, removedFields, changed
// }

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
