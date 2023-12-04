package docs

import "github.com/mvrilo/go-redoc"

func Initialize() redoc.Redoc {
	doc := redoc.Redoc{
		Title:       "Documentation of ShopSystem",
		Description: "Documentation describes working procedures of AuthSystem like structs, handlers, etc.",
		SpecFile:    "./docs/api_shopV1.swagger.json",
		SpecPath:    "/docs/api_shopV1.swagger.json",
		DocsPath:    "/docs",
	}

	return doc
}
