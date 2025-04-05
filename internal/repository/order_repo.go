package repository

import (
	"gorm.io/gorm"
	"rest-project/internal/models"
)

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{db: db}
}

func (o OrderRepositoryImpl) GetAll() ([]models.Order, error) {
	var orders []models.Order
	err := o.db.Find(&orders).Error
	return orders, err
}

func (o OrderRepositoryImpl) GetById(id int) (*models.Order, error) {
	var order models.Order
	err := o.db.First(&order, id).Error
	return &order, err
}

func (o OrderRepositoryImpl) Create(order *models.Order) error {
	return o.db.Create(order).Error
}

func (o OrderRepositoryImpl) Update(id int, order *models.OrderEdit) error {
	return o.db.Model(&models.Order{}).Where("id = ?", id).Omit("id, CreatedAt").Updates(order).Error
}

func (o OrderRepositoryImpl) Delete(orderID int) error {
	return o.db.Delete(&models.Order{}, orderID).Error
}

func (o OrderRepositoryImpl) Filter(status, product, search string) ([]models.Order, error) {
	var orders []models.Order
	query := o.db.Model(&models.Order{})

	if status != "" {
		query = query.Where("status = ?", status)
	}
	if product != "" {
		query = query.Where("product = ?", product)
	}
	if search != "" {
		query = query.Where("customer_name ILIKE ?", "%"+search+"%")
	}

	err := query.Find(&orders).Error
	return orders, err
}
