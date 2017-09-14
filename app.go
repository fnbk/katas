package main

import (
	"net/http"
)

type App struct {
	HTTPPortal
}

func (a *App) Run() {
	http.ListenAndServe(":8080", &a.HTTPPortal)
}
