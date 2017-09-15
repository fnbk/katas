package app

import (
	"net/http"

	"bitbucket.scm.otto.de/scm/primary/pim/app/portal"
)

type App struct {
	portal.HTTPPortal
}

func (a *App) Run() {
	http.ListenAndServe(":8080", &a.HTTPPortal)
}
