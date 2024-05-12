// generate.go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type CustomError struct {
	Message    string `json:"message"`
	Code       string `json:"code"`
	HttpStatus int    `json:"httpStatus"`
}

func main() {
	jsonFile, err := os.Open("errors/errors.json")
	if err != nil {
		panic(err)
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			panic(err)
		}
	}(jsonFile)

	bytes, _ := io.ReadAll(jsonFile)
	var errors []CustomError
	if err := json.Unmarshal(bytes, &errors); err != nil {
		panic(err)

	}

	out, err := os.Create("go-errors/errors.go")
	if err != nil {
		panic(err)
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			panic(err)
		}
	}(out)

	_, _ = out.WriteString("package goerrors\n\n")

	_, _ = out.WriteString("type CustomError struct {\n\tMessage string\n\tCode string\n\tHttpStatus int\n}\n\n")
	_, _ = out.WriteString("var (\n")
	for _, e := range errors {
		_, _ = out.WriteString(fmt.Sprintf("\t%s = &CustomError{\n\t\tMessage: \"%s\",\n\t\tCode: \"%s\",\n\t\tHttpStatus: %d,\n\t}\n", e.Code, e.Message, e.Code, e.HttpStatus))
	}
	_, _ = out.WriteString(")\n")
}
