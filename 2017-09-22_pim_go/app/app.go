package app

import (
	"net/http"

	"github.com/fnbk/pim/app/portal"
)

type App struct {
	portal.HTTPPortal
}

func (a *App) Run() {
	http.ListenAndServe(":8080", &a.HTTPPortal)
}
