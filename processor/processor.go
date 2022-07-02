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

func GBColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: 0, G: originalColor.G, B: originalColor.B, A: originalColor.A,
	}
	return c
}

func RBColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: originalColor.R, G: 0, B: originalColor.B, A: originalColor.A,
	}
	return c
}

func RGColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: originalColor.R, G: originalColor.G, B: 0, A: originalColor.A,
	}
	return c
}

func RColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: originalColor.R, G: 0, B: 0, A: originalColor.A,
	}
	return c
}

func GColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: 0, G: originalColor.G, B: 0, A: originalColor.A,
	}
	return c
}

func BColorProcessor(pixel color.Color) color.Color {
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
	processors[6] = GBColorProcessor
	processors[7] = RBColorProcessor
	processors[8] = RGColorProcessor
	processors[9] = RColorProcessor
	processors[10] = GColorProcessor
	processors[11] = BColorProcessor

	processor_num := rand.Intn(len(processors))

	return processors[processor_num](pixel)
}

func PhotometricGrayscaleColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	// Y = 0.2126 R + 0.7152 G + 0.0722 B
	Y := 0.2126*float32(originalColor.R) + 0.7152*float32(originalColor.G) + 0.0722*float32(originalColor.B)
	c := color.RGBA{
		R: uint8(Y), G: uint8(Y), B: uint8(Y), A: originalColor.A,
	}

	return c
}

func PhotometricRedscaleColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	// Y = 0.2126 R + 0.7152 G + 0.0722 B
	Y := 0.2126*float32(originalColor.R) + 0.7152*float32(originalColor.G) + 0.0722*float32(originalColor.B)
	c := color.RGBA{
		R: uint8(Y), G: 0, B: 0, A: originalColor.A,
	}

	return c
}

func PhotometricGreenscaleColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	// Y = 0.2126 R + 0.7152 G + 0.0722 B
	Y := 0.2126*float32(originalColor.R) + 0.7152*float32(originalColor.G) + 0.0722*float32(originalColor.B)
	c := color.RGBA{
		R: 0, G: uint8(Y), B: 0, A: originalColor.A,
	}

	return c
}

func PhotometricBluescaleColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	// Y = 0.2126 R + 0.7152 G + 0.0722 B
	Y := 0.2126*float32(originalColor.R) + 0.7152*float32(originalColor.G) + 0.0722*float32(originalColor.B)
	c := color.RGBA{
		R: 0, G: 0, B: uint8(Y), A: originalColor.A,
	}

	return c
}

func PhotometricGraychromeColorProcessor(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	// Y = 0.2126 R + 0.7152 G + 0.0722 B
	Y := 0.2126*float32(originalColor.R) + 0.7152*float32(originalColor.G) + 0.0722*float32(originalColor.B)

	k := uint8(Y)
	var c color.RGBA

	// 0 < x <= 64 -- 0
	// 64 < x <= 128 -- 85
	// 128 < x <= 192 -- 180
	// 192 < x <= 255 -- 255

	if k <= 64 {
		c = color.RGBA{
			R: 0, G: 0, B: 0, A: originalColor.A,
		}
	} else if k <= 128 {
		c = color.RGBA{
			R: 85, G: 85, B: 85, A: originalColor.A,
		}
	} else if k <= 192 {
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
