package defaultconfigs

import (
	"api-builder/templates"
	"api-builder/templates/files"
	hSnip "api-builder/templates/handlers/snippets"
	mSnip "api-builder/templates/models/snippets"
	rSnip "api-builder/templates/routes/snippets"
	sharedSnip "api-builder/templates/shared/snippets"
)

var RouteConfigDefaultTemplates templates.FileTemplates = templates.FileTemplates{

	Base:    files.Package,
	Package: sharedSnip.SharedPackage,
	Imports: rSnip.RoutesImports,
	Functions: []templates.Function{
		{
			Start: rSnip.RoutesFunctionGroup,
			Content: []templates.ContentTemplates{{
				Templates: []templates.Template{
					rSnip.RoutesDeclarationGroup,
					rSnip.RoutesGetAll,
					rSnip.RoutesCount,
					rSnip.RoutesNew,
					rSnip.RoutesPatch,
					rSnip.RoutesDelete,
				},
			},
			},
			End: rSnip.RoutesFunctionGroupEnd},
		{
			Start: rSnip.RoutesFunctionCheckError,
		},
	},
}

var HandlerConfigDefaultTemplates templates.FileTemplates = templates.FileTemplates{

	Base:    files.Package,
	Package: sharedSnip.SharedPackage,
	Imports: hSnip.HandlersImports,
	Functions: []templates.Function{
		{
			Content: []templates.ContentTemplates{{
				Templates: []templates.Template{
					hSnip.HandlersGetAll,
					hSnip.HandlersCount,
					hSnip.HandlersNew,
					hSnip.HandlersPatch,
					hSnip.HandlersDelete,
				},
			},
			},
		},
	},
}

var ModelConfigDefaultTemplates templates.FileTemplates = templates.FileTemplates{

	Base:    files.Package,
	Package: sharedSnip.SharedPackage,
	Imports: mSnip.ModelsImports,
	Functions: []templates.Function{
		{
			Content: []templates.ContentTemplates{{
				Templates: []templates.Template{
					mSnip.ModelsGetAll,
					mSnip.ModelsCount,
					mSnip.ModelsNew,
					mSnip.ModelsPatch,
					mSnip.ModelsDelete,
				},
			},
			},
		},
	},
}
