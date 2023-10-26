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
	GetAllPaginate(page int, limit int, filter string, orderBy string, column string, order string) ([]models.Comics, error)
	Delete(id int) error
	Count() (int64, error)
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

func (repo ComicsBooksRepositoryDb) GetAllPaginate(page int, limit int, filter string, orderBy string, column string, order string) ([]models.Comics, error) {
	var comics []models.Comics

	orderQuery := "id"

	if orderBy != "" {
		orderQuery = orderBy
	}

	result := database.DB.
		Order(orderQuery + " " + order).
		Model(&models.Comics{}).
		Offset((page - 1) * limit).
		Limit(limit)
	if filter != "" {
		result = result.Where(column+" LIKE ?", "%"+filter+"%")
	}

	result = result.Find(&comics)

	if result.Error != nil {
		return comics, result.Error
	}

	return comics, nil
}

func (repo ComicsBooksRepositoryDb) Delete(id int) error {
	result := database.DB.
		Model(&models.Comics{}).
		Where("id = ?", id).
		Delete(&models.Comics{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo ComicsBooksRepositoryDb) Count() (int64, error) {
	count := int64(0)
	err := repo.Db.Model(&models.Comics{}).Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}
