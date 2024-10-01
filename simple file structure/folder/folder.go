package folder

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type AddFile struct {
	Filedetails string `json:"filedetails"`
}

func (A AddFile) Display() {
	fmt.Println("f.filedetails")
}
func (A AddFile) Save() error {
	amiin := "A.json"

	JSON, err := json.Marshal(A)
	if err != nil {
		return err
	}
	return os.WriteFile(amiin, JSON, 0644)
}

func NewSalaam(name string) (AddFile, error) {
	if name == "" {
		return AddFile{}, errors.New("invalid input: file name or size is zero")
	}
	return AddFile{

		Filedetails: name,
	}, nil
}
