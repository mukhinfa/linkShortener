package link

// CreateRequest is a struct with Create body
type CreateRequest struct {
	URL string `json:"url" validate:"required,url"`
}
