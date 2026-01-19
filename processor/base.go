package processor

import (
	"fmt"
	"image/color"
)

// BaseProcessor provides common functionality for all processors.
type BaseProcessor struct {
	name string
}

// Name returns the processor's identifier.
func (b *BaseProcessor) Name() string {
	return b.name
}

// ToRGBA converts a color.Color to color.RGBA with error handling.
func (b *BaseProcessor) ToRGBA(pixel color.Color) (color.RGBA, error) {
	rgba, ok := color.RGBAModel.Convert(pixel).(color.RGBA)
	if !ok {
		return color.RGBA{}, fmt.Errorf("%s: failed to convert color to RGBA", b.name)
	}
	return rgba, nil
}
