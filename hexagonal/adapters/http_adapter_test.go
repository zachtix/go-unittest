package adapters

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	"zachtix/hexagonal/core"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockOrderService struct {
	mock.Mock
}

func (m *MockOrderService) CreateOrder(order core.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func TestCreateOrderHandle(t *testing.T) {
	mockService := new(MockOrderService)
	handler := NewHttpOrderHandler(mockService)

	app := fiber.New()
	app.Post("/orders", handler.CreateOrder)

	t.Run("Successful order creation", func(t *testing.T) {
		mockService.On("CreateOrder", mock.AnythingOfType("core.Order")).Return(nil)

		req := httptest.NewRequest("POST", "/orders", bytes.NewBufferString(`{"total":100}`))
		req.Header.Set("Content-Type", "application/json")
		resq, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resq.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("Fail order creation (total less than 0)", func(t *testing.T) {
		mockService.ExpectedCalls = nil
		mockService.On("CreateOrder", mock.AnythingOfType("core.Order")).Return(errors.New("total must positive"))

		req := httptest.NewRequest("POST", "/orders", bytes.NewBufferString(`{"total":-200}`))
		req.Header.Set("Content-Type", "application/json")
		resq, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resq.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("Invalid request body", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/orders", bytes.NewBufferString(`{"total":"invalid"}`))
		req.Header.Set("Content-Type", "application/json")
		resq, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resq.StatusCode)
	})

	t.Run("Order service error", func(t *testing.T) {
		mockService.ExpectedCalls = nil
		mockService.On("CreateOrder", mock.AnythingOfType("core.Order")).Return(errors.New("service error"))

		req := httptest.NewRequest("POST", "/orders", bytes.NewBufferString(`{"total":100}`))
		req.Header.Set("Content-Type", "application/json")
		resq, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resq.StatusCode)
		mockService.AssertExpectations(t)
	})
}
