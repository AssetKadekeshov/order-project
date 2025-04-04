package service

import (
	"rest-project/internal/models"
)

type OrderRepository interface {
	GetAll() ([]models.Order, error)
	GetById(id int) (*models.Order, error)
	Create(order *models.Order) error
	Update(id int, order *models.OrderEdit) error
	Delete(orderID int) error
}

type OrderService struct {
	repo OrderRepository
}

func NewOrderService(orderRepo OrderRepository) *OrderService {
	return &OrderService{repo: orderRepo}
}

func (s *OrderService) GetAllOrders() ([]models.Order, error) {
	return s.repo.GetAll()
}

func (s *OrderService) GetOrderByID(id int) (*models.Order, error) {
	return s.repo.GetById(id)
}

func (s *OrderService) Create(customerName, product string, quantity int, status string) (*models.Order, error) {
	order := &models.Order{
		CustomerName: customerName,
		Product:      product,
		Quantity:     quantity,
		Status:       status,
	}
	err := s.repo.Create(order)
	return order, err
}

func (s *OrderService) Update(id int, orderEdit *models.OrderEdit) (*models.Order, error) {
	err := s.repo.Update(id, orderEdit)
	if err != nil {
		return nil, err
	}
	return s.GetOrderByID(id)
}

func (s *OrderService) DeleteOrder(orderID int) error {
	return s.repo.Delete(orderID)
}
