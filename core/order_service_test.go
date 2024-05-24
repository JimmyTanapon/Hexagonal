package core

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockOrderRepo struct {
	SaveFunc func(order Order) error
}

func (m *mockOrderRepo) SaveOrder(order Order) error {
	return m.SaveFunc(order)
}

func TestCreateOrder(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		repo := &mockOrderRepo{
			SaveFunc: func(order Order) error {
				return nil
			},
		}

		service := NewOrderService(repo)
		err := service.CreateOrder(Order{Total: 10})
		assert.NoError(t, err)

	})
	t.Run("Total-less-than-0", func(t *testing.T) {
		repo := &mockOrderRepo{
			SaveFunc: func(order Order) error {
				return nil
			},
		}

		service := NewOrderService(repo)
		err := service.CreateOrder(Order{Total: -10})
		assert.Error(t, err)
		assert.Equal(t, "Total must be Positive", err.Error())

	})
	t.Run("repository errro", func(t *testing.T) {
		repo := &mockOrderRepo{
			SaveFunc: func(order Order) error {
				return errors.New("DataBase ERROR!!")
			},
		}

		service := NewOrderService(repo)
		err := service.CreateOrder(Order{Total: 100})
		assert.Error(t, err)
		assert.Equal(t, "DataBase ERROR!!", err.Error())

	})
}
