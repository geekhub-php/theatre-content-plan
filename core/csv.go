package core

import (
	"encoding/csv"
	"log"
	"os"
)
var w = csv.NewWriter(os.Stdout)

func WriteCsv(record []string) {
	err := w.Write(record)
	if nil != err {
		log.Fatalln("error writing record to csv:", err)
	}
}

func FlushCsv() {
	w.Flush()
}
