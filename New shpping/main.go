package main

import (
	"fmt"

	"example.com/new-shopping/alaab"
	"example.com/new-shopping/item"
)

type UpdateAmount interface {
	Save() error
}
type output interface {
	Display()
	UpdateAmount
}

func main() {
	{

		fmt.Println("New Shopping Item")

		// Create a new shopping item from alaab package
		alaabta, err := alaab.NewShopping("perfumes", 999.4, 110)
		if err != nil {
			fmt.Println("Error creating shopping item:", err)
			return
		}
		outputData(alaabta)

		// Add quantity to alaabta
		err = alaabta.AddQuantity(20)
		if err != nil {
			fmt.Println("Error adding quantity:", err)
		} else {
			fmt.Println("Item added successfully!")
			outputData(alaabta)
		}

		// Remove quantity from alaabta
		err = alaabta.RemoveQuantity(5)
		if err != nil {
			fmt.Println("Error removing quantity:", err)
		} else {
			fmt.Println("Item removed successfully!")

		}
		outputData(alaabta)

		// Create a new shopping item from item package
		amiinItem, err := item.NewShopping(10)
		if err != nil {
			fmt.Println("Error creating item:", err)
			return
		}
		outputData(amiinItem)

	}

}

func outputData(data output) {
	data.Display()
	iskudarAlaab(data)
}

func iskudarAlaab(data UpdateAmount) {
	err := data.Save()
	if err != nil {
		fmt.Println("Error removing quantity:")
	} else {
		fmt.Println("Item removed successfully!")
		return
	}

}
