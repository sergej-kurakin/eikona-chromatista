package processor

import (
	"fmt"
	"image/color"
	"math/rand"
)

type ColorProcessor func(color.Color) (color.Color, error)

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

func RandomColorProcessor(pixel color.Color) (color.Color, error) {
	var processors [12]ColorProcessor
	processors[0] = GrayColorProcessor
	processors[1] = RBGColorProcessor
	processors[2] = GBRColorProcessor
	processors[3] = GRBColorProcessor
	processors[4] = BRGColorProcessor
	processors[5] = BGRColorProcessor
	processors[6] = XGBColorProcessor
	processors[7] = RXBColorProcessor
	processors[8] = RGXColorProcessor
	processors[9] = RXXColorProcessor
	processors[10] = XGXColorProcessor
	processors[11] = XXBColorProcessor

	processorKey := rand.Intn(len(processors))

	return processors[processorKey](pixel)
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

func NegativeColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("NegativeColorProcessor: failed to convert color to RGBA")
	}

	c := color.RGBA{
		R: 255 - originalColor.R, G: 255 - originalColor.G, B: 255 - originalColor.B, A: originalColor.A,
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

func SepiaColorProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("SepiaColorProcessor: failed to convert color to RGBA")
	}

	tr := 0.393*float32(originalColor.R) + 0.769*float32(originalColor.G) + 0.189*float32(originalColor.B)
	tg := 0.349*float32(originalColor.R) + 0.686*float32(originalColor.G) + 0.168*float32(originalColor.B)
	tb := 0.272*float32(originalColor.R) + 0.534*float32(originalColor.G) + 0.131*float32(originalColor.B)

	var r uint8
	var g uint8
	var b uint8

	if tr > 255 {
		r = 255
	} else {
		r = uint8(tr)
	}
	if tg > 255 {
		g = 255
	} else {
		g = uint8(tg)
	}
	if tb > 255 {
		b = 255
	} else {
		b = uint8(tb)
	}

	c := color.RGBA{
		R: r, G: g, B: b, A: originalColor.A,
	}

	return c, nil
}

func calculateLuminocity(originalColor color.RGBA) uint8 {
	// Y = 0.2126 R + 0.7152 G + 0.0722 B
	return uint8(0.2126*float32(originalColor.R) + 0.7152*float32(originalColor.G) + 0.0722*float32(originalColor.B))
}

func InfraredProcessor(pixel color.Color) (color.Color, error) {
	originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

	if !ok {
		return nil, fmt.Errorf("InfraredProcessor: failed to convert color to RGBA")
	}

	c := color.RGBA{
		R: originalColor.B,
		G: originalColor.G,
		B: originalColor.R,
		A: originalColor.A,
	}
	return c, nil
}
