package main

import (
	"encoding/json"
	"net/http"
	"regexp"
)

type Product struct {
	ID          string
	Name        string
	StructureID string
}

func main() {
	http.HandleFunc("/", router)
	http.ListenAndServe(":8080", nil)
}

var showProductPath = regexp.MustCompile("^/products/([a-zA-Z0-9]+)$")

func router(w http.ResponseWriter, r *http.Request) {
	m := showProductPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return
	}
	ShowProductHandler(w, r, m[1])
}

func ShowProductHandler(w http.ResponseWriter, r *http.Request, id string) {
	product := ShowProduct(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func ShowProduct(id string) *Product {
	return &Product{
		ID:          id,
		Name:        "myName",
		StructureID: "myStructureID",
	}
}
