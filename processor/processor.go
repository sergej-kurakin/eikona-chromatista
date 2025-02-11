package processor

import (
	"image/color"
	"math/rand"
)

type ProcessorFunc func(color.Color) color.Color

func GrayColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	// Offset colors a little, adjust it to your taste
	r := float64(originalColor.R) * 0.92126
	g := float64(originalColor.G) * 0.97152
	b := float64(originalColor.B) * 0.90722
	// average
	grey := uint8((r + g + b) / 3)
	c := color.RGBA{
		R: grey, G: grey, B: grey, A: originalColor.A,
	}
	return c
}

func GBRColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: originalColor.G, G: originalColor.B, B: originalColor.R, A: originalColor.A,
	}
	return c
}

func GRBColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: originalColor.G, G: originalColor.R, B: originalColor.B, A: originalColor.A,
	}
	return c
}

func BRGColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: originalColor.B, G: originalColor.R, B: originalColor.G, A: originalColor.A,
	}
	return c
}

func BGRColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: originalColor.B, G: originalColor.G, B: originalColor.R, A: originalColor.A,
	}
	return c
}

func RBGColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: originalColor.R, G: originalColor.B, B: originalColor.G, A: originalColor.A,
	}
	return c
}

func XGBColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: 0, G: originalColor.G, B: originalColor.B, A: originalColor.A,
	}
	return c
}

func RXBColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: originalColor.R, G: 0, B: originalColor.B, A: originalColor.A,
	}
	return c
}

func RGXColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: originalColor.R, G: originalColor.G, B: 0, A: originalColor.A,
	}
	return c
}

func RBXColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: originalColor.R, G: originalColor.B, B: 0, A: originalColor.A,
	}
	return c
}

func GRXColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: originalColor.G, G: originalColor.R, B: 0, A: originalColor.A,
	}
	return c
}

func GBXColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: originalColor.G, G: originalColor.B, B: 0, A: originalColor.A,
	}
	return c
}

func BRXColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: originalColor.B, G: originalColor.R, B: 0, A: originalColor.A,
	}
	return c
}

func BGXColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: originalColor.B, G: originalColor.G, B: 0, A: originalColor.A,
	}
	return c
}

func RXGColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: originalColor.R, G: 0, B: originalColor.G, A: originalColor.A,
	}
	return c
}

func GXRColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: originalColor.G, G: 0, B: originalColor.R, A: originalColor.A,
	}
	return c
}

func GXBColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: originalColor.G, G: 0, B: originalColor.B, A: originalColor.A,
	}
	return c
}

func BXRColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: originalColor.B, G: 0, B: originalColor.R, A: originalColor.A,
	}
	return c
}

func BXGColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: originalColor.B, G: 0, B: originalColor.G, A: originalColor.A,
	}
	return c
}

func XRGColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: 0, G: originalColor.R, B: originalColor.G, A: originalColor.A,
	}
	return c
}

func XRBColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: 0, G: originalColor.R, B: originalColor.B, A: originalColor.A,
	}
	return c
}

func XGRColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: 0, G: originalColor.G, B: originalColor.R, A: originalColor.A,
	}
	return c
}

func XBRColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: 0, G: originalColor.B, B: originalColor.R, A: originalColor.A,
	}
	return c
}

func XBGColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: 0, G: originalColor.B, B: originalColor.G, A: originalColor.A,
	}
	return c
}

func RXXColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: originalColor.R, G: 0, B: 0, A: originalColor.A,
	}
	return c
}

func XGXColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: 0, G: originalColor.G, B: 0, A: originalColor.A,
	}
	return c
}

func XXBColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: 0, G: 0, B: originalColor.B, A: originalColor.A,
	}
	return c
}

func RandomColorProcessor(pixel color.Color) color.Color {

	var processors [12]func(color.Color) color.Color
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

	processor_num := rand.Intn(len(processors))

	return processors[processor_num](pixel)
}

func PhotometricGrayscaleColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	luminocity := calculateLuminocity(originalColor)
	c := color.RGBA{
		R: luminocity, G: luminocity, B: luminocity, A: originalColor.A,
	}

	return c
}

func PhotometricRedscaleColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	luminocity := calculateLuminocity(originalColor)
	c := color.RGBA{
		R: luminocity, G: 0, B: 0, A: originalColor.A,
	}

	return c
}

func PhotometricGreenscaleColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	luminocity := calculateLuminocity(originalColor)
	c := color.RGBA{
		R: 0, G: luminocity, B: 0, A: originalColor.A,
	}

	return c
}

func PhotometricBluescaleColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	luminocity := calculateLuminocity(originalColor)
	c := color.RGBA{
		R: 0, G: 0, B: luminocity, A: originalColor.A,
	}

	return c
}

func PhotometricGraychromeColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	luminocity := calculateLuminocity(originalColor)

	var c color.RGBA

	// 0 < x <= 64 -- 0
	// 64 < x <= 128 -- 85
	// 128 < x <= 192 -- 180
	// 192 < x <= 255 -- 255

	if luminocity <= 64 {
		c = color.RGBA{
			R: 0, G: 0, B: 0, A: originalColor.A,
		}
	} else if luminocity <= 128 {
		c = color.RGBA{
			R: 85, G: 85, B: 85, A: originalColor.A,
		}
	} else if luminocity <= 192 {
		c = color.RGBA{
			R: 180, G: 180, B: 180, A: originalColor.A,
		}
	} else {
		c = color.RGBA{
			R: 255, G: 255, B: 255, A: originalColor.A,
		}
	}

	return c
}

func NegativeColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: 255 - originalColor.R, G: 255 - originalColor.G, B: 255 - originalColor.B, A: originalColor.A,
	}
	return c
}

func PhotometricGraychromeNegativeColorProcessor(pixel color.Color) color.Color {
	graychrome_pixel := PhotometricGraychromeColorProcessor(pixel)
	return NegativeColorProcessor(graychrome_pixel)
}

func SepiaColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)

	tr := 0.393*float32(originalColor.R) + 0.769*float32(originalColor.G) + 0.189*float32(originalColor.B)
	tg := 0.349*float32(originalColor.R) + 0.686*float32(originalColor.G) + 0.168*float32(originalColor.B)
	tb := 0.272*float32(originalColor.R) + 0.534*float32(originalColor.G) + 0.131*float32(originalColor.B)

	var r uint8 = 0
	var g uint8 = 0
	var b uint8 = 0

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

	return c
}

func calculateLuminocity(originalColor color.RGBA) uint8 {
	// Y = 0.2126 R + 0.7152 G + 0.0722 B
	return uint8(0.2126*float32(originalColor.R) + 0.7152*float32(originalColor.G) + 0.0722*float32(originalColor.B))
}

func InfraredProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: originalColor.B,
		G: originalColor.G,
		B: originalColor.R,
		A: originalColor.A,
	}
	return c
}
