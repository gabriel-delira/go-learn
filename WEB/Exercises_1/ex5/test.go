package ex5

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	text "text/template"

	"../entities"
)

func parseCsvToTable() entities.Table {

	csvfile, errRead := os.Open("ex5/table.csv")
	if errRead != nil {
		log.Printf("Error reading file. \n Description: %v \n", errRead.Error())
		return entities.Table{}
	}
	r := csv.NewReader(csvfile)
	isHeader := true
	table := entities.Table{}
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if isHeader {
			table.Header.Date = record[0]
			table.Header.Open = record[1]
			table.Header.High = record[2]
			table.Header.Low = record[3]
			table.Header.Close = record[4]
			table.Header.Volume = record[5]
			table.Header.AdjClose = record[6]
		} else {
			date := record[0]
			open, _ := strconv.ParseFloat(record[1], 64)
			high, _ := strconv.ParseFloat(record[2], 64)
			low, _ := strconv.ParseFloat(record[3], 64)
			close, _ := strconv.ParseFloat(record[4], 64)
			volume, _ := strconv.ParseFloat(record[5], 64)
			adjClose, _ := strconv.ParseFloat(record[6], 64)
			newRow := entities.Row{
				Date:     date,
				Open:     open,
				High:     high,
				Low:      low,
				Close:    close,
				Volume:   volume,
				AdjClose: adjClose,
			}
			table.Rows = append(table.Rows, newRow)
		}

		isHeader = false
	}

	return table
}

// Test - Executes test
func Test(tpl *text.Template) {

	parsedTable := parseCsvToTable()

	err := tpl.Execute(os.Stdout, parsedTable)
	if err != nil {
		log.Println("ERROR:", err.Error())
		return
	}

}
