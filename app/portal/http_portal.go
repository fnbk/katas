package portal

import (
	"encoding/json"
	"net/http"
	"regexp"

	"bitbucket.scm.otto.de/scm/primary/pim/app/core"
	"bitbucket.scm.otto.de/scm/primary/pim/app/model"
)

var showProductPath = regexp.MustCompile("^/products/([a-zA-Z0-9]+)$")
var listProductIDsPath = regexp.MustCompile("^/products/$")

type HTTPPortal struct {
	core.Core
}

func (p *HTTPPortal) ServeHTTP(writer http.ResponseWriter, reader *http.Request) {
	handleNotFound := func() {
		p.HandleNotFound(writer, reader)
	}
	handleListProductIDs := func() {
		p.HandleListProductIDs(writer, reader)
	}
	handleShowProduct := func() {
		p.HandleShowProduct(writer, reader)
	}
	p.route(reader, handleNotFound, handleListProductIDs, handleShowProduct)
}

func (p *HTTPPortal) route(reader *http.Request, handleNotFound, handleListProductIDs, handleShowProduct func()) {
	showProductPathRegex := showProductPath.FindStringSubmatch(reader.URL.Path)
	listProductIDsPathRegex := listProductIDsPath.FindStringSubmatch(reader.URL.Path)
	if showProductPathRegex != nil {
		handleShowProduct()
	} else if listProductIDsPathRegex != nil {
		handleListProductIDs()
	} else {
		handleNotFound()
	}
}

func (p *HTTPPortal) HandleNotFound(writer http.ResponseWriter, reader *http.Request) {
	http.NotFound(writer, reader)
}

func (p *HTTPPortal) HandleListProductIDs(writer http.ResponseWriter, reader *http.Request) {
	ids := p.Core.ListProductIDs()
	p.writeListProductIDsResponse(writer, ids)
}

func (p *HTTPPortal) writeListProductIDsResponse(writer http.ResponseWriter, ids []string) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(ids)
}

func (p *HTTPPortal) HandleShowProduct(writer http.ResponseWriter, reader *http.Request) {
	id := p.parseProductID(reader)
	product := p.Core.ShowProduct(id)
	p.writeShowProductResponse(writer, product)
}

func (p *HTTPPortal) parseProductID(reader *http.Request) string {
	m := showProductPath.FindStringSubmatch(reader.URL.Path)
	return m[1]
}

func (p *HTTPPortal) writeShowProductResponse(writer http.ResponseWriter, product *model.Product) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(product)
}
