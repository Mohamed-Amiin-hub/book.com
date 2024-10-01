package alaab

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Item struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func (item Item) Display() {
	fmt.Printf("Item: %v | Price: $%.2f | Quantity: %d\n", item.Name, item.Price, item.Quantity)
}

func (item Item) AddQuantity(amount int) error {
	if amount <= 0 {
		return errors.New("invalid amount. Must be greater than zero")
	}
	item.Quantity += amount
	fmt.Printf("Added %d to %s. New Quantity: %d\n", amount, item.Name, item.Quantity)
	return nil
}

func (item Item) RemoveQuantity(amount int) error {
	if amount <= 0 {
		return errors.New("invalid amount. Must be greater than zero")
	}
	if amount > item.Quantity {
		return errors.New("not enough stock to remove")
	}
	item.Quantity -= amount
	fmt.Printf("Removed %v from %s. New Quantity: %d\n", amount, item.Name, item.Quantity)
	return nil
}

func (item Item) Save() error {
	fileName := strings.ReplaceAll(item.Name, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(item)

	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}

func NewShopping(name string, price float64, quantity int) (Item, error) {
	return Item{
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}, nil
}
