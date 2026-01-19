package processor

import (
	"fmt"
	"image/color"
)

func RXXColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("RXXColorProcessor: failed to convert color to RGBA")
	}

	c := color.RGBA{
		R: originalColor.R, G: 0, B: 0, A: originalColor.A,
	}
	return c, nil
}

func XGXColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("XGXColorProcessor: failed to convert color to RGBA")
	}

	c := color.RGBA{
		R: 0, G: originalColor.G, B: 0, A: originalColor.A,
	}
	return c, nil
}

func XXBColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("XXBColorProcessor: failed to convert color to RGBA")
	}

	c := color.RGBA{
		R: 0, G: 0, B: originalColor.B, A: originalColor.A,
	}
	return c, nil
}

func PhotometricRedscaleColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("PhotometricRedscaleColorProcessor: failed to convert color to RGBA")
	}

	luminocity := calculateLuminocity(originalColor)
	c := color.RGBA{
		R: luminocity, G: 0, B: 0, A: originalColor.A,
	}

	return c, nil
}

func PhotometricGreenscaleColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("PhotometricGreenscaleColorProcessor: failed to convert color to RGBA")
	}

	luminocity := calculateLuminocity(originalColor)
	c := color.RGBA{
		R: 0, G: luminocity, B: 0, A: originalColor.A,
	}

	return c, nil
}

func PhotometricBluescaleColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("PhotometricBluescaleColorProcessor: failed to convert color to RGBA")
	}

	luminocity := calculateLuminocity(originalColor)
	c := color.RGBA{
		R: 0, G: 0, B: luminocity, A: originalColor.A,
	}

	return c, nil
}
