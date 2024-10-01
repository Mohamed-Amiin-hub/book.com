package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"example.com/shopping/amiin"
)

func main() {
	name, price, quantity := shoppingoutput()

	amiinNote, err := amiin.New(name, price, quantity)

	if err != nil {
		fmt.Println(err)
		return
	}

	amiinNote.Display()
	amiinNote.UpdateQuantity(10) // Assuming you want to update quantity by 1
	err = amiinNote.Save()

	if err != nil {
		fmt.Println("invalid number of quantity")
		return
	}

	fmt.Println("Quantity updated successfully!")

}

func shoppingoutput() (string, float64, int) {
	name := getSHoppingInput("Enter the name: ")
	price := getFloat("Enter the price: ")
	quantity := getInteger("Enter the quantity: ")

	return name, price, quantity
}

func getSHoppingInput(promt string) string {
	fmt.Printf("%v ", promt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}

func getInteger(promt string) int {
	fmt.Printf("%v ", promt)
	var value string
	fmt.Scan(&value)
	intval, err := strconv.Atoi(value)

	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return getInteger(promt)

	}
	return intval
}

func getFloat(promt string) float64 {
	fmt.Printf("%v ", promt)
	var value string
	fmt.Scan(&value)
	floatval, err := strconv.ParseFloat(value, 64)

	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return getFloat(promt)

	}
	return floatval
}
