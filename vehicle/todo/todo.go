package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Details string `json:"details"`
}

type Todos struct {
	Details string `json:"details"`
}

type Todoss struct {
	Details string `json:"details"`
}

func (todo Todo) Display() {
	fmt.Println("todo.details:")
}
func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}
func (todos Todos) Display() {
	fmt.Println("todos.details:")

}
func (Todos Todos) Save() error {
	fileName := "todos.json"

	json, err := json.Marshal(Todos)

	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}
func (Todoss Todoss) Display() {
	fmt.Println("Todoss.Details")

}

func (Todoss Todoss) Save() error {
	fileName := "Todoss.json"
	json, err := json.Marshal(Todoss)

	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}

func NewCar(model string) (Todo, error) {
	if model == "" {
		return Todo{}, errors.New("invalid input")
	}
	return Todo{

		Details: model,
	}, nil
}

func NewMotorcycle(model string) (Todos, error) {
	if model == "" {
		return Todos{}, errors.New("invalid input")
	}
	return Todos{
		Details: model,
	}, nil
}

func NewTruct(model string) (Todoss, error) {
	if model == "" {
		return Todoss{}, errors.New("invalid input")
	}
	return Todoss{

		Details: model,
	}, nil
}
