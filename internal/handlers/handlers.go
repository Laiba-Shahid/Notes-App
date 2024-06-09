package handlers

import (
	"net/http"
	"notes-app/internal/config"
	"notes-app/internal/models"
	"notes-app/internal/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "notes.page.tmpl", &models.TemplateData{})
}

// Login is the handler for the about page
func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	// send data to the template
	render.RenderTemplate(w, r, "login.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Login is the handler for the about page
func (m *Repository) SignUp(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	// send data to the template
	render.RenderTemplate(w, r, "signup.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Postlogin is the handler for the about page
func (m *Repository) PostLogin(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	w.Write([]byte("Posted to login"))

}
