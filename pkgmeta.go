/*
Package pkgmeta allows loading project metadata from a manifest file
such as package.json, manifest.json etc. It currently only supports JSON files.
*/
package pkgmeta

import (
	"encoding/json"
	"io/ioutil"
)

// Metadata is a struct of common key-value pairs in package manifest
// It is largely influenced by Node.js/npm package.json. See https://docs.npmjs.com/files/package.json
type Metadata struct {
	Name, Version string
}

// Load reads the package manifest file at `filepath` and unmarshals its content into the receiving variable
func Load(filepath string, result interface{}) error {
	metadata, err := ioutil.ReadFile(filepath)

	if err != nil {
		return err
	}

	err = json.Unmarshal(metadata, result)

	if err != nil {
		return err
	}

	return nil
}

// LoadDefaults reads the package manifest file using the default Metadata struct
func LoadDefaults(filepath string) (Metadata, error) {
	var metadata Metadata
	err := Load(filepath, &metadata)
	return metadata, err
}
