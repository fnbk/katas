package main

type ProductProvider struct {
	IDs []string
}

func (s *ProductProvider) ProductIDs() []string {
	return s.IDs
}
