package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/hashicorp/hcl"
	h "github.com/mitchellh/mapstructure"
	"github.com/olekukonko/tablewriter"
)

func MakeTable(td [][]string) {

	table := tablewriter.NewWriter(os.Stdout)

	if len(td[0]) == 2 {

		table.SetHeader([]string{"TF Block Type", "Name"})
		table.SetFooter([]string{"", ""})

		table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.BgCyanColor, tablewriter.FgWhiteColor}, tablewriter.Colors{tablewriter.Bold, tablewriter.BgCyanColor, tablewriter.FgWhiteColor})
		table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlueColor}, tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiGreenColor})
		table.SetCaption(true, fmt.Sprintf("TF Block: %v", td[0][0]))

	} else if len(td[0]) == 3 {

		table.SetHeader([]string{"TF Block Type", "Resource Type", "Name"})
		table.SetFooter([]string{"", "", ""})

		table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.BgCyanColor, tablewriter.FgWhiteColor}, tablewriter.Colors{tablewriter.Bold, tablewriter.BgCyanColor, tablewriter.FgWhiteColor}, tablewriter.Colors{tablewriter.Bold, tablewriter.BgCyanColor, tablewriter.FgWhiteColor})
		table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlueColor}, tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiGreenColor}, tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor})
		table.SetCaption(true, fmt.Sprintf("TF Block: %v", td[0][0]))

	} else {
		fmt.Println("Error: unknown table data..")
	}

	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.SetCenterSeparator("|")

	table.AppendBulk(td)
	fmt.Println("\n")
	table.Render()
	fmt.Println("\n")

}

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

type test struct {
	Locals   []map[string]interface{}
	Output   []map[string]interface{}
	Resource []map[string]interface{}
	Data     []map[string]interface{}
	Module   []map[string]interface{}
	Provider []map[string]interface{}
	Variable []map[string]interface{}
}

var rest [][]string

func (t *test) showLocal() {
	// fmt.Println(t.Locals)
	for _, v := range t.Locals {
		for k1, _ := range v {
			fmt.Println(k1)
		}
	}
}

func (t *test) showVar() {
	// fmt.Println(t.Locals)
	for _, v := range t.Variable {
		for k1, _ := range v {
			fmt.Println(k1)
		}
	}
}

func (t *test) showRes(s [][]string) {
	for _, v := range t.Resource {
		for k1, v1 := range v {
			// fmt.Println(k1)
			for _, v2 := range v1.([]map[string]interface{}) {
				for _, v3 := range reflect.ValueOf(v2).MapKeys() {
					// fmt.Println(v3)
					s = append(s, []string{k1, string(v3.FieldByName()})
				}
			}
		}
	}
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

	// c := make(map[string][]string)
	// for k, v := range m {
	// 	switch k {
	// 	case "resource", "data":
	// 		for _, v1 := range v.([]map[string]interface{}) {
	// 			for k2, v2 := range v1 {
	// 				for _, v3 := range v2.([]map[string]interface{}) {
	// 					for k4 := range v3 {
	// 						c[k] = append(c[k], fmt.Sprintf("%v:%v", k2, k4))
	// 					}
	// 				}
	// 			}
	// 		}

	// 	default:
	// 		for _, v1 := range v.([]map[string]interface{}) {
	// 			for k2 := range v1 {
	// 				c[k] = append(c[k], k2)
	// 			}
	// 		}
	// 	}

	// }

	// k := os.Args[1]

	// switch k {
	// case "all":
	// 	for k := range c {
	// 		MakeTable(GetTableData(c, k))
	// 	}

	// default:
	// 	MakeTable(GetTableData(c, k))
	// }

	// fmt.Println(reflect.TypeOf(m))
	// fmt.Println(m["locals"])

	var res test

	config := &h.DecoderConfig{
		Result: &res,
	}

	decoder, err := h.NewDecoder(config)
	if err != nil {
		panic(err)
	}

	if err := decoder.Decode(m); err != nil {
		panic(err)
	}

	// err = h.WeakDecodeMetadata(m, &res, meta)
	// if err != nil {
	// 	panic(err)
	// }

	// b := os.Args[1]
	// res.showLocal()
	// res.showVar()
	res.showRes()

	// fmt.Println(meta)
	// fmt.Printf("Unused keys: %#v", meta.Keys[])

}
