package amiin

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// var newQuantity int = 100

// Fields represents a record in the inventory table.
type Fields struct {
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
}

func (F Fields) Display() {
	fmt.Printf("the name if item is %v and price of item %v and quantity of item %v\n\n", F.Name, F.Price, F.Quantity)
}

func (F Fields) Save() error {
	magaca := strings.ReplaceAll("F.Name", " ", "_")
	magaca = strings.ToLower(magaca) + ".json"

	json, err := json.Marshal(F)

	if err != nil {
		fmt.Println("Error data  :", err)
		return err
	}
	return os.WriteFile(magaca, json, 0644)

}

func (F Fields) UpdateQuantity(newQuantity int) error {
	for {
		fmt.Println("1. add quantity")
		fmt.Println("1. add quantity")
		fmt.Println("3. exit")

		var item int
		fmt.Print("Enter your item: ")
		fmt.Scan(&item)

		if item == 1 {
			fmt.Print("amount of add item: ")
			var amountItem int
			fmt.Scan(&amountItem)
		}

	}
	// 	fmt.Print("newQuantity: ")
	// 	var newQuantity float64
	// 	fmt.Scan(&newQuantity)
	// 	if newQuantity > 0 {
	// 		return errors.New("error")

	// 	} else if newQuantity < 0 {
	// 		return errors.New("error")
	// 	}
	// 	F.quantity -= newQuantity
	// 	return nil
}

func New(name string, price float64, quantity int) (Fields, error) {
	if name == "" || price == 0 || quantity == 0 {
		return Fields{}, errors.New("error")
	}
	return Fields{
		Name:      name,
		Price:     price,
		Quantity:  quantity,
		CreatedAt: time.Now(),
	}, nil
}
