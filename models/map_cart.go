package models

import (
	"fmt"
	"strings"
)

type ICart interface {
	GetCountByProductID(productID int64) int64
}

type CartMap map[int64]int64

func (c CartMap) String() string {
	builder := strings.Builder{}
	
	for productId, count := range c {
		builder.WriteString(fmt.Sprintf("product_id %d count %d\n", productId, count))
	}
	
	return builder.String()
}

func (c CartMap) GetCountByProductID(productId int64) int64 {
	return c[productId]
}
