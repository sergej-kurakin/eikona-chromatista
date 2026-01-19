package processor

import (
	"fmt"
	"image/color"
	"math/rand"
)

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
