package core

import (
	"reflect"
	"testing"

	"github.com/fnbk/pim/app/model"
)

func TestAddRemoveAttributes(t *testing.T) {
	testcases := []struct {
		ProductIn   model.Product
		StructureIn model.Structure
		ProductOut  model.Product
	}{
		// P_In:{} S:{B1} P_Out:{B1}
		{
			ProductIn: model.Product{
				Bs: []model.B{model.B{}},
			},
			StructureIn: model.Structure{
				Settings: []model.Setting{
					{Name: "B1", Tier: model.TierTwo},
				},
			},
			ProductOut: model.Product{
				Attributes: []model.Attribute{},
				Bs: []model.B{
					{
						Attributes: []model.Attribute{
							{
								Name:  "B1",
								Value: "",
							},
						},
					},
				},
			},
		},

		// P_In:{A1X} S:{A1} P_Out:{A1X}
		{
			ProductIn: model.Product{
				Attributes: []model.Attribute{
					{
						Name:  "A1",
						Value: "X",
					},
				},
			},
			StructureIn: model.Structure{
				Settings: []model.Setting{
					{Name: "A1", Tier: model.TierOne},
				},
			},
			ProductOut: model.Product{
				Attributes: []model.Attribute{
					{
						Name:  "A1",
						Value: "X",
					},
				},
			},
		},

		// P_In:{A1X} S:{} P_Out:{}
		{
			ProductIn: model.Product{
				Attributes: []model.Attribute{
					{
						Name:  "A1",
						Value: "X",
					},
				},
			},
			StructureIn: model.Structure{},
			ProductOut:  model.Product{Attributes: []model.Attribute{}},
		},

		// P_In:{A1X,B1X} S:{B1} P_Out:{B1X}
		{
			ProductIn: model.Product{
				Attributes: []model.Attribute{
					{
						Name:  "A1",
						Value: "X",
					},
				},
				Bs: []model.B{
					{
						Attributes: []model.Attribute{
							{
								Name:  "B1",
								Value: "X",
							},
						},
					},
				},
			},
			StructureIn: model.Structure{
				Settings: []model.Setting{
					{Name: "B1", Tier: model.TierTwo},
				},
			},
			ProductOut: model.Product{
				Attributes: []model.Attribute{},
				Bs: []model.B{
					{
						Attributes: []model.Attribute{
							{
								Name:  "B1",
								Value: "X",
							},
						},
					},
				},
			},
		},

		// P_In:{A1X} S:{A1,A2} P_Out:{A1X,A2E}
		{
			ProductIn: model.Product{
				Attributes: []model.Attribute{
					{
						Name:  "A1",
						Value: "X",
					},
				},
			},
			StructureIn: model.Structure{
				Settings: []model.Setting{
					{Name: "A1", Tier: model.TierOne},
					{Name: "A2", Tier: model.TierOne},
				},
			},
			ProductOut: model.Product{
				Attributes: []model.Attribute{
					{
						Name:  "A1",
						Value: "X",
					},
					{
						Name:  "A2",
						Value: "",
					},
				},
			},
		},

		// P_In:{A1X,B1X,C1X} S:{A1,A2,B1,B2,C1,C2} P_Out:{A1X,A2E,B1X,B2E,C1X,C2E}
		{
			ProductIn: model.Product{
				Attributes: []model.Attribute{
					{
						Name:  "A1",
						Value: "X",
					},
				},
				Bs: []model.B{
					{
						Attributes: []model.Attribute{
							{
								Name:  "B1",
								Value: "X",
							},
						},
						Cs: []model.C{
							{
								Attributes: []model.Attribute{
									{
										Name:  "C1",
										Value: "X",
									},
								},
							},
						},
					},
				},
			},
			StructureIn: model.Structure{
				Settings: []model.Setting{
					{Name: "A1", Tier: model.TierOne},
					{Name: "A2", Tier: model.TierOne},
					{Name: "B1", Tier: model.TierTwo},
					{Name: "B2", Tier: model.TierTwo},
					{Name: "C1", Tier: model.TierThree},
					{Name: "C2", Tier: model.TierThree},
				},
			},
			ProductOut: model.Product{
				Attributes: []model.Attribute{
					{
						Name:  "A1",
						Value: "X",
					},
					{
						Name:  "A2",
						Value: "",
					},
				},
				Bs: []model.B{
					{
						Attributes: []model.Attribute{
							{
								Name:  "B1",
								Value: "X",
							},
							{
								Name:  "B2",
								Value: "",
							},
						},
						Cs: []model.C{
							{
								Attributes: []model.Attribute{
									{
										Name:  "C1",
										Value: "X",
									},
									{
										Name:  "C2",
										Value: "",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	// act
	for i, tt := range testcases {
		expected := &tt.ProductOut
		AddRemoveAttributes(&tt.ProductIn, &tt.StructureIn)
		actual := &tt.ProductIn

		// assert
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("AddRemoveAttributes() failed!\ncase:%d\nexpected:%+v\n  actual:%+v\n", i, expected, actual)
		}
	}
}
