package core

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockOrderRepository struct {
	saved   *Order
	saveErr error
}

func (m *mockOrderRepository) Save(order *Order) error {
	m.saved = order
	return m.saveErr
}

func TestCreateOrder_Success(t *testing.T) {
	repo := &mockOrderRepository{}
	service := NewOrderService(repo)

	order := Order{ID: 1, Total: 100.0}
	err := service.CreateOrder(order)

	assert.NoError(t, err)
	assert.NotNil(t, repo.saved)
	assert.Equal(t, order.ID, repo.saved.ID)
	assert.Equal(t, order.Total, repo.saved.Total)
}

func TestCreateOrder_RepositoryError(t *testing.T) {
	repo := &mockOrderRepository{saveErr: errors.New("save failed")}
	service := NewOrderService(repo)

	err := service.CreateOrder(Order{ID: 1, Total: 100.0})

	assert.Error(t, err)
	assert.EqualError(t, err, "save failed")
}
