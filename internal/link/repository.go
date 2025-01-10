package link

import "github.com/muhinfa/linkShortener/pkg/db"

// LinkRepository represents a repository for working with links in a database
type LinkRepository struct {
	DataBase *db.Db
}

// NewLinkRepository creates a new LinkRepository instance with the given database
func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository{
		DataBase: database,
	}
}

// Create adds a new link to the repository
func (repo *LinkRepository) Create(link *Link) {

}
