package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/hashicorp/hcl"
)

// HandleError logs the error onto the console
func HandleError(e error, msg string) {
	if e != nil {
		log.Fatal(e, msg)
	}
}

// ReadSource reads the source path for terraform files
func ReadSource(path string) map[string]interface{} {
	var v interface{}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		HandleError(err, "Error: Unable to read the source path")
	}
	for _, files := range files {
		if files.IsDir() != true && filepath.Ext(files.Name()) == ".tf" {
			input, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", path, files.Name()))
			if err != nil {
				HandleError(err, fmt.Sprintf("Error: Unable to read from file %v", files.Name()))
			}

			err = hcl.Unmarshal(input, &v)
			if err != nil {
				HandleError(err, "Error: Unable to parse HCL files.")
			}

		}
	}

	jsonData, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		HandleError(err, "Error: Unable to marshal json.")
	}

	m := make(map[string]interface{})
	err = json.Unmarshal(jsonData, &m)
	if err != nil {
		HandleError(err, "Error: cannot unmarshal json data.")
	}
	return m
}
