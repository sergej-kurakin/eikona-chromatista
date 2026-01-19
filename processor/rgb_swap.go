package processor

import (
	"image/color"
)

// GBRProcessor swaps RGB channels to GBR order.
type GBRProcessor struct {
	BaseProcessor
}

// NewGBRProcessor creates a new GBRProcessor instance.
func NewGBRProcessor() *GBRProcessor {
	return &GBRProcessor{
		BaseProcessor: BaseProcessor{name: "gbr"},
	}
}

// Process swaps RGB channels to GBR.
func (p *GBRProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	return color.RGBA{
		R: originalColor.G, G: originalColor.B, B: originalColor.R, A: originalColor.A,
	}, nil
}

// GRBProcessor swaps RGB channels to GRB order.
type GRBProcessor struct {
	BaseProcessor
}

// NewGRBProcessor creates a new GRBProcessor instance.
func NewGRBProcessor() *GRBProcessor {
	return &GRBProcessor{
		BaseProcessor: BaseProcessor{name: "grb"},
	}
}

// Process swaps RGB channels to GRB.
func (p *GRBProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	return color.RGBA{
		R: originalColor.G, G: originalColor.R, B: originalColor.B, A: originalColor.A,
	}, nil
}

// BRGProcessor swaps RGB channels to BRG order.
type BRGProcessor struct {
	BaseProcessor
}

// NewBRGProcessor creates a new BRGProcessor instance.
func NewBRGProcessor() *BRGProcessor {
	return &BRGProcessor{
		BaseProcessor: BaseProcessor{name: "brg"},
	}
}

// Process swaps RGB channels to BRG.
func (p *BRGProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	return color.RGBA{
		R: originalColor.B, G: originalColor.R, B: originalColor.G, A: originalColor.A,
	}, nil
}

// BGRProcessor swaps RGB channels to BGR order.
type BGRProcessor struct {
	BaseProcessor
}

// NewBGRProcessor creates a new BGRProcessor instance.
func NewBGRProcessor() *BGRProcessor {
	return &BGRProcessor{
		BaseProcessor: BaseProcessor{name: "bgr"},
	}
}

// Process swaps RGB channels to BGR.
func (p *BGRProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	return color.RGBA{
		R: originalColor.B, G: originalColor.G, B: originalColor.R, A: originalColor.A,
	}, nil
}

// RBGProcessor swaps RGB channels to RBG order.
type RBGProcessor struct {
	BaseProcessor
}

// NewRBGProcessor creates a new RBGProcessor instance.
func NewRBGProcessor() *RBGProcessor {
	return &RBGProcessor{
		BaseProcessor: BaseProcessor{name: "rbg"},
	}
}

// Process swaps RGB channels to RBG.
func (p *RBGProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	return color.RGBA{
		R: originalColor.R, G: originalColor.B, B: originalColor.G, A: originalColor.A,
	}, nil
}
