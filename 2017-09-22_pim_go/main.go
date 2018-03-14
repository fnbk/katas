package main

import (
	"github.com/fnbk/pim/app"
	"github.com/fnbk/pim/app/core"
	"github.com/fnbk/pim/app/model"
	"github.com/fnbk/pim/app/portal"
	"github.com/fnbk/pim/app/provider"
)

func main() {
	productProvider := provider.ProductProvider{
		Products: []model.Product{
			{
				ID:   "1",
				Name: "Name1",
			},
			{
				ID:   "2",
				Name: "Name2",
			},
			{
				ID:   "3",
				Name: "Name3",
			},
		},
	}
	structureProvider := provider.StructureProvider{
		Structures: []model.Structure{
			{
				ID:   "123",
				Name: "Name123",
			},
		},
	}
	c := core.Core{ProductProvider: productProvider, StructureProvider: structureProvider}
	httpPortal := portal.HTTPPortal{Core: c}
	a := app.App{HTTPPortal: httpPortal}
	a.Run()
}
