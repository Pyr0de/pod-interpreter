package main

import (
	"errors"
	"os"
)

type GenerateFunc func() (string, error);

func main() {
	generate_code(OPERATOR_FILE, OperatorGen)
}



func generate_code(filePath string, c GenerateFunc) error {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o644)
	defer file.Close()
	
	if err != nil {
		return errors.New("Could not find & create file " + filePath)
	}

	code, err := c()
	if err != nil {
		return err
	}

	_, err = file.WriteString(code)
	if err != nil {
		return errors.New("Failed to write to file " + filePath)
	}

	return nil
}
