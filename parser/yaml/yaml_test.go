package yaml

import (
	"io"
	"log"
	"os"
	"testing"
)

func TestYamlExtensions(t *testing.T) {
	yp := &Parser{}
	exts := yp.Extensions()
	// Make sure we have at least the expected extensions registered
	expect := []string{"yaml", "yml"}
	for _, actual := range exts {
		for i, expected := range expect {
			if expected == actual {
				expect[i] = ""
			}
		}
	}
	for _, expected := range expect {
		if expected != "" {
			t.Fatalf("Missing extension for YAML pareser. Missing: %s", expected)
		}
	}
}

func TestParseYamlStdin(t *testing.T) {
	f, err := os.Open("testdata/rt.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	tmp := os.Stdin
	defer func() {
		os.Stdin = tmp
	}()
	os.Stdin = r
	go func() {
		io.Copy(w, f)
		w.Close()
	}()
	yp := &Parser{}
	_, err = yp.Parse("-")
	if err != nil {
		t.Fatal(err)
	}
}

func TestParseYamlFile(t *testing.T) {
	yp := &Parser{}
	scene, err := yp.Parse("testdata/rt.yaml")
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%#v", scene)
}
