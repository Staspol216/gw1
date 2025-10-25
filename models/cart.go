package models

import (
	"github.com/samber/lo"
)

type Product struct {
	ID    int64 `json:"id"`
	Count int64 `json:"count"`
}

type Cart struct {
	UserID   int64      `json:"user_id"`
	Products []*Product `json:"products"`
}

func (c *Cart) GetCountByProductID(productId int64) int64 {
	if c == nil {
		return 0
	}
	
	 product, exists := lo.Find(c.Products, func(product *Product) bool {
		return product.ID == productId
	 })
	 
	 if !exists {
		return 0
	 }
	 
	 return product.Count
}
