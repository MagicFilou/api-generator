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
	handlerSnips "api-builder/templates/handlers/snippets"
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

		//Create models
		err = os.Mkdir(path+"/models", os.ModePerm)
		if err != nil {
			log.Println(err)
		}

		//Create handlers
		err = os.Mkdir(path+"/handlers", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	resourceDirs, err := ioutil.ReadDir(cfg.GetConfig().Repo.Models)
	if err != nil {
		panic(err)
	}

	var (
		allResources  resources.AllResources
		globalConfigs templates.GlobalConfigs
	)

	allResources.Resources = make(map[string]resources.Resource)

	//go through all models dir
	for _, dir := range resourceDirs {

		bufferResource, err := resources.ResourceFromFiles(cfg.GetConfig().Repo.Models+"/"+dir.Name(), dir.Name())
		if err != nil {
			panic(err)
		}

		allResources.Resources[dir.Name()] = bufferResource

		allResources.Destination = cfg.GetConfig().Repo.Service
	}

	//Ignored for now
	// allResources.Relations = allResources.GatherRelations()
	// allResources.GatherPriorities()

	for _, r := range allResources.Resources {

		//State disabled
		// if r.Deleted {
		// 	// fmt.Println("SKIP")
		// 	continue
		// }

		// GLOBAL ROUTES
		globalConfigs.GinRoutesConfig.Imports = append(
			globalConfigs.GinRoutesConfig.Imports,
			"\""+cfg.GetConfig().Repo.Service+"/routes/"+r.Name.Singular+"\"")

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

		modelConfig.Templates.Functions[0].Content[0].Custom = []templates.Template{}

		//Custom disabled
		// if r.PluralUnderscored == "service_configurations" {

		// 	// fmt.Println("WHATWHAT", r.PluralUnderscored)

		// 	modelConfig.Templates.Functions[0].Content[0].Custom = []templates.Template{modelSnips.ModelsCustomNestedServiceConfigurations}
		// }

		//Related disabled
		// relatedResources, ok := allResources.Relations[r.Resource.PluralUnderscored]
		// if ok {
		// 	for _, r := range utils.DeduplicateNames(relatedResources) {
		// 		modelConfig.Preload = resources.ToPreloads(relatedResources)
		// 		modelConfig.Imports = append(modelConfig.Imports, "\""+cfg.GetConfig().Gitea.ServiceRepo+"/models/"+strings.ToLower(strings.ReplaceAll(r.SingularUnderscored, "_", ""))+"\"")
		// 	}
		// }

		modelConfig.Templates.Struct = templates.Template(r.ToGoStruct())

		fmt.Println(r.Storage.Config)

		dir := strings.Join([]string{cfg.GetConfig().Repo.Service, cfg.MODELS_DIR, r.Singular}, "/")

		modelConfig.Write(dir, dir+"/"+r.Singular+".go")
		// END MODELS

		// DOCUMENTATION
		docTemplate := templates.DocConfig{
			Name: r.Name,
			//Relation disabled
			// Typescript:         r.ToTypescriptInterface(relatedResources),
			// TypescriptDTO:      r.ToTypescriptInterfaceDto(relatedResources),
			// TypescriptFullFile: r.ToTypescriptFullFile(relatedResources),
			Typescript:    r.ToTypescriptInterface(),
			TypescriptDTO: r.ToTypescriptInterfaceDto(),
			DocTemplates: templates.DocTemplates{
				Paths:     templates.Template(docSnips.Root),
				PathsByID: templates.Template(docSnips.ByID),
			},
		}

		//Relation disabled
		// if len(relatedResources) != 0 {
		// 	docTemplate.HasRelations = true
		// }

		// fmt.Println("RELATED RESOURCES")
		// fmt.Println(Pretty(relatedResources))

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
			Name:      r.Name,
			Templates: defaultconfigs.HandlerConfigDefaultTemplates,
		}

		handlerConfig.Templates.Functions[0].Content[0].Custom = []templates.Template{}

		if r.PluralUnderscored == "service_configurations" {
			// fmt.Println("WHATWHAT", r.PluralUnderscored)

			// handlerConfig.Templates.Functions[0].Content = append(handlerConfig.Templates.Functions[0].Content, handlerSnips.HandlersCustomNestedServiceConfigurations)
			handlerConfig.Templates.Functions[0].Content[0].Custom = []templates.Template{handlerSnips.HandlersCustomNestedServiceConfigurations}
		}

		dir = strings.Join([]string{cfg.GetConfig().Repo.Service, cfg.HANDLERS_DIR, r.Singular}, "/")
		handlerConfig.Write(dir, dir+"/"+r.Singular+".go")
		// END HANDLERS

		// ROUTES
		routeConfig := templates.FileConfig{
			Name:      r.Name,
			Templates: defaultconfigs.RouteConfigDefaultTemplates,
		}

		routeConfig.Templates.Functions[0].Content[0].Custom = []templates.Template{}

		if r.PluralUnderscored == "service_configurations" {

			// fmt.Println("WHATWHAT", r.PluralUnderscored)
			// content := routeConfig.Templates.Functions[0].Content
			// content = append(content, routeSnips.RoutesCustomNestedServiceConfigurations)
			// routeConfig.Templates.Functions[0].Content = append(routeConfig.Templates.Functions[0].Content, routeSnips.RoutesCustomNestedServiceConfigurations)
			routeConfig.Templates.Functions[0].Content[0].Custom = []templates.Template{routeSnips.RoutesCustomNestedServiceConfigurations}
			// fmt.Println(len(content))
			// routeConfig.Templates.Functions[0].Content = append(routeConfig.Templates.Functions[0].Content, routeSnips.RoutesCustomNestedServiceConfigurations)
			// fmt.Println(Pretty(routeConfig.Templates.Functions[0]))
		}

		dir = strings.Join([]string{cfg.GetConfig().Repo.Service, cfg.ROUTES_DIR, r.Singular}, "/")
		routeConfig.Write(dir, dir+"/"+r.Singular+".go")
		// END ROUTES

		// WRITE BACK STATE
		//state disabled
		//r.Resource.WriteResourceFile(cfg.GetConfig().Gitea.ModelsRepo + "/models/" + r.Resource.SingularUnderscored + "/" + r.Resource.SingularUnderscored + "_state.yaml")
		// END WRITE BACK STATE

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

	//Priorities disables
	// sortedKeys := make([]string, 0, len(allResources.Priorities))

	// for key := range allResources.Priorities {
	// 	// fmt.Println(key)
	// 	sortedKeys = append(sortedKeys, key)
	// }

	// sort.SliceStable(sortedKeys, func(i, j int) bool {
	// 	return allResources.Priorities[sortedKeys[i]] < allResources.Priorities[sortedKeys[j]]
	// })

	// fmt.Println(Pretty(sortedKeys))

	// for _, resourceName := range sortedKeys {

	// 	resourceWithState := allResources.Resources[resourceName]
	// 	// fmt.Println(Pretty(resourceWithState))

	// 	// fmt.Println(resourceName)
	// 	// fmt.Println("THIS", resourceWithState.New)
	// 	if resourceWithState.New {

	// 		err = resourceWithState.Resource.WriteSQLFiles(migrationDir, migrationVersion)
	// 		if err != nil {
	// 			panic(err)
	// 		}

	// 		migrationVersion++

	// 		resourceWithState.State = resourceWithState.Resource

	// 		err = resourceWithState.State.WriteResourceFile(cfg.GetConfig().Gitea.ModelsRepo + "/models/" + resourceWithState.Resource.SingularUnderscored + "/" + resourceWithState.Resource.SingularUnderscored + "_state.yaml")
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 	} else {

	// 		// CHECK FOR CHANGES

	// 		err := resourceWithState.HandleFieldChanges(migrationDir, migrationVersion)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 	}

	//}

	globalConfigs.GinRoutesConfig.Imports = append(
		globalConfigs.GinRoutesConfig.Imports,
		"cfg \""+allResources.Destination+"/configs\"")

	// fmt.Println(Pretty(globalConfigs.GinRoutesConfig))
	// fmt.Println(len(globalConfigs.GinRoutesConfig.Imports))
	// fmt.Println(len(globalConfigs.GinRoutesConfig.Templates.Functions))

	err = templates.WriteRoutesCollectFile(globalConfigs.GinRoutesConfig, routeSnips.RoutesFunctionCollectGroups, cfg.GetConfig().Repo.Service+"/routes/routes.go")
	if err != nil {
		panic(err)
	}

	sort.Strings(globalConfigs.OpenAPIRootConfig.Tags)

	err = templates.WriteRootDoc(cfg.GetConfig().Repo.Service+"/docs/wult-api.yaml", globalConfigs.OpenAPIRootConfig, docSnips.OpenAPISpec)
	if err != nil {
		panic(err)
	}

	// fmt.Println(Pretty(allResources.Relations))

	return 0
}

func Pretty(input interface{}) string {

	output, err := json.MarshalIndent(input, "", " ")
	if err != nil {
		fmt.Println(strings.ToUpper(err.Error()))
	}

	return string(output)
}
