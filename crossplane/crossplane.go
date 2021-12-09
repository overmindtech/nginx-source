package crossplane

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
)

type Error struct {
	// the full path of the config file
	File string `json:"file,omitempty"`

	// integer line number the directive that caused the error
	Line int `json:"line,omitempty"`

	// the error message
	Error string `json:"error,omitempty"`
}

type Directive struct {
	// the name of the directive
	Directive string `json:"directive,omitempty"`

	// integer line number the directive started on
	Line int `json:"line,omitempty"`

	// Array of String arguments
	Args []string `json:"args,omitempty"`

	// Array of integers (included if this is an include directive)
	Inlcudes []int `json:"includes,omitempty"`

	// Array of Directive Objects (included iff this is a block)
	Block []Directive `json:"block,omitempty"`
}

type Config struct {
	// the full path of the config file
	File string `json:"file,omitempty"`

	// "ok" or "failed" if errors is not empty array
	Status string `json:"status,omitempty"`

	// Array of Error objects
	Errors []Error `json:"errors,omitempty"`

	// Array of Directive objects
	Parsed []Directive `json:"parsed,omitempty"`
}

type Response struct {
	// "ok" or "failed" if "errors" is not empty
	Status string `json:"status,omitempty"`

	// aggregation of "errors" from Config objects
	Errors []Error `json:"errors,omitempty"`

	// Array of Config objects
	Config []Config `json:"config,omitempty"`
}

func Parse(ctx context.Context, content string) (Response, error) {
	var tempFile *os.File
	var err error

	tempFile, err = ioutil.TempFile("", "crossplane")

	if err != nil {
		return Response{}, err
	}

	defer os.Remove(tempFile.Name())

	_, err = tempFile.WriteString(content)

	if err != nil {
		return Response{}, err
	}

	return ParseFile(ctx, tempFile.Name())
}

func ParseFile(ctx context.Context, filePath string) (Response, error) {
	var out []byte
	var err error
	var response Response

	if err != nil {
		return response, err
	}

	cmd := exec.CommandContext(ctx, "crossplane", "parse", filePath)

	out, err = cmd.Output()

	if err != nil {
		return response, err
	}

	err = json.Unmarshal(out, &response)

	return response, err
}
