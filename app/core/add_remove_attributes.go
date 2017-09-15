package core

import (
	"bitbucket.scm.otto.de/scm/primary/pim/app/model"
)

func AddRemoveAttributes(p *model.Product, s *model.Structure) {
	var filteredAttributes []model.Attribute

	// filter attributes
	filteredAttributes = []model.Attribute{}
	for _, attr := range p.Attributes {
		if settingExists(s, &attr, model.TierOne) {
			filteredAttributes = append(filteredAttributes, attr)
		}
	}
	p.Attributes = filteredAttributes

	for _, b := range p.Bs {
		filteredAttributes = []model.Attribute{}
		for _, attr := range b.Attributes {
			if settingExists(s, &attr, model.TierOne) {
				filteredAttributes = append(filteredAttributes, attr)
			}
		}
		b.Attributes = filteredAttributes
		for _, c := range b.Cs {
			filteredAttributes = []model.Attribute{}
			for _, attr := range c.Attributes {
				if settingExists(s, &attr, model.TierOne) {
					filteredAttributes = append(filteredAttributes, attr)
				}
			}
			c.Attributes = filteredAttributes
		}
	}

	// add missing attributes as empty values
	for _, setting := range s.Settings {
		switch setting.Tier {
		case model.TierOne:
			found := false
			for _, attr := range p.Attributes {
				if attr.Name == setting.Name {
					found = true
				}
			}
			if !found {
				attr := model.Attribute{
					Name:  setting.Name,
					Value: "",
				}
				if p.Attributes == nil {
					p.Attributes = []model.Attribute{attr}
				} else {
					p.Attributes = append(p.Attributes, attr)
				}
			}
		case model.TierTwo:
			for i, _ := range p.Bs {
				found := false
				for _, attr := range p.Bs[i].Attributes {
					if attr.Name == setting.Name {
						found = true
					}
				}
				if !found {
					attr := model.Attribute{
						Name:  setting.Name,
						Value: "",
					}
					if p.Bs[i].Attributes == nil {
						p.Bs[i].Attributes = []model.Attribute{attr}
					} else {
						p.Bs[i].Attributes = append(p.Bs[i].Attributes, attr)
					}
				}
			}
		case model.TierThree:
			for i, _ := range p.Bs {
				for j, _ := range p.Bs[i].Cs {
					found := false
					for _, attr := range p.Bs[i].Cs[j].Attributes {
						if attr.Name == setting.Name {
							found = true
						}
					}
					if !found {
						attr := model.Attribute{
							Name:  setting.Name,
							Value: "",
						}
						if p.Bs[i].Cs[j].Attributes == nil {
							p.Bs[i].Cs[j].Attributes = []model.Attribute{attr}
						} else {
							p.Bs[i].Cs[j].Attributes = append(p.Bs[i].Cs[j].Attributes, attr)
						}
					}
				}
			}
		}
	}
}

func settingExists(s *model.Structure, attr *model.Attribute, tier model.Tier) bool {
	for _, setting := range s.Settings {
		if setting.Name == attr.Name && setting.Tier == tier {
			return true
		}
	}
	return false
}
