package crossplane

import (
	"context"
	"os"
	"path"
	"runtime"
	"testing"
)

func TestParseFile(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	exampleFile := path.Join(path.Dir(filename), "test/nginx.conf")

	response, err := ParseFile(context.Background(), exampleFile)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(response)
}

func TestParse(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	exampleFile := path.Join(path.Dir(filename), "test/nginx.conf")

	b, err := os.ReadFile(exampleFile)

	if err != nil {
		t.Fatal(err)
	}

	response, err := Parse(context.Background(), string(b))

	if err != nil {
		t.Fatal(err)
	}

	t.Log(response)
}
