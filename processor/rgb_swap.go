package processor

import (
	"fmt"
	"image/color"
)

func GBRColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("GBRColorProcessor: failed to convert color to RGBA")
	}

	c := color.RGBA{
		R: originalColor.G, G: originalColor.B, B: originalColor.R, A: originalColor.A,
	}
	return c, nil
}

func GRBColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("GRBColorProcessor: failed to convert color to RGBA")
	}

	c := color.RGBA{
		R: originalColor.G, G: originalColor.R, B: originalColor.B, A: originalColor.A,
	}
	return c, nil
}

func BRGColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("BRGColorProcessor: failed to convert color to RGBA")
	}

	c := color.RGBA{
		R: originalColor.B, G: originalColor.R, B: originalColor.G, A: originalColor.A,
	}
	return c, nil
}

func BGRColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("BGRColorProcessor: failed to convert color to RGBA")
	}

	c := color.RGBA{
		R: originalColor.B, G: originalColor.G, B: originalColor.R, A: originalColor.A,
	}
	return c, nil
}

func RBGColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("RBGColorProcessor: failed to convert color to RGBA")
	}

	c := color.RGBA{
		R: originalColor.R, G: originalColor.B, B: originalColor.G, A: originalColor.A,
	}
	return c, nil
}
