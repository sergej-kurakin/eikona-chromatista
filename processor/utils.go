package processor

import "image/color"

// LuminosityCalculator provides luminosity calculation for processors that need it.
type LuminosityCalculator interface {
	Calculate(c color.RGBA) uint8
}

// PhotometricLuminosityCalculator implements the standard photometric formula.
type PhotometricLuminosityCalculator struct{}

// Calculate returns the luminosity using the photometric formula:
// Y = 0.2126 R + 0.7152 G + 0.0722 B.
func (c *PhotometricLuminosityCalculator) Calculate(originalColor color.RGBA) uint8 {
	return uint8(0.2126*float32(originalColor.R) + 0.7152*float32(originalColor.G) + 0.0722*float32(originalColor.B))
}

// DefaultLuminosityCalculator is the standard calculator instance.
var DefaultLuminosityCalculator LuminosityCalculator = &PhotometricLuminosityCalculator{}
