// Package app ties together application resources and handlers.
package app

import (
	"net/http"
	
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"

	"github.com/dhax/go-base/logging"
)

type ctxKey int

const (
	ctxAccount ctxKey = iota
	ctxProfile
)

// API provides application resources and handlers.
type API struct {
	Links *LinksResource
}

// NewAPI configures and returns application API.
func NewAPI() (*API, error) {
	links := NewLinksResource()

	api := &API{
		Links: links,
	}
	return api, nil
}

// Router provides application routes.
func (a *API) Router() *chi.Mux {
	r := chi.NewRouter()

	//	r.Group(func(r chi.Router) {
	//		r.Use(jwt.Authenticator)
	//		r.Mount("/profile", a.Profile.router())
	//	})

	r.Group(func(r chi.Router) {
		r.Mount("/links", a.Links.router())
	})

	return r
}

func log(r *http.Request) logrus.FieldLogger {
	return logging.GetLogEntry(r)
}
