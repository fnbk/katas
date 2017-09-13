package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/products/", ListProductsHandler)
	http.ListenAndServe(":8080", nil)
}

func ListProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ids := ProductIDs()
	json.NewEncoder(w).Encode(ids)
}

func ProductIDs() []string {
	return []string{"1", "2"}
}
