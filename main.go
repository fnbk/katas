package main

import (
	"bitbucket.scm.otto.de/scm/primary/pim/app"
	"bitbucket.scm.otto.de/scm/primary/pim/app/core"
	"bitbucket.scm.otto.de/scm/primary/pim/app/model"
	"bitbucket.scm.otto.de/scm/primary/pim/app/portal"
	"bitbucket.scm.otto.de/scm/primary/pim/app/provider"
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
