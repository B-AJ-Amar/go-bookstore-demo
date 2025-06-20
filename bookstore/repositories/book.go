package repositories

import (
	"errors"

	"github.com/B-AJ-Amar/go-bookstore-demo/bookstore/models"
	"gorm.io/gorm"
)

type BookRepository interface {
	Create(book *models.Book) error
	GetByID(id uint) (*models.Book, error)
	GetAll() ([]models.Book, error)
	Update(book *models.Book) error
	Delete(id uint) error
	UpdateBookQuantity(bookID uint, quantity int) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db}
}

func (r *bookRepository) Create(book *models.Book) error {
	if err := book.ValidateBook(); err != nil {
		return err
	}
	return r.db.Create(book).Error
}

func (r *bookRepository) GetByID(id uint) (*models.Book, error) {
	var book models.Book
	if err := r.db.Preload("Author").First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *bookRepository) GetAll() ([]models.Book, error) {
	var books []models.Book
	if err := r.db.Preload("Author").Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (r *bookRepository) Update(book *models.Book) error {
	if book.ID == 0 {
		return errors.New("book ID is required")
	}
	if err := book.ValidateBook(); err != nil {
		return err
	}
	return r.db.Save(book).Error
}

func (r *bookRepository) Delete(id uint) error {
	return r.db.Delete(&models.Book{}, id).Error
}


func (r *bookRepository) UpdateBookQuantity(bookID uint, quantity int) error {
	return r.db.Model(&models.Book{}).Where("id = ?", bookID).Update("quantity", quantity).Error
}