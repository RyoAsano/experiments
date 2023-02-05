package util

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/path"
)

func OutputToCsv(pth path.Path, dirname string, filename string) error {
	var colSize int = 1 + pth.Dim() // t, X_1(t), ..., X_n(t)

	records := make([][]string, pth.Grid().Size()*colSize)
	for k := 0; k <= pth.Grid().Size(); k++ {
		record := make([]string, colSize)

		t := pth.Grid().Get(k) // t
		x, err := pth.At(k)    // X(t)
		if err != nil {
			return err
		}
		// Get the k-th row.
		record[0] = fmt.Sprintf("%.20f", t)
		for dim := 1; dim <= x.Dim(); dim++ {
			val, err := x.Pr(dim)
			if err != nil {
				return err
			}
			record[dim] = fmt.Sprintf("%.20f", val)
		}
		// Save the k-th row
		records[k] = record
	}

	fullDir := fmt.Sprintf("db/%s", dirname)
	os.MkdirAll(fullDir, 0777)
	savingFile, err := os.Create(fmt.Sprintf("%s/%s", fullDir, filename))
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	w := csv.NewWriter(savingFile)
	return w.WriteAll(records)
}
