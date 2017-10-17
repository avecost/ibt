package csv

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/avecost/ibt/ties"
	"strconv"
)

func processCSV(rc io.Reader) (ch chan []string) {
	ch = make(chan []string, 10)
	go func() {
		r := csv.NewReader(rc)
		if _, err := r.Read(); err != nil { // read header
			log.Fatal(err)
		}
		defer close(ch)

		for {
			rec, err := r.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}
			ch <- rec
		}
	}()
	return
}

func ImportCSV(fCSV string) (int, error) {
	f, err := os.Open(fCSV)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	i := 0
	for row := range processCSV(f) {
		i++
		bb, _ := strconv.ParseFloat(row[2], 64)
		bp, _ := strconv.ParseFloat(row[3], 64)
		bt, _ := strconv.ParseFloat(row[4], 64)
		tp, _ := strconv.ParseFloat(row[5], 64)
		ties.InsertTie(ties.Tie{0, row[0], row[1], bb, bp, bt, tp, row[8], row[9], row[10], row[11]})
	}
	return i, nil
}
