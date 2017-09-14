package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
)

type apiFeature struct {
	// app  App
	HTTPPortal
	resp *httptest.ResponseRecorder
}

func (a *apiFeature) resetResponse(interface{}) {
	a.resp = httptest.NewRecorder()
}

func (a *apiFeature) thereAreIDsInTheProductDatabase(IDTable *gherkin.DataTable) error {
	products := []Product{}
	for i, row := range IDTable.Rows {
		if i == 0 {
			continue
		}
		products = append(products, Product{ID: row.Cells[0].Value})
	}
	productProvider := ProductProvider{Products: products}
	a.HTTPPortal = HTTPPortal{ProductProvider: productProvider}
	// a.app = App{portal}
	// a.app = App{ProductProvider: provider}
	return nil
}

func (a *apiFeature) iSendRequestTo(method, endpoint string) (err error) {
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

	switch endpoint {
	case "/products":
		// func (a *App) ListProductIDsHandler(w http.ResponseWriter, r *http.Request) {
		// a.app.ListProductIDsHandler(a.resp, req)
		job := ListProductIDsJob{a.resp, req, a.HTTPPortal.ProductProvider}
		job.Execute()
	default:
		err = fmt.Errorf("unknown endpoint: %s", endpoint)
	}
	return
}

func (a *apiFeature) theResponseCodeShouldBe(code int) error {
	if code != a.resp.Code {
		if a.resp.Code >= 400 {
			return fmt.Errorf("expected response code to be: %d, but actual is: %d, response message: %s", code, a.resp.Code, string(a.resp.Body.Bytes()))
		}
		return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, a.resp.Code)
	}
	return nil
}

func (a *apiFeature) theResponseShouldMatchJson(body *gherkin.DocString) (err error) {
	var expectedBlob, actualBlob []byte
	var expected, actual string
	var data interface{}
	if err = json.Unmarshal([]byte(body.Content), &data); err != nil {
		return
	}
	if expectedBlob, err = json.Marshal(data); err != nil {
		return
	}
	actualBlob = a.resp.Body.Bytes()
	actual = strings.TrimSpace(string(actualBlob))
	expected = strings.TrimSpace(string(expectedBlob))

	if actual != expected {
		err = fmt.Errorf("expected json %s, does not match actual: %s", expected, actual)
	}
	return
}

func FeatureContext(s *godog.Suite) {
	api := &apiFeature{}

	s.BeforeScenario(api.resetResponse)

	s.Step(`^there are IDs in the product database:$`, api.thereAreIDsInTheProductDatabase)
	s.Step(`^I send "([^"]*)" request to "([^"]*)"$`, api.iSendRequestTo)
	s.Step(`^the response code should be (\d+)$`, api.theResponseCodeShouldBe)
	s.Step(`^the response should match json:$`, api.theResponseShouldMatchJson)
}
