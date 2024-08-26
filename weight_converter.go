package main

var WeightUnits = []string{"milligram", "gram", "kilogram", "ounce", "pound"}

type WeightConverter struct {
	input  float64
	from   string
	to     string
	result float64

	milligram float64
	gram      float64
	kilogram  float64
	ounce     float64
	pound     float64
}

func (wc *WeightConverter) initWeightConverter() {
	switch wc.from {
	case "milligram":
		wc.milligram = wc.input
		wc.from_Milligram()
	case "gram":
		wc.gram = wc.input
	case "kilogram":
		wc.kilogram = wc.input
		wc.from_Kilogram()
	case "ounce":
		wc.ounce = wc.input
		wc.from_Ounce()
	case "pound":
		wc.pound = wc.input
		wc.from_Pound()
	}

	switch wc.to {
	case "milligram":
		wc.to_Milligram()
		wc.result = wc.milligram
	case "gram":
		wc.result = wc.gram
	case "kilogram":
		wc.to_Kilogram()
		wc.result = wc.kilogram
	case "ounce":
		wc.to_Ounce()
		wc.result = wc.ounce
	case "pound":
		wc.to_Pound()
		wc.result = wc.pound
	}
}

func (wc *WeightConverter) to_Milligram() { wc.milligram = wc.gram * 1000 }
func (wc *WeightConverter) to_Kilogram()  { wc.kilogram = wc.gram / 1000 }
func (wc *WeightConverter) to_Ounce()     { wc.ounce = wc.gram / 28.35 }
func (wc *WeightConverter) to_Pound()     { wc.pound = wc.gram / 453.6 }

func (wc *WeightConverter) from_Milligram() { wc.gram = wc.milligram / 1000 }
func (wc *WeightConverter) from_Kilogram()  { wc.gram = wc.kilogram * 1000 }
func (wc *WeightConverter) from_Ounce()     { wc.gram = wc.ounce * 28.35 }
func (wc *WeightConverter) from_Pound()     { wc.gram = wc.pound * 453.6 }
