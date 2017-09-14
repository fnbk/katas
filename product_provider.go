package main

type ProductProvider struct {
	Products []Product
}

func (s *ProductProvider) ProductIDs() []string {
	IDs := []string{}
	for _, p := range s.Products {
		IDs = append(IDs, p.ID)
	}
	return IDs
}

func (s *ProductProvider) GetProduct(id string) *Product {
	for _, p := range s.Products {
		if p.ID == id {
			return &p
		}
	}
	return &Product{}
}
