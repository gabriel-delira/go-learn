package ex2

import (
	"log"
	"os"
	text "text/template"
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

type hotels struct {
	Hotels []hotel
}

// Test - Executes test
func Test(tpl *text.Template) {

	californiaHotels := hotels{
		Hotels: []hotel{
			hotel{"Hilton San Diego Bayfront", "1 Park Blvd", "San Diego", "92101", "Southern"},
			hotel{"Freeport Wine Country Inn", "8201 Freeport Blvd", "Sacramento", "95832", "Central"},
			hotel{"Auberge du Soleil", "180 Rutherford Hill Rd", "Rutherford", "94573", "Northern"},
		},
	}

	err := tpl.Execute(os.Stdout, californiaHotels)
	if err != nil {
		log.Println("ERROR:", err.Error())
		return
	}

}
