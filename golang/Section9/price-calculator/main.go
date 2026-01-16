package main

import (
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	calculatePrices(taxRates...)

}

func calculatePrices(taxRates ...float64) {

	for _, taxRate := range taxRates {
		pricesJob := prices.NewTaxIncludedPriceJob(taxRate)
		pricesJob.Process()
	}

}
