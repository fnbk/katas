package main

import (
	"reflect"
	"testing"
)

func TestAddRemoveAttributes(t *testing.T) {
	testcases := []struct {
		ProductIn   Product
		StructureIn Structure
		ProductOut  Product
	}{
		// P_In:{} S:{B1} P_Out:{B1}
		{
			ProductIn: Product{
				Bs: []B{B{}},
			},
			StructureIn: Structure{
				Settings: []Setting{
					{Name: "B1", Tier: TierTwo},
				},
			},
			ProductOut: Product{
				Attributes: []Attribute{},
				Bs: []B{
					{
						Attributes: []Attribute{
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
			ProductIn: Product{
				Attributes: []Attribute{
					{
						Name:  "A1",
						Value: "X",
					},
				},
			},
			StructureIn: Structure{
				Settings: []Setting{
					{Name: "A1", Tier: TierOne},
				},
			},
			ProductOut: Product{
				Attributes: []Attribute{
					{
						Name:  "A1",
						Value: "X",
					},
				},
			},
		},

		// P_In:{A1X} S:{} P_Out:{}
		{
			ProductIn: Product{
				Attributes: []Attribute{
					{
						Name:  "A1",
						Value: "X",
					},
				},
			},
			StructureIn: Structure{},
			ProductOut:  Product{Attributes: []Attribute{}},
		},

		// P_In:{A1X,B1X} S:{B1} P_Out:{B1X}
		{
			ProductIn: Product{
				Attributes: []Attribute{
					{
						Name:  "A1",
						Value: "X",
					},
				},
				Bs: []B{
					{
						Attributes: []Attribute{
							{
								Name:  "B1",
								Value: "X",
							},
						},
					},
				},
			},
			StructureIn: Structure{
				Settings: []Setting{
					{Name: "B1", Tier: TierTwo},
				},
			},
			ProductOut: Product{
				Attributes: []Attribute{},
				Bs: []B{
					{
						Attributes: []Attribute{
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
			ProductIn: Product{
				Attributes: []Attribute{
					{
						Name:  "A1",
						Value: "X",
					},
				},
			},
			StructureIn: Structure{
				Settings: []Setting{
					{Name: "A1", Tier: TierOne},
					{Name: "A2", Tier: TierOne},
				},
			},
			ProductOut: Product{
				Attributes: []Attribute{
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
			ProductIn: Product{
				Attributes: []Attribute{
					{
						Name:  "A1",
						Value: "X",
					},
				},
				Bs: []B{
					{
						Attributes: []Attribute{
							{
								Name:  "B1",
								Value: "X",
							},
						},
						Cs: []C{
							{
								Attributes: []Attribute{
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
			StructureIn: Structure{
				Settings: []Setting{
					{Name: "A1", Tier: TierOne},
					{Name: "A2", Tier: TierOne},
					{Name: "B1", Tier: TierTwo},
					{Name: "B2", Tier: TierTwo},
					{Name: "C1", Tier: TierThree},
					{Name: "C2", Tier: TierThree},
				},
			},
			ProductOut: Product{
				Attributes: []Attribute{
					{
						Name:  "A1",
						Value: "X",
					},
					{
						Name:  "A2",
						Value: "",
					},
				},
				Bs: []B{
					{
						Attributes: []Attribute{
							{
								Name:  "B1",
								Value: "X",
							},
							{
								Name:  "B2",
								Value: "",
							},
						},
						Cs: []C{
							{
								Attributes: []Attribute{
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
