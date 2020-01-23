package helpers

import (
	"os"
	"fmt"
	"github.com/olekukonko/tablewriter"
)

// MakeTable renders the data into a table
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