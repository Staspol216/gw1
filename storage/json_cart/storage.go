package json_cart

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Staspol216/gw1/models"
	"github.com/samber/lo"
)

type Storage struct {
	carts []*models.Cart
	path string
}

func New(path string) (*Storage, error) {
	b, err := os.ReadFile(path)
	
	if err != nil {
		return nil, fmt.Errorf("os.ReadFile: %w", err)
	}
	
	carts := make([]*models.Cart, 0)

	err = json.Unmarshal(b, &carts)
	
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}
	
	return &Storage{
		carts: carts,
		path: path,
	}, nil
}

func (s *Storage) AddToCart(userID, productID, count int64) {
	defer func() {
		err := s.saveCartsToFile()
		if err != nil {
			fmt.Println(err)
		}
	}()

	cart, exists := lo.Find(s.carts, func(cart *models.Cart) bool {
		return cart.UserID == userID
	})
	
	if !exists {
		newCart := &models.Cart{
			UserID: userID,
			Products: []*models.Product{
				{
					ID: productID,
					Count: count,
				},
			},
		}
		s.carts = append(s.carts, newCart)
		return
	}
	
	product, exists := lo.Find(cart.Products, func(product *models.Product) bool {
		return product.ID == productID
	})
	
	if !exists {
		newProduct := &models.Product{
			ID: productID,
			Count: count,
		}
		
		cart.Products = append(cart.Products, newProduct)
		return
	}
	
	product.Count += count
}

func (s *Storage) saveCartsToFile() error {
	f, err := os.OpenFile(s.path, os.O_RDWR|os.O_TRUNC, 0666)
	
	if err != nil {
		return err
	}
	
	defer f.Close()
	
	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "\t")
	
	encoderError := encoder.Encode(s.carts)
	if encoderError == nil {
		return encoderError
	}
	
	fmt.Println("Struct successfully written to json")
	
	return nil
}

func (s *Storage) GetCartByUserID(userID int64) models.ICart {
	cart, _ := lo.Find(s.carts, func(cart *models.Cart) bool {
		return cart.UserID == userID
	})
	
	return cart
}
func (s *Storage) DeleteAllCarts()  {
	fmt.Println("dsfsd")
}