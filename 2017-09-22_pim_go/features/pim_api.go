package features

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/fnbk/pim/app/core"
	"github.com/fnbk/pim/app/model"
	"github.com/fnbk/pim/app/portal"
	"github.com/fnbk/pim/app/provider"

	"github.com/DATA-DOG/godog/gherkin"
)

type pimAPI struct {
	portal.HTTPPortal
	resp *httptest.ResponseRecorder
}

func (a *pimAPI) resetResponse(interface{}) {
	productProvider := provider.ProductProvider{Products: []model.Product{}}
	structureProvider := provider.StructureProvider{Structures: []model.Structure{}}
	c := core.Core{ProductProvider: productProvider, StructureProvider: structureProvider}
	a.HTTPPortal = portal.HTTPPortal{Core: c}

	a.resp = httptest.NewRecorder()
}

func (a *pimAPI) thereAreIDsInTheProductDatabase(IDTable *gherkin.DataTable) error {
	products := []model.Product{}
	for i, row := range IDTable.Rows {
		if i == 0 {
			continue
		}
		products = append(products, model.Product{ID: row.Cells[0].Value})
	}
	a.HTTPPortal.Core.ProductProvider.Products = products
	return nil
}

func (a *pimAPI) aProductWithTheIDAndStructureIDExists(productID, structureID string) error {
	a.HTTPPortal.Core.ProductProvider.Products = []model.Product{{ID: productID, StructureID: structureID}}
	return nil
}

func (a *pimAPI) aStructureWithTheIDExists(id string) error {
	a.HTTPPortal.Core.StructureProvider.Structures = []model.Structure{{ID: id}}
	return nil
}

func (a *pimAPI) productAndStructureWithAttributesAndSettingsInTheDatabase(table *gherkin.DataTable) error {
	for i, row := range table.Rows {
		if i == 0 {
			continue
		}
		if row.Cells[0].Value != "" {
			attr := strings.Split(row.Cells[0].Value, ",")
			if attr[0][0] == 'A' {
				a.HTTPPortal.Core.ProductProvider.Products[0].Attributes = append(a.HTTPPortal.Core.ProductProvider.Products[0].Attributes, model.Attribute{Name: attr[0], Value: attr[1]})
			}
			if attr[0][0] == 'B' {
				if len(a.HTTPPortal.Core.ProductProvider.Products[0].Bs) == 0 {
					a.HTTPPortal.Core.ProductProvider.Products[0].Bs = []model.B{{Attributes: []model.Attribute{}}}
				}
				a.HTTPPortal.Core.ProductProvider.Products[0].Bs[0].Attributes = append(a.HTTPPortal.Core.ProductProvider.Products[0].Bs[0].Attributes, model.Attribute{Name: attr[0], Value: attr[1]})
			}
			if attr[0][0] == 'C' {
				if len(a.HTTPPortal.Core.ProductProvider.Products[0].Bs[0].Cs) == 0 {
					a.HTTPPortal.Core.ProductProvider.Products[0].Bs[0].Cs = []model.C{{Attributes: []model.Attribute{}}}
				}
				a.HTTPPortal.Core.ProductProvider.Products[0].Bs[0].Cs[0].Attributes = append(a.HTTPPortal.Core.ProductProvider.Products[0].Bs[0].Cs[0].Attributes, model.Attribute{Name: attr[0], Value: attr[1]})
			}
		}
		if row.Cells[1].Value != "" {
			attr := strings.Split(row.Cells[1].Value, ",")
			tier := model.TierOne
			if attr[0][0] == 'A' {
				tier = model.TierOne
			} else if attr[0][0] == 'B' {
				tier = model.TierTwo
			} else if attr[0][0] == 'C' {
				tier = model.TierThree
			}
			a.HTTPPortal.Core.StructureProvider.Structures[0].Settings = append(a.HTTPPortal.Core.StructureProvider.Structures[0].Settings, model.Setting{Name: attr[0], Tier: tier})
		}
	}
	return nil
}

func (a *pimAPI) iSendRequestTo(method, endpoint string) (err error) {
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return
	}

	// handle panic
	defer func() {
		switch t := recover().(type) {
		case string:
			err = fmt.Errorf(t)
		case error:
			err = t
		}
	}()

	a.HTTPPortal.ServeHTTP(a.resp, req)
	return
}

func (a *pimAPI) theResponseCodeShouldBe(code int) error {
	if code != a.resp.Code {
		if a.resp.Code >= 400 {
			return fmt.Errorf("expected response code to be: %d, but actual is: %d, response message: %s", code, a.resp.Code, string(a.resp.Body.Bytes()))
		}
		return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, a.resp.Code)
	}
	return nil
}

func (a *pimAPI) theResponseShouldMatchJson(body *gherkin.DocString) (err error) {
	var expectedBlob, actualBlob []byte
	var expected, actual string
	var data interface{}

	if err = json.Unmarshal([]byte(body.Content), &data); err != nil {
		return
	}
	if expectedBlob, err = json.Marshal(data); err != nil {
		return
	}
	expected = strings.TrimSpace(string(expectedBlob))

	if err = json.Unmarshal(a.resp.Body.Bytes(), &data); err != nil {
		return
	}
	if actualBlob, err = json.Marshal(data); err != nil {
		return
	}
	actual = strings.TrimSpace(string(actualBlob))

	if actual != expected {
		err = fmt.Errorf("json does not match!\nexpected:\n%s\nactual:\n%s", expected, actual)
	}
	return
}
