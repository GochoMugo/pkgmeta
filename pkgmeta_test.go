package pkgmeta

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func ExampleLoad() {
	var pkg Metadata
	err := Load("package.json", &pkg)
	if err != nil {
		log.Fatal("could not load package metadata")
	}
	fmt.Println(pkg.Name)
	// Output: pkgmeta
}

func TestSuccessfulLoad(t *testing.T) {
	var pkg Metadata
	err := Load("package.json", &pkg)

	if err != nil {
		t.Error("failed to Load: ", err)
	}

	if pkg.Name != "pkgmeta" {
		t.Error("pkg.name mismatch")
	}
}

func TestLoadMissingFile(t *testing.T) {
	var pkg Metadata
	err := Load("missing.json", &pkg)

	if err == nil {
		t.Error("error not returned with missing file")
	}
}

func TestLoadInvalidJSON(t *testing.T) {
	filename := "_tmp.json"
	ioutil.WriteFile(filename, []byte(`{ name: "missing-quotes" }`), 0644)
	var pkg Metadata
	err := Load(filename, &pkg)

	if err == nil {
		t.Error("error not returned with invalid json")
	}
}

func TestLoadUsingCustomStruct(t *testing.T) {
	type customStruct struct {
		Name, Homepage string
	}
	var result customStruct
	err := Load("package.json", &result)

	if err != nil {
		t.Error("unexpected error")
	}

	if result.Name != "pkgmeta" {
		t.Error("failed to use custom struct")
	}
}
