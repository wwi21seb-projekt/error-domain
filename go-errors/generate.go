// Package main generate.go
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type CustomError struct {
	Message    string
	Code       string
	HttpStatus int
}

func main() {
	// Path to the JSON file containing the errors
	path := filepath.Join("..", "errors", "errors.json")
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var errors []CustomError
	err = json.Unmarshal(file, &errors)
	if err != nil {
		panic(err)
	}

	// Generate errors.go in the goerrors directory
	outPath := filepath.Join("./", "goerrors", "errors.go")
	if err := os.MkdirAll(filepath.Dir(outPath), os.ModePerm); err != nil {
		panic(err)
	}

	out, err := os.Create(outPath)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	_, _ = out.WriteString("// CODE GENERATED - DO NOT EDIT\n")
	_, _ = out.WriteString("package goerrors\n\n")

	_, _ = out.WriteString("type CustomError struct {\n\tMessage string\n\tCode string\n\tHttpStatus int\n}\n\n")

	_, _ = out.WriteString("var (\n")
	for _, e := range errors {
		_, _ = out.WriteString(fmt.Sprintf("\t%s = &CustomError{\n\t\tMessage: \"%s\",\n\t\tCode: \"%s\",\n\t\tHttpStatus: %d,\n\t}\n", e.Code, e.Message, e.Code, e.HttpStatus))
	}
	_, _ = out.WriteString(")\n")

	// Run go fmt on the generated file

}
