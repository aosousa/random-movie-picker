package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	var (
		movie  string
		movies [27][]string
	)

	XLSXLocation := `C:\Users\André\OneDrive\random\Movies.xlsx`

	// pick random number to identify column (0 = A, 26 = AA)
	column := random(0, 26)

	// open XLSX document
	xlsx, err := excelize.OpenFile(XLSXLocation)
	handleError(err)

	// get all rows in worksheet
	rows := xlsx.GetRows("Movies")

	// build movies slice
	for _, row := range rows {
		for i := 0; i < 27; i++ {
			if row[i] != "" {
				movies[i] = append(movies[i], row[i])
			}
		}
	}

	// pick random row
	row := random(0, len(movies[column]))

	// get movie from random column and row
	movie = movies[column][row]
	fmt.Printf("Movie: %s\n", movie)
}
