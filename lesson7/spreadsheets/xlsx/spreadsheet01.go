package main

import (
	"fmt"

	"github.com/tealeg/xlsx/v3"
)

// jméno generovaného souboru
const spreadsheetName = "test01.xlsx"

func main() {
	// konstrukce struktury typu File
	worksheet := xlsx.NewFile()

	fmt.Println(worksheet)

	// pokus o uložení sešitu
	err := worksheet.Save(spreadsheetName)
	if err != nil {
		panic(err)
	}
}
