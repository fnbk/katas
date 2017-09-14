package main

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
