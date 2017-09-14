package main

import (
	"encoding/json"
	"net/http"
	"regexp"
)

type ShowProductJob struct {
	reader *http.Request
	writer http.ResponseWriter
	ProductProvider
	StructureProvider
	ShowProductPath *regexp.Regexp
}

func (job *ShowProductJob) Execute() {
	id := job.parseID(job.reader.URL.Path)
	product := job.showProduct(id)
	job.writeResponse(product)
}

func (job *ShowProductJob) parseID(url string) string {
	m := job.ShowProductPath.FindStringSubmatch(url)
	return m[1]
}

func (job *ShowProductJob) showProduct(id string) *Product {
	product := job.ProductProvider.GetProduct(id)
	structure := job.StructureProvider.GetStructure(id)
	AddRemoveAttributes(product, structure)
	return product
}

func (job *ShowProductJob) writeResponse(product *Product) {
	job.writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(job.writer).Encode(product)
}
