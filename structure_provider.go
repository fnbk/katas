package main

type Structure struct {
	ID       string
	Name     string
	Settings []Setting
}

type Setting struct {
	ID       string
	Name     string
	Tier     Tier
	Datatype string
}

type Tier int

const (
	TierOne Tier = 1 + iota
	TierTwo
	TierThree
)

type StructureProvider struct {
	Structures []Structure
}

func (s *StructureProvider) GetStructure(id string) *Structure {
	for _, p := range s.Structures {
		if p.ID == id {
			return &p
		}
	}
	return &Structure{}
}
