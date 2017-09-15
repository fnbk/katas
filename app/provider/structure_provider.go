package provider

import "bitbucket.scm.otto.de/scm/primary/pim/app/model"

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
