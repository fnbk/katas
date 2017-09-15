package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("visit: http:localhost:8080/products/\n")
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
