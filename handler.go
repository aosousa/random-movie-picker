package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/aosousa/go-movielookup/models"
)

// Pick a random movie from my Excel watchlist
func pickRandomMovie() {
	var (
		movieTitle string
		movies     [27][]string
	)

	XLSXLocation := `Movies.xlsx`

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
	movieTitle = movies[column][row]

	printMovieInformation(movieTitle)
}

// Print movie's OMDB information if it is available
func printMovieInformation(title string) {
	var (
		queryURL string
		apiError models.Error
	)

	// build OMDB query URL
	title = strings.Replace(title, " ", "+", -1)
	queryURL = fmt.Sprintf("%st=%s&type=movie", baseURL, title)

	// call OMDB API
	res, err := http.Get(queryURL)
	if res.StatusCode != 200 || err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	handleError(err)

	json.Unmarshal(content, &apiError)
	if apiError.Response == "True" {
		var movie models.Movie
		json.Unmarshal(content, &movie)
		movie.PrintMovie()
	} else {
		fmt.Printf("Movie: %s\n", title)
	}
}
