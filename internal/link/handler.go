package link

import (
	"fmt"
	"net/http"
)

// LinkHandlerDeps contains dependencies for the link handler
type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
}

// LinkHandler processes requests related to links
type LinkHandler struct {
	LinkRepository *LinkRepository
}

// NewLinkHandler creates a new link handler and registers it with the router
func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
	}
	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
	router.HandleFunc("GET /{hash}", handler.GoTo())
}

// Create handles the creation of a new link
func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Update handles updating an existing link
func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

// Delete handles deleting an existing link
func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Println(id)
	}
}

// GoTo handles redirection by hash
func (handler *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
