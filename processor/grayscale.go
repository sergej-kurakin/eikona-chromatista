package processor

import (
	"image/color"
)

// GrayProcessor converts colors to grayscale using weighted averaging.
type GrayProcessor struct {
	BaseProcessor
}

// NewGrayProcessor creates a new GrayProcessor instance.
func NewGrayProcessor() *GrayProcessor {
	return &GrayProcessor{
		BaseProcessor: BaseProcessor{name: "gray"},
	}
}

// Process transforms the pixel to grayscale.
func (p *GrayProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	// Offset colors a little, adjust it to your taste
	r := float64(originalColor.R) * 0.92126
	g := float64(originalColor.G) * 0.97152
	b := float64(originalColor.B) * 0.90722
	// average
	grey := uint8((r + g + b) / 3)

	return color.RGBA{R: grey, G: grey, B: grey, A: originalColor.A}, nil
}

// PhotometricGrayscaleProcessor converts colors using photometric luminosity.
type PhotometricGrayscaleProcessor struct {
	BaseProcessor
	luminosityCalc LuminosityCalculator
}

// NewPhotometricGrayscaleProcessor creates a processor with injectable luminosity calculator.
func NewPhotometricGrayscaleProcessor(calc LuminosityCalculator) *PhotometricGrayscaleProcessor {
	if calc == nil {
		calc = DefaultLuminosityCalculator
	}
	return &PhotometricGrayscaleProcessor{
		BaseProcessor:  BaseProcessor{name: "photometric_grayscale"},
		luminosityCalc: calc,
	}
}

// Process transforms the pixel to photometric grayscale.
func (p *PhotometricGrayscaleProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	luminocity := p.luminosityCalc.Calculate(originalColor)
	return color.RGBA{R: luminocity, G: luminocity, B: luminocity, A: originalColor.A}, nil
}

// PhotometricGraychromeProcessor creates posterized grayscale effect.
type PhotometricGraychromeProcessor struct {
	BaseProcessor
	luminosityCalc LuminosityCalculator
}

// NewPhotometricGraychromeProcessor creates a graychrome processor.
func NewPhotometricGraychromeProcessor(calc LuminosityCalculator) *PhotometricGraychromeProcessor {
	if calc == nil {
		calc = DefaultLuminosityCalculator
	}
	return &PhotometricGraychromeProcessor{
		BaseProcessor:  BaseProcessor{name: "photometric_graychrome"},
		luminosityCalc: calc,
	}
}

// Process transforms the pixel to posterized grayscale.
func (p *PhotometricGraychromeProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	luminocity := p.luminosityCalc.Calculate(originalColor)

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

// PhotometricGraychromeNegativeProcessor chains graychrome with negative.
type PhotometricGraychromeNegativeProcessor struct {
	BaseProcessor
	graychrome Processor
	negative   Processor
}

// NewPhotometricGraychromeNegativeProcessor creates a processor with injectable dependencies.
func NewPhotometricGraychromeNegativeProcessor(graychrome, negative Processor) *PhotometricGraychromeNegativeProcessor {
	if graychrome == nil {
		graychrome = NewPhotometricGraychromeProcessor(nil)
	}
	if negative == nil {
		negative = NewNegativeProcessor()
	}
	return &PhotometricGraychromeNegativeProcessor{
		BaseProcessor: BaseProcessor{name: "photometric_graychrome_negative"},
		graychrome:    graychrome,
		negative:      negative,
	}
}

// Process applies graychrome then negative transformation.
func (p *PhotometricGraychromeNegativeProcessor) Process(pixel color.Color) (color.Color, error) {
	graychromePixel, err := p.graychrome.Process(pixel)
	if err != nil {
		return nil, err
	}

	return p.negative.Process(graychromePixel)
}
