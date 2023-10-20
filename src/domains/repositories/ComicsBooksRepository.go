package repositories

import (
	"github.com/VieiraGabrielAlexandre/api_comics_and_books/src/config/database"
	"github.com/VieiraGabrielAlexandre/api_comics_and_books/src/models"
	"gorm.io/gorm"
)

type ComicsBooksRepository interface {
	Create(comics models.Comics) (models.Comics, error)
}

type ComicsBooksRepositoryDb struct {
	Db *gorm.DB
}

func (repo ComicsBooksRepositoryDb) Create(comics models.Comics) (models.Comics, error) {
	err := database.DB.
		Where(models.Comics{Title: comics.Title, Volume: comics.Volume}).
		Assign(comics).
		FirstOrCreate(&comics).
		Error

	if err != nil {
		return comics, err
	}

	return comics, nil
}
