package helpers

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

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
	var v map[string]interface{}
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
	return v
}

// GetTableData obtains the required data form the data structure to populate the table.
func GetTableData(m map[string][]string, resKey string) [][]string {
	var tableData [][]string

	if _, ok := m[resKey]; ok {
		for _, v := range m[resKey] {
			l := strings.Split(v, ":")
			tableData = append(tableData, append([]string{resKey}, l...))
		}
	} else {
		fmt.Println("Error: requested resource type doesnt exist")
	}

	return tableData
}

// InitFiles initalises the set of empty terraform configuration files in a provided directory
func InitFiles(path string ) {
	
}