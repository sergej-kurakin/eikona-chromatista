package processor

import (
	"image/color"
)

// RXXProcessor keeps only the red channel.
type RXXProcessor struct {
	BaseProcessor
}

// NewRXXProcessor creates a new RXXProcessor instance.
func NewRXXProcessor() *RXXProcessor {
	return &RXXProcessor{
		BaseProcessor: BaseProcessor{name: "red"},
	}
}

// Process keeps only the red channel.
func (p *RXXProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	return color.RGBA{
		R: originalColor.R, G: 0, B: 0, A: originalColor.A,
	}, nil
}

// XGXProcessor keeps only the green channel.
type XGXProcessor struct {
	BaseProcessor
}

// NewXGXProcessor creates a new XGXProcessor instance.
func NewXGXProcessor() *XGXProcessor {
	return &XGXProcessor{
		BaseProcessor: BaseProcessor{name: "green"},
	}
}

// Process keeps only the green channel.
func (p *XGXProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	return color.RGBA{
		R: 0, G: originalColor.G, B: 0, A: originalColor.A,
	}, nil
}

// XXBProcessor keeps only the blue channel.
type XXBProcessor struct {
	BaseProcessor
}

// NewXXBProcessor creates a new XXBProcessor instance.
func NewXXBProcessor() *XXBProcessor {
	return &XXBProcessor{
		BaseProcessor: BaseProcessor{name: "blue"},
	}
}

// Process keeps only the blue channel.
func (p *XXBProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	return color.RGBA{
		R: 0, G: 0, B: originalColor.B, A: originalColor.A,
	}, nil
}

// PhotometricRedscaleProcessor converts to redscale using photometric luminosity.
type PhotometricRedscaleProcessor struct {
	BaseProcessor
	luminosityCalc LuminosityCalculator
}

// NewPhotometricRedscaleProcessor creates a processor with injectable luminosity calculator.
func NewPhotometricRedscaleProcessor(calc LuminosityCalculator) *PhotometricRedscaleProcessor {
	if calc == nil {
		calc = DefaultLuminosityCalculator
	}
	return &PhotometricRedscaleProcessor{
		BaseProcessor:  BaseProcessor{name: "redscale"},
		luminosityCalc: calc,
	}
}

// Process converts to redscale using photometric luminosity.
func (p *PhotometricRedscaleProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	luminocity := p.luminosityCalc.Calculate(originalColor)
	return color.RGBA{R: luminocity, G: 0, B: 0, A: originalColor.A}, nil
}

// PhotometricGreenscaleProcessor converts to greenscale using photometric luminosity.
type PhotometricGreenscaleProcessor struct {
	BaseProcessor
	luminosityCalc LuminosityCalculator
}

// NewPhotometricGreenscaleProcessor creates a processor with injectable luminosity calculator.
func NewPhotometricGreenscaleProcessor(calc LuminosityCalculator) *PhotometricGreenscaleProcessor {
	if calc == nil {
		calc = DefaultLuminosityCalculator
	}
	return &PhotometricGreenscaleProcessor{
		BaseProcessor:  BaseProcessor{name: "greenscale"},
		luminosityCalc: calc,
	}
}

// Process converts to greenscale using photometric luminosity.
func (p *PhotometricGreenscaleProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	luminocity := p.luminosityCalc.Calculate(originalColor)
	return color.RGBA{R: 0, G: luminocity, B: 0, A: originalColor.A}, nil
}

// PhotometricBluescaleProcessor converts to bluescale using photometric luminosity.
type PhotometricBluescaleProcessor struct {
	BaseProcessor
	luminosityCalc LuminosityCalculator
}

// NewPhotometricBluescaleProcessor creates a processor with injectable luminosity calculator.
func NewPhotometricBluescaleProcessor(calc LuminosityCalculator) *PhotometricBluescaleProcessor {
	if calc == nil {
		calc = DefaultLuminosityCalculator
	}
	return &PhotometricBluescaleProcessor{
		BaseProcessor:  BaseProcessor{name: "bluescale"},
		luminosityCalc: calc,
	}
}

// Process converts to bluescale using photometric luminosity.
func (p *PhotometricBluescaleProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	luminocity := p.luminosityCalc.Calculate(originalColor)
	return color.RGBA{R: 0, G: 0, B: luminocity, A: originalColor.A}, nil
}
