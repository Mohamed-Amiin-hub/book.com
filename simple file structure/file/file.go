package file

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type File struct {
	Name string `json:"filename"`
	Size int    `json:"size"`
}

func (f File) GetSize() int {
	return f.Size

}

func (f File) Display() {
	fmt.Printf("File name is %v and Size is %v bytes\n", f.Name, f.Size)
}
func (f File) Save() error {
	amiin := strings.ReplaceAll("f.name", "_", " ")
	amiin = strings.ToLower(amiin) + ".json"

	JSON, err := json.Marshal(f)
	if err != nil {
		return err
	}
	return os.WriteFile(amiin, JSON, 0644)
}

func NewSalaam(Name string, Size int) (File, error) {
	if Name == "" || Size == 0 {
		return File{}, errors.New("invalid input: file name or size is zero")
	}
	return File{
		Name: Name,
		Size: Size,
	}, nil
}
