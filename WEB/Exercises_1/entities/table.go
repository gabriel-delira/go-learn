package entities

type Header struct {
	Date, Open, High, Low, Close, Volume, AdjClose string
}

type Row struct {
	Date                                     string
	Open, High, Low, Close, Volume, AdjClose float64
}

type Table struct {
	Header Header
	Rows   []Row
}
