package core

import (
	"github.com/fnbk/pim/app/model"
	"github.com/fnbk/pim/app/provider"
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
