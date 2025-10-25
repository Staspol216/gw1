package main

import (
	"flag"
	"fmt"

	"github.com/Staspol216/gw1/models"
	"github.com/Staspol216/gw1/storage/json_cart"
)

const (
	countOfArgs = 3
)

type storage interface {
	AddToCart(userID, productID, count int64)
	GetCartByUserID(userID int64) models.ICart
}


func main() {
	
	var (
		userID 	  = flag.Int64("user", 0, "user id")
		productID = flag.Int64("product", 0, "product id")
		count 	  = flag.Int64("count", 0, "count")
	)
	
	flag.Parse()
	
	if nFlag := flag.NFlag(); nFlag != countOfArgs {
		fmt.Printf("oh no. have %d, need %d", nFlag, countOfArgs)
		return
	}
		
	jsonCarts, err := json_cart.New("storage/json_cart/carts.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	var cartsStorage storage = jsonCarts
	cartsStorage.AddToCart(*userID, *productID, *count)
	
	fmt.Println(cartsStorage.GetCartByUserID(*userID).GetCountByProductID(*productID))
}
