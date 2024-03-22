package main

import (
	"fmt"

	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
)

func main() {
	lev := metrics.NewLevenshtein()
	lev.CaseSensitive = false
	lev.InsertCost = 1
	lev.ReplaceCost = 2
	lev.DeleteCost = 1

	similarity := strutil.Similarity("make", "Cake", lev)
	fmt.Printf("%.2f\n", similarity) // Output: 0.50
}
