package handlers

import (
	"net/http"

	"github.com/xinglixing/book-go/config"
	"github.com/xinglixing/book-go/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(rw http.ResponseWriter, r *http.Request) {
	sm := make(map[string]string)
	sm["test"] = "Hello there"

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(rw, "home.page.tmpl", &config.TemplateData{
		StringMap: sm,
	})
}

func (m *Repository) About(rw http.ResponseWriter, r *http.Request) {
	sm := make(map[string]string)

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	sm["remote_ip"] = remoteIP
	render.RenderTemplate(rw, "about.page.tmpl", &config.TemplateData{
		StringMap: sm,
	})
}
