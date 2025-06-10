package repository

import (
	"errors"

	"github.com/B-AJ-Amar/go-bookstore-demo/bookstore/models"
	"gorm.io/gorm"
)

type AuthorRepository interface {
	Create(author *models.Author) error
	GetByID(id uint) (*models.Author, error)
	GetAll() ([]models.Author, error)
	Update(author *models.Author) error
	Delete(id uint) error
}

type authorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) AuthorRepository {
	return &authorRepository{db}
}

func (r *authorRepository) Create(author *models.Author) error {
	if err := author.ValidateAuthor(); err != nil {
		return err
	}
	return r.db.Create(author).Error
}

func (r *authorRepository) GetByID(id uint) (*models.Author, error) {
	var author models.Author
	if err := r.db.Preload("Books").First(&author, id).Error; err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *authorRepository) GetAll() ([]models.Author, error) {
	var authors []models.Author
	if err := r.db.Preload("Books").Find(&authors).Error; err != nil {
		return nil, err
	}
	return authors, nil
}

func (r *authorRepository) Update(author *models.Author) error {
	if author.ID == 0 {
		return errors.New("author ID is required")
	}
	if err := author.ValidateAuthor(); err != nil {
		return err
	}
	return r.db.Save(author).Error
}

func (r *authorRepository) Delete(id uint) error {
	return r.db.Delete(&models.Author{}, id).Error
}