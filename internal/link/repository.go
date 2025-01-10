package link

import "github.com/muhinfa/linkShortener/pkg/db"

// Repository represents a repository for working with links in a database
type Repository struct {
	DataBase *db.Db
}

// NewLinkRepository creates a new LinkRepository instance with the given database
func NewLinkRepository(database *db.Db) *Repository {
	return &Repository{
		DataBase: database,
	}
}

// Create adds a new link to the repository
func (repo *Repository) Create(link *Link) (*Link, error) {
	result := repo.DataBase.DB.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}
