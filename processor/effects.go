package processor

import (
	"image/color"
	"math/rand"
)

// NegativeProcessor inverts colors.
type NegativeProcessor struct {
	BaseProcessor
}

// NewNegativeProcessor creates a new NegativeProcessor instance.
func NewNegativeProcessor() *NegativeProcessor {
	return &NegativeProcessor{
		BaseProcessor: BaseProcessor{name: "negative"},
	}
}

// Process inverts the colors.
func (p *NegativeProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	return color.RGBA{
		R: 255 - originalColor.R,
		G: 255 - originalColor.G,
		B: 255 - originalColor.B,
		A: originalColor.A,
	}, nil
}

// SepiaProcessor applies sepia tone effect.
type SepiaProcessor struct {
	BaseProcessor
}

// NewSepiaProcessor creates a new SepiaProcessor instance.
func NewSepiaProcessor() *SepiaProcessor {
	return &SepiaProcessor{
		BaseProcessor: BaseProcessor{name: "sepia"},
	}
}

// Process applies sepia tone effect.
func (p *SepiaProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	tr := 0.393*float32(originalColor.R) + 0.769*float32(originalColor.G) + 0.189*float32(originalColor.B)
	tg := 0.349*float32(originalColor.R) + 0.686*float32(originalColor.G) + 0.168*float32(originalColor.B)
	tb := 0.272*float32(originalColor.R) + 0.534*float32(originalColor.G) + 0.131*float32(originalColor.B)

	var r, g, b uint8

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

	return color.RGBA{R: r, G: g, B: b, A: originalColor.A}, nil
}

// InfraredProcessor simulates infrared film effect by swapping R and B.
type InfraredProcessor struct {
	BaseProcessor
}

// NewInfraredProcessor creates a new InfraredProcessor instance.
func NewInfraredProcessor() *InfraredProcessor {
	return &InfraredProcessor{
		BaseProcessor: BaseProcessor{name: "infrared"},
	}
}

// Process simulates infrared film effect.
func (p *InfraredProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	return color.RGBA{
		R: originalColor.B,
		G: originalColor.G,
		B: originalColor.R,
		A: originalColor.A,
	}, nil
}

// RandomProcessor randomly selects from a pool of processors.
type RandomProcessor struct {
	BaseProcessor
	processors []Processor
	randFunc   func(n int) int
}

// NewRandomProcessor creates a processor with injectable processor pool.
func NewRandomProcessor(processors []Processor) *RandomProcessor {
	if processors == nil {
		// Default pool matching original implementation
		processors = []Processor{
			NewGrayProcessor(),
			NewRBGProcessor(),
			NewGBRProcessor(),
			NewGRBProcessor(),
			NewBRGProcessor(),
			NewBGRProcessor(),
			NewXGBProcessor(),
			NewRXBProcessor(),
			NewRGXProcessor(),
			NewRXXProcessor(),
			NewXGXProcessor(),
			NewXXBProcessor(),
		}
	}
	return &RandomProcessor{
		BaseProcessor: BaseProcessor{name: "rnd"},
		processors:    processors,
		randFunc:      rand.Intn,
	}
}

// WithRandFunc allows injecting a custom random function for testing.
func (p *RandomProcessor) WithRandFunc(f func(n int) int) *RandomProcessor {
	p.randFunc = f
	return p
}

// Process randomly selects a processor and applies it.
func (p *RandomProcessor) Process(pixel color.Color) (color.Color, error) {
	idx := p.randFunc(len(p.processors))
	return p.processors[idx].Process(pixel)
}
