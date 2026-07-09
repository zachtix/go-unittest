package adapters

import (
	"net/http/httptest"
	"strings"
	"testing"

	"zachtix/hexagonal/core"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

type mockOrderService struct {
	createErr error
}

func (m *mockOrderService) CreateOrder(order core.Order) error {
	return m.createErr
}

func TestHttpOrderHandler_CreateOrder_Success(t *testing.T) {
	app := fiber.New()
	handler := NewHttpOrderHandler(&mockOrderService{})
	app.Post("/order", handler.CreateOrder)

	req := httptest.NewRequest("POST", "/order", strings.NewReader(`{"ID":1,"Total":100}`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
}

func TestHttpOrderHandler_CreateOrder_InvalidBody(t *testing.T) {
	app := fiber.New()
	handler := NewHttpOrderHandler(&mockOrderService{})
	app.Post("/order", handler.CreateOrder)

	req := httptest.NewRequest("POST", "/order", strings.NewReader(`{invalid`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}
