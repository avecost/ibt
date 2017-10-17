package main

import (
	"github.com/avecost/ibt/config"
	"github.com/avecost/ibt/csv"
	"os"
)

func main() {
	fToUpload := os.Args[1]
	if _, err := os.Stat(fToUpload); os.IsNotExist(err) {
		panic("File does not exist")
	}

	dbSourceName := os.Args[2]
	config.InitDB(dbSourceName)

	csv.ImportCSV(fToUpload)
}
