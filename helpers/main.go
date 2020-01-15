package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/hashicorp/hcl"
)

func Keys(m map[string]interface{}) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func main() {

	var m map[string]interface{}
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("unable to read from path")
	}
	for _, files := range files {
		if files.IsDir() != true && filepath.Ext(files.Name()) == ".tf" {

			input, err := ioutil.ReadFile(files.Name())
			if err != nil {
				fmt.Println("unable to read from file %v: %s", files.Name(), err)
			}

			err = hcl.Unmarshal(input, &m)
			if err != nil {
				fmt.Println("unable to parse HCL: %s", err)
			}

		}
	}

	// data := [][]string{}
	c := make(map[string][]string)
	for k, v := range m {
		switch k {
		case "resource":
			for _, v1 := range v.([]map[string]interface{}) {
				for k2, v2 := range v1 {
					for _, v3 := range v2.([]map[string]interface{}) {
						for k4 := range v3 {
							c[k] = append(c[k], k2, k4)
						}
					}
				}
			}

		default:
			for _, v1 := range v.([]map[string]interface{}) {
				for k2 := range v1 {
					c[k] = append(c[k], k2)
				}
			}
		}

	}

	// for _, val := range data {
	// 	resType := val[0]
	// 	_, keyExists := c[resType]
	// 	if keyExists {
	// 		for _, dat := range val[1:] {
	// 			c[resType] = append(c[resType], dat)
	// 		}
	// 	} else {
	// 		for _, dat := range val[1:] {
	// 			c[resType] = append(c[resType], dat)
	// 		}
	// 	}

	// }
	fmt.Println(c)

	// table := tablewriter.NewWriter(os.Stdout)
	// table.SetHeader([]string{"TF Object Type", "Resource Type", "Name"})
	// table.SetFooter([]string{"", "", ""})
	// table.SetAutoMergeCells(true)
	// table.SetRowLine(true)
	// table.SetCenterSeparator("|")

	// table.SetHeaderColor(
	// 	tablewriter.Colors{tablewriter.Bold, tablewriter.BgRedColor},
	// 	tablewriter.Colors{tablewriter.Bold, tablewriter.BgRedColor},
	// 	tablewriter.Colors{tablewriter.Bold, tablewriter.BgRedColor})

	// table.SetColumnColor(
	// 	tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlueColor},
	// 	tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
	// 	tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiGreenColor})

	// table.SetFooterColor(tablewriter.Colors{}, tablewriter.Colors{},
	// 	tablewriter.Colors{tablewriter.Bold})

	// table.AppendBulk(data)
	// table.Render()

}
