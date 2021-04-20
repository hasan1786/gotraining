package writers

import (
	"fmt"
	"os"

	"github.com/hasan1786/gotraining/ticker"
	"github.com/olekukonko/tablewriter"
)

func WriteToTable(data map[string]ticker.Ticker, timestamp string) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Sign", "Rating"})

	table.ClearRows()
	for currency, t := range data {
		//b, err := json.Marshal(t)
		table.Append([]string{currency, fmt.Sprintf("%f", t.Buy), fmt.Sprintf("%f", t.Sell), t.Symbol, timestamp})
		//append(data, append([]string(b), curret)
	}

	table.Render()

}
