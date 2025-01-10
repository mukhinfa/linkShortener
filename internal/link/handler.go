package link

import (
	"fmt"
	"net/http"

	"github.com/muhinfa/linkShortener/pkg/req"
	"github.com/muhinfa/linkShortener/pkg/res"
)

// HandlerDeps contains dependencies for the link handler
type HandlerDeps struct {
	Repository *Repository
}

// Handler processes requests related to links
type Handler struct {
	Repository *Repository
}

// NewLinkHandler creates a new link handler and registers it with the router
func NewLinkHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := &Handler{
		Repository: deps.Repository,
	}
	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
	router.HandleFunc("GET /{hash}", handler.GoTo())
}

// Create handles the creation of a new link
func (handler *Handler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[CreateRequest](&w, r)
		if err != nil {
			return
		}
		link := NewLink(body.URL)
		createdLink, err := handler.Repository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.JSON(w, createdLink, http.StatusCreated)
	}
}

// Update handles updating an existing link
func (handler *Handler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

// Delete handles deleting an existing link
func (handler *Handler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Println(id)
	}
}

// GoTo handles redirection by hash
func (handler *Handler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
