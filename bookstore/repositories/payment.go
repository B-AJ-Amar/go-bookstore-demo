package repositories

import (
	"errors"

	"github.com/B-AJ-Amar/go-bookstore-demo/bookstore/models"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(payment *models.Payment) error
	GetByID(id uint) (*models.Payment, error)
	GetAll() ([]models.Payment, error)
	Update(payment *models.Payment) error
	Delete(id uint) error
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db}
}

func (r *paymentRepository) Create(payment *models.Payment) error {
	if err := payment.ValidatePayment(); err != nil {
		return err
	}
	return r.db.Create(payment).Error
}

func (r *paymentRepository) GetByID(id uint) (*models.Payment, error) {
	var payment models.Payment
	if err := r.db.First(&payment, id).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepository) GetAll() ([]models.Payment, error) {
	var payments []models.Payment
	if err := r.db.Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}

func (r *paymentRepository) Update(payment *models.Payment) error {
	if payment.ID == 0 {
		return errors.New("payment ID is required")
	}
	if err := payment.ValidatePayment(); err != nil {
		return err
	}
	return r.db.Save(payment).Error
}

func (r *paymentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Payment{}, id).Error
}
