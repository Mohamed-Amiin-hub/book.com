package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"example.com/vehicle/todo"
	"example.com/vehicle/vehicles"
)

type saver interface {
	Save() error
}

type output interface {
	Display()
	saver
}

func main() {
	fmt.Println("Display basic information about different types of vehicles.")

	vehicleType := getVehicles("Please enter a vehicle type (car, motorcycle, truck):")
	fmt.Println("You entered:", vehicleType)

	switch vehicleType {
	case "car":

		fmt.Println("Enter car details:")
		make := getVehicles("Make:")
		model := getVehicles("Model:")
		year := getInteger("Year:")

		todoshowDetails := getVehicles("Show Details of cars:")
		todo, err := todo.NewCar(todoshowDetails)

		if err != nil {
			fmt.Println(err)
			return
		}

		car, err := vehicles.NewCar(make, model, year)
		if err != nil {
			fmt.Println("Error creating car:")
			return
		}

		err = outputtableDetails(todo)

		if err != nil {
			fmt.Println(err)
			return
		}

		outputtableDetails(car)

	case "motorcycle":

		fmt.Println("Enter motorcycle details:")
		make := getVehicles("Make:")
		model := getVehicles("Model:")
		year := getInteger("Year:")

		showDetails := getVehicles("Show Details of motorcycle:")
		todos, err := todo.NewMotorcycle(showDetails)
		if err != nil {
			fmt.Println(err)
			return
		}
		motorcycle, err := vehicles.NewMotorcycle(make, model, year)
		if err != nil {
			fmt.Println("Error creating motorcycle:")
			return
		}
		err = outputtableDetails(todos)

		if err != nil {
			fmt.Println(err)
			return
		}

		outputtableDetails(motorcycle)

		fmt.Println("Saving motorcycle succeeded!")

	case "truck":

		fmt.Println("Enter truck details:")
		make := getVehicles("Make:")
		model := getVehicles("Model:")
		year := getInteger("Year:")
		showDetails := getVehicles("Show Details of truck:")
		todoss, err := todo.NewTruct(showDetails)

		if err != nil {
			fmt.Println(err)
			return
		}

		truck, err := vehicles.NewTruct(make, model, year)
		if err != nil {
			fmt.Println("Error creating truck:")
			return
		}

		err = outputtableDetails(todoss)

		if err != nil {
			fmt.Println(err)
			return
		}

		outputtableDetails(truck)

	default:
		fmt.Println("Unknown vehicle type.")

	}
}

func outputtableDetails(details output) error {
	details.Display()
	return showData(details)
}

func showData(details saver) error {
	err := details.Save()

	if err != nil {
		fmt.Println("Error saving truck:")
		return err
	}

	fmt.Println("Saving truck succeeded!")
	return nil
}

func getVehicles(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text

}

func getInteger(prompt string) int {
	fmt.Printf("%v ", prompt)
	var value string
	fmt.Scanln(&value)
	intValue, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return getInteger(prompt)
	}
	return intValue
}
