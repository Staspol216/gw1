package map_cart

import (
	"github.com/Staspol216/gw1/models"
)

type Storage struct {
	carts map[int64]models.CartMap
}

func New() *Storage {
	return &Storage {
		carts: make(map[int64]models.CartMap),
	}
}

func (s *Storage) AddToCart(userID, productID, count int64) {
	if _, ok := s.carts[userID]; !ok {
		s.carts[userID] = make(models.CartMap)
	}
	s.carts[userID][productID] += count
}

func (s *Storage) GetCartByUserID(userID int64) models.ICart {
	return s.carts[userID]
}