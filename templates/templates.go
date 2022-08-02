package templates

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"os"
	"text/template"

	cfg "api-builder/configs"
	"api-builder/utils/name"
)

type Template string

type GlobalConfigs struct {
	GinRoutesConfig   FileConfig
	OpenAPIRootConfig RootConfig
}

type FileConfig struct {
	ParentFolder string
	Name         name.Name
	Preload      string
	Imports      []string
	Templates    FileTemplates
}

type DocConfig struct {
	Name               name.Name
	Typescript         string
	TypescriptDTO      string
	TypescriptFullFile string
	//HasRelations       bool
	DocTemplates
}

type DocTemplates struct {
	Component Template
	Paths     Template
	PathsByID Template
}

type FileTemplates struct {
	Base      Template
	Package   Template
	Imports   Template
	Struct    Template
	Functions []Function
}

type Function struct {
	Start   Template
	Content []ContentTemplates
	End     Template
}

type ContentTemplates struct {
	Templates []Template
	Custom    []Template
}

func (t Template) BuildFromFileConfig(f FileConfig) Template {

	ut, err := template.New("template").Parse(string(t))
	if err != nil {
		panic(err)
	}

	var b bytes.Buffer

	err = ut.Execute(&b, f)
	if err != nil {
		panic(err)
	}

	return Template(b.String())
}

func (t Template) BuildFromDocConfig(d DocConfig) Template {

	ut, err := template.New("template").Parse(string(t))
	if err != nil {
		panic(err)
	}

	var b bytes.Buffer

	err = ut.Execute(&b, d)
	if err != nil {
		panic(err)
	}

	return Template(b.String())
}

func (f FileConfig) Build() string {

	template := f.Templates.Base.BuildFromFileConfig(f)

	formatted, err := format.Source([]byte(template.BuildFromFileConfig(f)))
	if err != nil {
		panic(err)
	}

	return string(formatted)
}

func (d DocConfig) BuildComponent() string {

	template := d.DocTemplates.Component.BuildFromDocConfig(d)

	return string(template.BuildFromDocConfig(d))
}

func (d DocConfig) BuildPaths() string {

	template := d.DocTemplates.Paths.BuildFromDocConfig(d)

	return string(template.BuildFromDocConfig(d))
}

func (d DocConfig) BuildPathsForByID() string {

	template := d.DocTemplates.PathsByID.BuildFromDocConfig(d)

	return string(template.BuildFromDocConfig(d))
}

func (f FileConfig) Write(dir, path string) error {

	os.Mkdir(dir, 0755)

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(path, []byte(f.Build()), 0644); err != nil {
		return err
	}

	return file.Close()
}

func (d DocConfig) Write(path string, resourceName name.Name, withID bool) (err error) {

	if withID {

		err = d.WriteComponent(path + "/schemas/" + resourceName.CamelLower + "_with_id" + cfg.CONFIG_EXTENSION)
		if err != nil {
			return err
		}
	} else {
		err = d.WriteComponent(path + "/schemas/" + resourceName.CamelLower + cfg.CONFIG_EXTENSION)
		if err != nil {
			return err
		}
	}

	err = d.WritePaths(path + "/paths/" + resourceName.CamelLower + cfg.CONFIG_EXTENSION)
	if err != nil {
		return err
	}

	err = d.WritePathsForByID(path + "/paths/" + resourceName.CamelLower + "_id" + cfg.CONFIG_EXTENSION)
	if err != nil {
		return err
	}

	return nil
}

func (d DocConfig) WriteComponent(path string) error {

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(path, []byte(d.BuildComponent()), 0644); err != nil {
		return err
	}

	return file.Close()
}

func (d DocConfig) WritePaths(path string) error {

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(path, []byte(d.BuildPaths()), 0644); err != nil {
		return err
	}

	return file.Close()
}

func (d DocConfig) WritePathsForByID(path string) error {

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(path, []byte(d.BuildPathsForByID()), 0644); err != nil {
		return err
	}

	return file.Close()
}

type RootConfig struct {
	Paths           []string
	Tags            []string
	TagDescriptions []string
}

func WriteRootDoc(path string, config RootConfig, rootSpec Template) (err error) {

	ut, err := template.New("template").Parse(string(rootSpec))

	var b bytes.Buffer

	err = ut.Execute(&b, config)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	if err = ioutil.WriteFile(path, []byte(b.String()), 0644); err != nil {
		panic(err)
	}
	return nil
}

func WriteRoutesCollectFile(config FileConfig, rootSpec Template, path string) (err error) {

	ut, err := template.New("template").Parse(string(rootSpec))

	var b bytes.Buffer

	err = ut.Execute(&b, config)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	formatted, err := format.Source([]byte(b.String()))
	if err != nil {
		panic(err)
	}

	if err = ioutil.WriteFile(path, formatted, 0644); err != nil {
		panic(err)
	}

	return nil
}
