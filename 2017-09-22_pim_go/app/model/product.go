package model

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
