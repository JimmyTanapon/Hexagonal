package adepters

import (
	"bytes"
	"errors"

	"github.com/JimmyTanapon/Hexagonal/core"

	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
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

func TestCreateOrderHandler(t *testing.T) {
	mockService := new(MockOrderService)
	handler := NewHttpOrderHandler(mockService)

	app := fiber.New()
	app.Post("/orders", handler.CreateOrder)

	t.Run("successful order creation", func(t *testing.T) {
		mockService.On("CreateOrder", mock.AnythingOfType("core.Order")).Return(nil)
		req := httptest.NewRequest("POST", "/orders", bytes.NewBufferString(`{"total":1000}`))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
		mockService.AssertExpectations(t)

	})
	t.Run("fail order creation (total less than 0)", func(t *testing.T) {
		mockService.ExpectedCalls = nil
		mockService.On("CreateOrder", mock.AnythingOfType("core.Order")).Return(errors.New("total must be positive"))
		req := httptest.NewRequest("POST", "/orders", bytes.NewBufferString(`{"total":-100}`))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("invalid request body", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/orders", bytes.NewBufferString(`{"total": "invalid"}`))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("order service error", func(t *testing.T) {
		mockService.ExpectedCalls = nil
		mockService.On("CreateOrder", mock.AnythingOfType("core.Order")).Return(errors.New("service error"))
		req := httptest.NewRequest("POST", "/orders", bytes.NewBufferString(`{"total":100}`))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

}
