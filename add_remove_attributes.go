package main

func AddRemoveAttributes(p *Product, s *Structure) {
	var filteredAttributes []Attribute

	// filter attributes
	filteredAttributes = []Attribute{}
	for _, attr := range p.Attributes {
		if settingExists(s, &attr, TierOne) {
			filteredAttributes = append(filteredAttributes, attr)
		}
	}
	p.Attributes = filteredAttributes

	for _, b := range p.Bs {
		filteredAttributes = []Attribute{}
		for _, attr := range b.Attributes {
			if settingExists(s, &attr, TierOne) {
				filteredAttributes = append(filteredAttributes, attr)
			}
		}
		b.Attributes = filteredAttributes
		for _, c := range b.Cs {
			filteredAttributes = []Attribute{}
			for _, attr := range c.Attributes {
				if settingExists(s, &attr, TierOne) {
					filteredAttributes = append(filteredAttributes, attr)
				}
			}
			c.Attributes = filteredAttributes
		}
	}

	// add missing attributes as empty values
	for _, setting := range s.Settings {
		switch setting.Tier {
		case TierOne:
			found := false
			for _, attr := range p.Attributes {
				if attr.Name == setting.Name {
					found = true
				}
			}
			if !found {
				attr := Attribute{
					Name:  setting.Name,
					Value: "",
				}
				if p.Attributes == nil {
					p.Attributes = []Attribute{attr}
				} else {
					p.Attributes = append(p.Attributes, attr)
				}
			}
		case TierTwo:
			for i, _ := range p.Bs {
				found := false
				for _, attr := range p.Bs[i].Attributes {
					if attr.Name == setting.Name {
						found = true
					}
				}
				if !found {
					attr := Attribute{
						Name:  setting.Name,
						Value: "",
					}
					if p.Bs[i].Attributes == nil {
						p.Bs[i].Attributes = []Attribute{attr}
					} else {
						p.Bs[i].Attributes = append(p.Bs[i].Attributes, attr)
					}
				}
			}
		case TierThree:
			for i, _ := range p.Bs {
				for j, _ := range p.Bs[i].Cs {
					found := false
					for _, attr := range p.Bs[i].Cs[j].Attributes {
						if attr.Name == setting.Name {
							found = true
						}
					}
					if !found {
						attr := Attribute{
							Name:  setting.Name,
							Value: "",
						}
						if p.Bs[i].Cs[j].Attributes == nil {
							p.Bs[i].Cs[j].Attributes = []Attribute{attr}
						} else {
							p.Bs[i].Cs[j].Attributes = append(p.Bs[i].Cs[j].Attributes, attr)
						}
					}
				}
			}
		}
	}
}

func settingExists(s *Structure, attr *Attribute, tier Tier) bool {
	for _, setting := range s.Settings {
		if setting.Name == attr.Name && setting.Tier == tier {
			return true
		}
	}
	return false
}
