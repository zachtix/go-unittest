package core

type OrderService interface {
	CreateOrder(order Order) error
}

type orderService struct {
	repo OrderRepository
}

func NewOrderService(repo OrderRepository) OrderService {
	return &orderService{repo: repo}
}

func (s *orderService) CreateOrder(order Order) error {
	return s.repo.Save(&order)
}
