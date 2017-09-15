package core

import (
	"bitbucket.scm.otto.de/scm/primary/pim/app/model"
	"bitbucket.scm.otto.de/scm/primary/pim/app/provider"
)

type Core struct {
	provider.ProductProvider
	provider.StructureProvider
}

func (k *Core) ShowProduct(id string) *model.Product {
	product := k.ProductProvider.GetProduct(id)
	structure := k.StructureProvider.GetStructure(id)
	AddRemoveAttributes(product, structure)
	return product
}

func (k *Core) ListProductIDs() []string {
	return k.ProductProvider.ProductIDs()
}
