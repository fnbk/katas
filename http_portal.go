package main

import (
	"net/http"
	"regexp"
)

var showProductPath = regexp.MustCompile("^/products/([a-zA-Z0-9]+)$")
var listProductIDsPath = regexp.MustCompile("^/products/$")

type HTTPPortal struct {
	ProductProvider   ProductProvider
	StructureProvider StructureProvider
}

func (p *HTTPPortal) ServeHTTP(writer http.ResponseWriter, reader *http.Request) {
	var job Job
	p.route(writer, reader, &job)
	job.Execute()
}

func (p *HTTPPortal) route(writer http.ResponseWriter, reader *http.Request, job *Job) {
	m1 := showProductPath.FindStringSubmatch(reader.URL.Path)
	m2 := listProductIDsPath.FindStringSubmatch(reader.URL.Path)
	if m1 != nil {
		*job = &ShowProductJob{reader: reader, writer: writer, ProductProvider: p.ProductProvider, StructureProvider: p.StructureProvider, ShowProductPath: showProductPath}
	} else if m2 != nil {
		*job = &ListProductIDsJob{reader: reader, writer: writer, ProductProvider: p.ProductProvider}
	} else {
		*job = &NotFoundJob{reader: reader, writer: writer}
	}
}
