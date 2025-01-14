package link

import (
	"github.com/muhinfa/linkShortener/pkg/db"
	"gorm.io/gorm/clause"
)

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

// GetByHash retrieves a link by its hash
func (repo *Repository) GetByHash(hash string) (*Link, error) {
	var link Link
	result := repo.DataBase.DB.First(&link, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

// Update updates a link in the repository
func (repo *Repository) Update(link *Link) (*Link, error) {
	result := repo.DataBase.DB.Clauses(clause.Returning{}).Updates(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

// Delete deletes a link from the repository
func (repo *Repository) Delete(id uint) error {
	result := repo.DataBase.DB.Delete(&Link{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// getByID retrieves a link by its ID
func (repo *Repository) getByID(id uint) (*Link, error) {
	var link Link
	result := repo.DataBase.DB.First(&link, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}
