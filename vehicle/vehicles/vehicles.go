package vehicles

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Car struct {
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

type Motorcycle struct {
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

type Truck struct {
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

func (car Car) Display() {
	fmt.Printf("Car - Make: %v, Model: %v, Year: %v\n", car.Make, car.Model, car.Year)
}
func (car Car) Save() error {
	fileName := strings.ReplaceAll(car.Make, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(car)

	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}
func (motorcycle Motorcycle) Display() {
	fmt.Printf("Motorcycle - Make: %s, Model: %s, Year: %d\n", motorcycle.Make, motorcycle.Model, motorcycle.Year)

}
func (motorcycle Motorcycle) Save() error {
	fileName := strings.ReplaceAll(motorcycle.Make, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(motorcycle)

	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}
func (truck Truck) Display() {
	fmt.Printf("Truck - Make: %s, Model: %s, Year: %d\n", truck.Make, truck.Model, truck.Year)

}

func (truck Truck) Save() error {
	fileName := strings.ReplaceAll(truck.Make, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(truck)

	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}

func NewCar(make, model string, year int) (Car, error) {
	if make == "" || model == "" || year == 0 {
		return Car{}, errors.New("invalid input")
	}
	return Car{
		Make:  make,
		Model: model,
		Year:  year,
	}, nil
}

func NewMotorcycle(make, model string, year int) (Motorcycle, error) {
	if make == "" || model == "" || year == 0 {
		return Motorcycle{}, errors.New("invalid input")
	}
	return Motorcycle{
		Make:  make,
		Model: model,
		Year:  year,
	}, nil
}

func NewTruct(make, model string, year int) (Truck, error) {
	if make == "" || model == "" || year == 0 {
		return Truck{}, errors.New("invalid input")
	}
	return Truck{
		Make:  make,
		Model: model,
		Year:  year,
	}, nil
}
