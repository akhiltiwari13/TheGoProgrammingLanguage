package units

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToK converts Celsius to Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c) - 273.15 }

// FToK converts Fahrenheit to Kelvin
func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }

// FtToMts converts Feet to Meters
func FtToMts(ft Feet) Meters { return Meters(ft) * .3048 }

// MtToFt converts Meters to Feet
func MtToFt(m Meters) Feet { return Feet(m) * 3.28084 }

// PdToKg converts Pounds to Kilogram
func PdToKg(p Pounds) Kilograms { return Kilograms(p) * .453592 }

// KgToPd converts Kilograms to Pounds
func KgToPd(k Kilograms) Pounds { return Pounds(k) * 2.20462 }
