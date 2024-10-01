package book

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Book struct {
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	ISBN      int       `json:"iSBN"`
	CreatedAt time.Time `json:"created_at"`
}

func (book Book) Display() {
	fmt.Printf("Details of the book titled '%v' by author %v with ISBN %v:\n", book.Title, book.Author, book.ISBN)

}

func (book Book) Save() error {
	fileName := strings.ReplaceAll(book.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(book)

	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}

func New(title, author string, iSBN int) (Book, error) {
	if title == "" || author == "" || iSBN == 0 {
		return Book{}, errors.New("invalid book output")
	}
	return Book{
		Title:     title,
		Author:    author,
		ISBN:      iSBN,
		CreatedAt: time.Now(),
	}, nil
}
