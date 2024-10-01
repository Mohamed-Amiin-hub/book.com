package item

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Item struct {
	ItemNumber int `json:"itemNumber"`
}

func (item Item) Display() {
	fmt.Printf("item.info")
}

// func (item Item) AddQuantity(amount int) error {
// 	if amount <= 0 {
// 		return errors.New("invalid amount. Must be greater than zero")
// 	}
// 	item.Quantity += amount
// 	fmt.Printf("Added %d to %s. New Quantity: %d\n", amount, item.Name, item.Quantity)
// 	return nil
// }

// func (item Item) RemoveQuantity(amount int) error {

// 	if amount <= 0 {
// 		return errors.New("invalid amount. Must be greater than zero")
// 	}
// 	if amount > item.Quantity {
// 		return errors.New("not enough stock to remove")
// 	}
// 	item.Quantity -= amount
// 	fmt.Printf("Removed %v from %s. New Quantity: %d\n", amount, item.Name, item.Quantity)
// 	return nil
// }

func (item Item) Save() error {
	fileName := "item.json"

	json, err := json.Marshal(item)

	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}

func NewShopping(quantity int) (Item, error) {
	if quantity <= 0 {
		return Item{}, errors.New("invalid item")
	}
	return Item{
		ItemNumber: quantity,
	}, nil
}
