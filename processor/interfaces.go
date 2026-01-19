package processor

import "image/color"

// Processor defines the interface for all color processors.
// Implementations transform a single pixel color.
type Processor interface {
	// Process transforms a color and returns the result.
	Process(pixel color.Color) (color.Color, error)

	// Name returns the processor's identifier used for output file suffixes.
	Name() string
}
