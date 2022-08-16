package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"

	cfg "api-builder/configs"
	"api-builder/defaultconfigs"
	"api-builder/resources"
	"api-builder/templates"
	docSnips "api-builder/templates/docs"
	routeSnips "api-builder/templates/routes/snippets"
	"api-builder/utils"
)

func Run() int {

	fmt.Println(Pretty(cfg.GetConfig()))

	//Git disabled
	// err := git.CloneOrPull(cfg.GetConfig().Gitea.ServiceRepo, cfg.GetConfig().Gitea.ServiceRepo, "wult")
	// if err != nil {
	// 	panic(err)
	// }

	// err = git.CloneOrPull(cfg.GetConfig().Gitea.ModelsRepo, cfg.GetConfig().Gitea.ModelsRepo, "wult")
	// if err != nil {
	// 	panic(err)
	// }

	path := cfg.GetConfig().Repo.Service
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
		//create for docs
		err = os.Mkdir(path+"/docs", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
		err = os.Mkdir(path+"/docs/paths", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
		err = os.Mkdir(path+"/docs/schemas", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
		//create for migrations
		err = os.Mkdir(path+"/migrations", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
		err = os.Mkdir(path+"/migrations/migration_files", os.ModePerm)
		if err != nil {
			log.Println(err)
		}

		//Create routes
		err = os.Mkdir(path+"/routes", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
		err = os.Mkdir(path+"/routes/"+cfg.GetConfig().Repo.SubModels, os.ModePerm)
		if err != nil {
			log.Println(err)
		}

		//Create models
		err = os.Mkdir(path+"/models", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
		err = os.Mkdir(path+"/models/"+cfg.GetConfig().Repo.SubModels, os.ModePerm)
		if err != nil {
			log.Println(err)
		}

		//Create handlers
		err = os.Mkdir(path+"/handlers", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
		err = os.Mkdir(path+"/handlers/"+cfg.GetConfig().Repo.SubModels, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	resourceDirs, err := ioutil.ReadDir(cfg.GetConfig().Repo.Models + "/" + cfg.GetConfig().Repo.SubModels)
	if err != nil {
		panic(err)
	}

	var (
		allResources  resources.AllResources
		globalConfigs templates.GlobalConfigs
	)

	allResources.Resources = make(map[string]resources.Resource)
	globalConfigs.GinRoutesConfig.ParentFolder = cfg.GetConfig().Repo.SubModels

	//go through all models dir
	for _, dir := range resourceDirs {

		bufferResource, err := resources.ResourceFromFiles(cfg.GetConfig().Repo.Models+"/"+cfg.GetConfig().Repo.SubModels, dir.Name())
		if err != nil {
			panic(err)
		}

		allResources.Resources[dir.Name()] = bufferResource

		allResources.Destination = cfg.GetConfig().Repo.Service
	}

	for _, r := range allResources.Resources {

		// GLOBAL ROUTES
		globalConfigs.GinRoutesConfig.Imports = append(
			globalConfigs.GinRoutesConfig.Imports,
			"\"generator-gw/routes/"+cfg.GetConfig().Repo.SubModels+"/"+r.Name.Singular+"\"")

		globalConfigs.GinRoutesConfig.Templates.Functions = append(
			globalConfigs.GinRoutesConfig.Templates.Functions,
			templates.Function{
				Start: templates.Template(r.Singular + "." + r.CamelUpper + "Group(r)\n"),
			})
		// END GLOBAL ROUTES

		// MODELS
		modelConfig := templates.FileConfig{
			Name:      r.Name,
			Templates: defaultconfigs.ModelConfigDefaultTemplates,
		}

		modelConfig.Templates.Struct = templates.Template(r.ToGoStruct())

		fmt.Println(r.Storage.Config)

		dir := strings.Join([]string{cfg.GetConfig().Repo.Service, cfg.MODELS_DIR, cfg.GetConfig().Repo.SubModels, r.Singular}, "/")

		modelConfig.Write(dir, dir+"/"+r.Singular+".go")
		// END MODELS

		// DOCUMENTATION
		docTemplate := templates.DocConfig{
			Name:               r.Name,
			Typescript:         r.ToTypescriptInterface(),
			TypescriptDTO:      r.ToTypescriptInterfaceDto(),
			TypescriptFullFile: r.ToTypescriptFullFile(),
			DocTemplates: templates.DocTemplates{
				Paths:     templates.Template(docSnips.Root),
				PathsByID: templates.Template(docSnips.ByID),
			},
		}

		docTemplate.Component = templates.Template(r.ToOpenAPIComponent(false))
		err = docTemplate.Write(cfg.GetConfig().Repo.Service+"/docs", r.Name, false)
		if err != nil {
			panic(err)
		}

		docTemplate.Component = templates.Template(r.ToOpenAPIComponent(true))
		err = docTemplate.Write(cfg.GetConfig().Repo.Service+"/docs", r.Name, true)
		if err != nil {
			panic(err)
		}

		globalConfigs.OpenAPIRootConfig.Paths = append(globalConfigs.OpenAPIRootConfig.Paths, r.ToOpenAPIPathsRef())
		globalConfigs.OpenAPIRootConfig.Tags = append(globalConfigs.OpenAPIRootConfig.Tags, "- "+r.CamelLower)
		globalConfigs.OpenAPIRootConfig.TagDescriptions = append(globalConfigs.OpenAPIRootConfig.TagDescriptions, r.ToOpenAPITagDescription())

		fmt.Println(r.ToOpenAPITagDescription())
		// END DOCUMENTATION

		// HANDLERS
		handlerConfig := templates.FileConfig{
			ParentFolder: cfg.GetConfig().Repo.SubModels,
			Name:         r.Name,
			Templates:    defaultconfigs.HandlerConfigDefaultTemplates,
		}

		dir = strings.Join([]string{cfg.GetConfig().Repo.Service, cfg.HANDLERS_DIR, cfg.GetConfig().Repo.SubModels, r.Singular}, "/")
		handlerConfig.Write(dir, dir+"/"+r.Singular+".go")
		// END HANDLERS

		// ROUTES
		routeConfig := templates.FileConfig{
			SecretType:   cfg.GetConfig().SecretType,
			ParentFolder: cfg.GetConfig().Repo.SubModels,
			Name:         r.Name,
			Templates:    defaultconfigs.RouteConfigDefaultTemplates,
		}

		routeConfig.Templates.Functions[0].Content[0].Custom = []templates.Template{}

		dir = strings.Join([]string{cfg.GetConfig().Repo.Service, cfg.ROUTES_DIR, cfg.GetConfig().Repo.SubModels, r.Singular}, "/")
		routeConfig.Write(dir, dir+"/"+r.Singular+".go")
		// END ROUTES

		allResources.Resources[r.SingularUnderscored] = r
	}

	migrationDir := cfg.GetConfig().Repo.Migrations + "/migrations/migration_files"

	_, err = utils.FileVersion(migrationDir)
	if err != nil {
		panic(err)
	}

	for _, res := range allResources.Resources {
		err = res.WriteSQLFiles(migrationDir, 1)
		if err != nil {
			panic(err)
		}
	}

	globalConfigs.GinRoutesConfig.Imports = append(
		globalConfigs.GinRoutesConfig.Imports,
		"cfg \"generator-gw/configs\"")

	err = templates.WriteRoutesCollectFile(globalConfigs.GinRoutesConfig, routeSnips.RoutesFunctionCollectGroups, cfg.GetConfig().Repo.Service+"/routes/routes.go")
	if err != nil {
		panic(err)
	}

	sort.Strings(globalConfigs.OpenAPIRootConfig.Tags)

	err = templates.WriteRootDoc(cfg.GetConfig().Repo.Service+"/docs/"+cfg.GetConfig().Repo.SubModels+".yaml", globalConfigs.OpenAPIRootConfig, docSnips.OpenAPISpec)
	if err != nil {
		panic(err)
	}

	return 0
}

func Pretty(input interface{}) string {

	output, err := json.MarshalIndent(input, "", " ")
	if err != nil {
		fmt.Println(strings.ToUpper(err.Error()))
	}

	return string(output)
}
