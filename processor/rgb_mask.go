package processor

import (
	"fmt"
	"image/color"
)

func XGBColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("XGBColorProcessor: failed to convert color to RGBA")
	}

	c := color.RGBA{
		R: 0, G: originalColor.G, B: originalColor.B, A: originalColor.A,
	}
	return c, nil
}

func RXBColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("RXBColorProcessor: failed to convert color to RGBA")
	}

	c := color.RGBA{
		R: originalColor.R, G: 0, B: originalColor.B, A: originalColor.A,
	}
	return c, nil
}

func RGXColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("RGXColorProcessor: failed to convert color to RGBA")
	}

	c := color.RGBA{
		R: originalColor.R, G: originalColor.G, B: 0, A: originalColor.A,
	}
	return c, nil
}

func RBXColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("RBXColorProcessor: failed to convert color to RGBA")
	}

	c := color.RGBA{
		R: originalColor.R, G: originalColor.B, B: 0, A: originalColor.A,
	}
	return c, nil
}

func GRXColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("GRXColorProcessor: failed to convert color to RGBA")
	}

	c := color.RGBA{
		R: originalColor.G, G: originalColor.R, B: 0, A: originalColor.A,
	}
	return c, nil
}

func GBXColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("GBXColorProcessor: failed to convert color to RGBA")
	}

	c := color.RGBA{
		R: originalColor.G, G: originalColor.B, B: 0, A: originalColor.A,
	}
	return c, nil
}

func BRXColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("BRXColorProcessor: failed to convert color to RGBA")
	}

	c := color.RGBA{
		R: originalColor.B, G: originalColor.R, B: 0, A: originalColor.A,
	}
	return c, nil
}

func BGXColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("BGXColorProcessor: failed to convert color to RGBA")
	}

	c := color.RGBA{
		R: originalColor.B, G: originalColor.G, B: 0, A: originalColor.A,
	}
	return c, nil
}

func RXGColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("RXGColorProcessor: failed to convert color to RGBA")
	}

	c := color.RGBA{
		R: originalColor.R, G: 0, B: originalColor.G, A: originalColor.A,
	}
	return c, nil
}

func GXRColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("GXRColorProcessor: failed to convert color to RGBA")
	}

	c := color.RGBA{
		R: originalColor.G, G: 0, B: originalColor.R, A: originalColor.A,
	}
	return c, nil
}

func GXBColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("GXBColorProcessor: failed to convert color to RGBA")
	}

	c := color.RGBA{
		R: originalColor.G, G: 0, B: originalColor.B, A: originalColor.A,
	}
	return c, nil
}

func BXRColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("BXRColorProcessor: failed to convert color to RGBA")
	}

	c := color.RGBA{
		R: originalColor.B, G: 0, B: originalColor.R, A: originalColor.A,
	}
	return c, nil
}

func BXGColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("BXGColorProcessor: failed to convert color to RGBA")
	}

	c := color.RGBA{
		R: originalColor.B, G: 0, B: originalColor.G, A: originalColor.A,
	}
	return c, nil
}

func XRGColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("XRGColorProcessor: failed to convert color to RGBA")
	}

	c := color.RGBA{
		R: 0, G: originalColor.R, B: originalColor.G, A: originalColor.A,
	}
	return c, nil
}

func XRBColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("XRBColorProcessor: failed to convert color to RGBA")
	}

	c := color.RGBA{
		R: 0, G: originalColor.R, B: originalColor.B, A: originalColor.A,
	}
	return c, nil
}

func XGRColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("XGRColorProcessor: failed to convert color to RGBA")
	}

	c := color.RGBA{
		R: 0, G: originalColor.G, B: originalColor.R, A: originalColor.A,
	}
	return c, nil
}

func XBRColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("XBRColorProcessor: failed to convert color to RGBA")
	}

	c := color.RGBA{
		R: 0, G: originalColor.B, B: originalColor.R, A: originalColor.A,
	}
	return c, nil
}

func XBGColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("XBGColorProcessor: failed to convert color to RGBA")
	}

	c := color.RGBA{
		R: 0, G: originalColor.B, B: originalColor.G, A: originalColor.A,
	}
	return c, nil
}
