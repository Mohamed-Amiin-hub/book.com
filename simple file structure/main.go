package main

import (
	"fmt"

	"file.com/file"
	"file.com/folder"
)

type Saver interface {
	Save() error
}
type Saverr interface {
	Display()
	Saver
}

func main() {
	// name, size := getOutPut()

	// amiin, err := file.NewSalaam(name, size)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// amiin.Display()
	// amiin.Save()
	folders, err := folder.NewSalaam("func file name")
	if err != nil {
		fmt.Println(err)
		return
	}
	amiin, err := file.NewSalaam("khalifa", 123)
	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(folders)

	outputData(amiin)
}
func outputData(data Saverr) error {
	data.Display()
	return mohamed(data)
}
func mohamed(data Saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Amiin saved failed.")
		return err
	}
	fmt.Println("Amiin saved successfully.")
	return nil
}

// func getOutPut() (string, int) {
// 	name := getLaqanyo("please enter name: ")
// 	size := getinteger("please enter size: ")

// 	return name, size
// }

// func getinteger(prompt string) int {
// 	fmt.Printf("%v ", prompt)
// 	var size string
// 	fmt.Scanln(&size)
// 	intSize, err := strconv.Atoi(size)
// 	if err != nil {
// 		fmt.Println("Invalid input. Please enter a number.")
// 		return getinteger(prompt)
// 	}
// 	return intSize

// }

// func getLaqanyo(promt string) string {
// 	var name string
// 	fmt.Printf("%v ", promt)
// 	fmt.Scanln(&name)
// 	return name

// }
