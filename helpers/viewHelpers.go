package helpers

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

// RenderTable renders the data into a table
func RenderTable(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"TF Object Type", "Resource Type", "Name"})
	table.SetFooter([]string{"", "", ""})
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.SetCenterSeparator("|")

	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgRedColor})

	table.SetColumnColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlueColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiGreenColor})

	// table.SetFooterColor(tablewriter.Colors{}, tablewriter.Colors{},
	// 	tablewriter.Colors{tablewriter.Bold})

	table.AppendBulk(data)
	table.Render()

}
