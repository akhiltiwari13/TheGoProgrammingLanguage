// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

// Celsius => exported type should have a comment to avoid the warning.
type Celsius float64

// Fahrenheit => commment for the exported type to avoid the green squiggly for the warning.
type Fahrenheit float64

// Kelvin => assignment from TGPL
type Kelvin float64

// exported constants for marking specific temperatures & also to avoid the warnings.
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	AbsoluteZeroK Kelvin  = 0
	FreezingK     Kelvin  = 273.15
	BoilingK      Kelvin  = 373.15
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g K", k) }
