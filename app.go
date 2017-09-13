package main

import (
	"encoding/json"
	"net/http"
)

type App struct {
	ProductProvider ProductProvider
}

func (a *App) Run() {
	http.HandleFunc("/products/", a.ListProductIDsHandler)
	http.ListenAndServe(":8080", nil)
}
func (a *App) ListProductIDsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ids := a.ProductProvider.ProductIDs()
	json.NewEncoder(w).Encode(ids)
}
