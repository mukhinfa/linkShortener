package link

// CreateRequest is a struct with Create body
type CreateRequest struct {
	URL string `json:"url" validate:"required,url"`
}

// UpdateRequest is a struct with Update body
type UpdateRequest struct {
	URL  string `json:"url" validate:"required,url"`
	Hash string `json:"hash"`
}
