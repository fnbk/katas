package main

type Product A

type A struct {
	ID          string
	Name        string
	StructureID string
	Attributes  []Attribute
	Bs          []B
}

type B struct {
	ID         string
	Name       string
	Attributes []Attribute
	Cs         []C
}

type C struct {
	ID         string
	Name       string
	Attributes []Attribute
}

type Attribute struct {
	ID    string
	Name  string
	Value string
	State string
}

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
