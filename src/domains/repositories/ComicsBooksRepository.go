package repositories

import (
	"errors"
	"github.com/VieiraGabrielAlexandre/api_comics_and_books/src/config/database"
	"github.com/VieiraGabrielAlexandre/api_comics_and_books/src/models"
	"gorm.io/gorm"
)

type ComicsBooksRepository interface {
	Create(comics models.Comics) (models.Comics, error)
	Update(comics models.Comics, id int) (models.Comics, error)
	GetByID(id int) (models.Comics, error)
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

func (repo ComicsBooksRepositoryDb) Update(comics models.Comics, id int) (models.Comics, error) {
	result := database.DB.
		Model(&comics).
		Assign(comics).
		Where("id = ?", id).
		Updates(&comics)

	if result.RowsAffected == 0 {
		return comics, errors.New("Not found")
	}

	if result.Error != nil {
		return comics, result.Error
	}

	return comics, nil
}

func (repo ComicsBooksRepositoryDb) GetByID(id int) (models.Comics, error) {
	var comics models.Comics

	result := database.DB.
		Model(&models.Comics{}).
		Where("id = ?", id).
		First(&comics)

	if result.Error != nil {
		return comics, result.Error
	}

	return comics, nil
}
