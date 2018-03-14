package provider

import (
	"github.com/fnbk/pim/app/model"
)

type ProductProvider struct {
	Products []model.Product
}

func (s *ProductProvider) ProductIDs() []string {
	IDs := []string{}
	for _, p := range s.Products {
		IDs = append(IDs, p.ID)
	}
	return IDs
}

func (s *ProductProvider) GetProduct(id string) *model.Product {
	for _, p := range s.Products {
		if p.ID == id {
			return &p
		}
	}
	return &model.Product{}
}
