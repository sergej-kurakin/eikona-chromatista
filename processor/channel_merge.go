package processor

import (
	"fmt"
	"image"
	"image/color"
)

// ChannelMerger combines independent R, G, and B channel images into one RGB image.
type ChannelMerger struct{}

// NewChannelMerger creates a new ChannelMerger instance.
func NewChannelMerger() *ChannelMerger {
	return &ChannelMerger{}
}

// Merge combines three same-sized images into a single RGB output.
func (m *ChannelMerger) Merge(redImg, greenImg, blueImg image.Image) (image.Image, error) {
	_ = m

	if redImg == nil || greenImg == nil || blueImg == nil {
		return nil, fmt.Errorf("all channel images must be provided")
	}

	redBounds := redImg.Bounds()
	if !redBounds.Eq(greenImg.Bounds()) || !redBounds.Eq(blueImg.Bounds()) {
		return nil, fmt.Errorf("channel images must have identical dimensions")
	}

	merged := image.NewRGBA(redBounds)

	for y := redBounds.Min.Y; y < redBounds.Max.Y; y++ {
		for x := redBounds.Min.X; x < redBounds.Max.X; x++ {
			r, _, _, _ := redImg.At(x, y).RGBA()
			_, g, _, _ := greenImg.At(x, y).RGBA()
			_, _, b, _ := blueImg.At(x, y).RGBA()

			merged.Set(x, y, color.RGBA{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: 255,
			})
		}
	}

	return merged, nil
}