package provider

import (
	"github.com/fnbk/pim/app/model"
)

type StructureProvider struct {
	Structures []model.Structure
}

func (s *StructureProvider) GetStructure(id string) *model.Structure {
	for _, p := range s.Structures {
		if p.ID == id {
			return &p
		}
	}
	return &model.Structure{}
}
