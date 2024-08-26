package main

type LengthConverter struct {
	input  float64
	from   string
	to     string
	result float64

	millimeter float64
	centimeter float64
	meter      float64
	kilometer  float64
	inch       float64
	foot       float64
	yard       float64
	mile       float64
}

func (lc *LengthConverter) initLengthConverter() {
	switch lc.from {
	case "millimeter":
		lc.millimeter = lc.input
		lc.from_Millimeter()
	case "centimeter":
		lc.centimeter = lc.input
		lc.from_Centimeter()
	case "meter":
		lc.meter = lc.input
	case "kilometer":
		lc.kilometer = lc.input
		lc.from_Kilometer()
	case "inch":
		lc.inch = lc.input
		lc.from_Inch()
	case "foot":
		lc.foot = lc.input
		lc.from_Foot()
	case "yard":
		lc.yard = lc.input
		lc.from_Yard()
	case "mile":
		lc.mile = lc.input
		lc.from_Mile()
	}

	switch lc.to {
	case "millimeter":
		lc.to_Millimeter()
		lc.result = lc.millimeter
	case "centimeter":
		lc.to_Centimeter()
		lc.result = lc.centimeter
	case "meter":
		lc.result = lc.meter
	case "kilometer":
		lc.to_Kilometer()
		lc.result = lc.kilometer
	case "inch":
		lc.to_Inch()
		lc.result = lc.inch
	case "foot":
		lc.to_Foot()
		lc.result = lc.foot
	case "yard":
		lc.to_Yard()
		lc.result = lc.yard
	case "mile":
		lc.to_Mile()
		lc.result = lc.mile
	}
}

func (lc *LengthConverter) to_Millimeter() { lc.millimeter = lc.meter * 1000 }
func (lc *LengthConverter) to_Centimeter() { lc.centimeter = lc.meter * 100 }
func (lc *LengthConverter) to_Kilometer()  { lc.kilometer = lc.meter / 1000 }
func (lc *LengthConverter) to_Inch()       { lc.inch = lc.meter * 39.37 }
func (lc *LengthConverter) to_Foot()       { lc.foot = lc.meter * 3.281 }
func (lc *LengthConverter) to_Yard()       { lc.foot = lc.meter * 1.094 }
func (lc *LengthConverter) to_Mile()       { lc.mile = lc.meter / 1609 }

func (lc *LengthConverter) from_Millimeter() { lc.meter = lc.millimeter / 1000 }
func (lc *LengthConverter) from_Centimeter() { lc.meter = lc.centimeter / 100 }
func (lc *LengthConverter) from_Kilometer()  { lc.meter = lc.kilometer * 1000 }
func (lc *LengthConverter) from_Inch()       { lc.meter = lc.inch / 39.37 }
func (lc *LengthConverter) from_Foot()       { lc.meter = lc.foot / 3.281 }
func (lc *LengthConverter) from_Yard()       { lc.meter = lc.yard / 1.094 }
func (lc *LengthConverter) from_Mile()       { lc.meter = lc.mile * 1609 }
