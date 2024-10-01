package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Info string `json:"info"`
}

func (todo Todo) Display() {
	fmt.Println("todo.info")

}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}

func New(Author string) (Todo, error) {
	if Author == "" {
		return Todo{}, errors.New("invalid book output")
	}
	return Todo{

		Info: Author,
	}, nil
}
