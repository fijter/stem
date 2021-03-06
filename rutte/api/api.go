package api

import (
	"net/http"

	"github.com/GeenPeil/stem/rutte/version"

	"github.com/Sirupsen/logrus"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/pressly/chi/render"
)

// API service for cpserver
type API struct {
	log *logrus.Entry
	db  *sqlx.DB
}

// New creates a new instance of the API service
func New(log *logrus.Entry, db *sqlx.DB) *API {
	return &API{
		log: log,
		db:  db,
	}
}

// AttachChiRouter attaches API routes to the provided chi router.
func (a *API) AttachChiRouter(r chi.Router) {
	a.log.Infoln("attach chi router")
	r.Use(middleware.NoCache)
	r.Get("/", a.getRoot)
	r.Post("/login", a.login())
	r.Route("/private", func(r chi.Router) {
		r.Use(a.newAccountAuthMiddleware())
		r.Get("/profile", a.getProfileOverview())
	})
}

// getRoot handles get requests to root (show api version?)
func (a *API) getRoot(w http.ResponseWriter, r *http.Request) {
	out := map[string]string{
		"apiVersion": "v0.0",
		"appVersion": version.String(),
	}
	render.JSON(w, r, out)
}
