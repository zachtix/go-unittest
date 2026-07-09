package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockOrderRepo struct {
	saveFunc func(order Order) error
}

func (m *mockOrderRepo) Save(order Order) error {
	return m.saveFunc(order)
}

func TestCreateOrder(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		repo := &mockOrderRepo{
			saveFunc: func(order Order) error {
				return nil
			},
		}
		service := NewOrderService(repo)

		err := service.CreateOrder(Order{Total: 100})
		assert.NoError(t, err)
	})

	t.Run("Total less than 0", func(t *testing.T) {
		repo := &mockOrderRepo{
			saveFunc: func(order Order) error { return nil },
		}
		service := NewOrderService(repo)

		err := service.CreateOrder(Order{Total: -10})
		assert.Error(t, err)
		assert.Equal(t, "total must be positive", err.Error())
	})
}
