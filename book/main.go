package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"example.com/book/book"
	"example.com/book/todo"
)

type Saver interface {
	Save() error
}

type output interface {
	Display()
	Saver
}

func main() {
	fmt.Println("welcome Basic book system to store and retrieve information about book.")

	title, author, iSBN := getBookNote()

	todoInfo := getBookInput("todo Info:")
	todo, err := todo.New(todoInfo)

	if err != nil {
		fmt.Println(err)
		return
	}

	noteBook, err := book.New(title, author, iSBN)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(noteBook)

}

func outputData(data output) error {
	data.Display()
	return saveData(data)
}
func saveData(data Saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving book is failed. ")
		return err
	}
	fmt.Println("Saving book is succeeded! .")
	return nil
}

func getBookNote() (string, string, int) {

	title := getBookInput("title:")

	author := getBookInput("Author:")

	iSBN := getInteger("ISBN:")

	return title, author, iSBN
}

func getBookInput(promt string) string {
	fmt.Printf("%v ", promt)
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
