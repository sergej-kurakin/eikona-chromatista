package processor

import (
	"fmt"
	"image/color"
)

func GrayColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("GrayColorProcessor: failed to convert color to RGBA")
	}

	// Offset colors a little, adjust it to your taste
	r := float64(originalColor.R) * 0.92126
	g := float64(originalColor.G) * 0.97152
	b := float64(originalColor.B) * 0.90722
	// average
	grey := uint8((r + g + b) / 3)
	c := color.RGBA{
		R: grey, G: grey, B: grey, A: originalColor.A,
	}
	return c, nil
}

func PhotometricGrayscaleColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("PhotometricGrayscaleColorProcessor: failed to convert color to RGBA")
	}

	luminocity := calculateLuminocity(originalColor)
	c := color.RGBA{
		R: luminocity, G: luminocity, B: luminocity, A: originalColor.A,
	}

	return c, nil
}

func PhotometricGraychromeColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("PhotometricGraychromeColorProcessor: failed to convert color to RGBA")
	}

	luminocity := calculateLuminocity(originalColor)

	var c color.RGBA

	// 0 < x <= 64 -- 0
	// 64 < x <= 128 -- 85
	// 128 < x <= 192 -- 180
	// 192 < x <= 255 -- 255

	switch {
	case luminocity <= 64:
		c = color.RGBA{
			R: 0, G: 0, B: 0, A: originalColor.A,
		}
	case luminocity <= 128:
		c = color.RGBA{
			R: 85, G: 85, B: 85, A: originalColor.A,
		}
	case luminocity <= 192:
		c = color.RGBA{
			R: 180, G: 180, B: 180, A: originalColor.A,
		}
	default:
		c = color.RGBA{
			R: 255, G: 255, B: 255, A: originalColor.A,
		}
	}

	return c, nil
}

func PhotometricGraychromeNegativeColorProcessor(pixel color.Color) (color.Color, error) {
	graychromePixel, err := PhotometricGraychromeColorProcessor(pixel)
	if err != nil {
		return nil, err
	}

	return NegativeColorProcessor(graychromePixel)
}
