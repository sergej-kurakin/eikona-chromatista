package processor

import (
	"image/color"
)

// XGBProcessor masks the red channel to zero.
type XGBProcessor struct {
	BaseProcessor
}

// NewXGBProcessor creates a new XGBProcessor instance.
func NewXGBProcessor() *XGBProcessor {
	return &XGBProcessor{
		BaseProcessor: BaseProcessor{name: "xgb"},
	}
}

// Process masks the red channel to zero.
func (p *XGBProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	return color.RGBA{
		R: 0, G: originalColor.G, B: originalColor.B, A: originalColor.A,
	}, nil
}

// RXBProcessor masks the green channel to zero.
type RXBProcessor struct {
	BaseProcessor
}

// NewRXBProcessor creates a new RXBProcessor instance.
func NewRXBProcessor() *RXBProcessor {
	return &RXBProcessor{
		BaseProcessor: BaseProcessor{name: "rxb"},
	}
}

// Process masks the green channel to zero.
func (p *RXBProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	return color.RGBA{
		R: originalColor.R, G: 0, B: originalColor.B, A: originalColor.A,
	}, nil
}

// RGXProcessor masks the blue channel to zero.
type RGXProcessor struct {
	BaseProcessor
}

// NewRGXProcessor creates a new RGXProcessor instance.
func NewRGXProcessor() *RGXProcessor {
	return &RGXProcessor{
		BaseProcessor: BaseProcessor{name: "rgx"},
	}
}

// Process masks the blue channel to zero.
func (p *RGXProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	return color.RGBA{
		R: originalColor.R, G: originalColor.G, B: 0, A: originalColor.A,
	}, nil
}

// RBXProcessor swaps G and B, then masks blue to zero.
type RBXProcessor struct {
	BaseProcessor
}

// NewRBXProcessor creates a new RBXProcessor instance.
func NewRBXProcessor() *RBXProcessor {
	return &RBXProcessor{
		BaseProcessor: BaseProcessor{name: "rbx"},
	}
}

// Process swaps G and B, then masks blue to zero.
func (p *RBXProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	return color.RGBA{
		R: originalColor.R, G: originalColor.B, B: 0, A: originalColor.A,
	}, nil
}

// GRXProcessor swaps R and G, then masks blue to zero.
type GRXProcessor struct {
	BaseProcessor
}

// NewGRXProcessor creates a new GRXProcessor instance.
func NewGRXProcessor() *GRXProcessor {
	return &GRXProcessor{
		BaseProcessor: BaseProcessor{name: "grx"},
	}
}

// Process swaps R and G, then masks blue to zero.
func (p *GRXProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	return color.RGBA{
		R: originalColor.G, G: originalColor.R, B: 0, A: originalColor.A,
	}, nil
}

// GBXProcessor moves G to R, B to G, masks blue to zero.
type GBXProcessor struct {
	BaseProcessor
}

// NewGBXProcessor creates a new GBXProcessor instance.
func NewGBXProcessor() *GBXProcessor {
	return &GBXProcessor{
		BaseProcessor: BaseProcessor{name: "gbx"},
	}
}

// Process moves G to R, B to G, masks blue to zero.
func (p *GBXProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	return color.RGBA{
		R: originalColor.G, G: originalColor.B, B: 0, A: originalColor.A,
	}, nil
}

// BRXProcessor moves B to R, R to G, masks blue to zero.
type BRXProcessor struct {
	BaseProcessor
}

// NewBRXProcessor creates a new BRXProcessor instance.
func NewBRXProcessor() *BRXProcessor {
	return &BRXProcessor{
		BaseProcessor: BaseProcessor{name: "brx"},
	}
}

// Process moves B to R, R to G, masks blue to zero.
func (p *BRXProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	return color.RGBA{
		R: originalColor.B, G: originalColor.R, B: 0, A: originalColor.A,
	}, nil
}

// BGXProcessor moves B to R, G stays, masks blue to zero.
type BGXProcessor struct {
	BaseProcessor
}

// NewBGXProcessor creates a new BGXProcessor instance.
func NewBGXProcessor() *BGXProcessor {
	return &BGXProcessor{
		BaseProcessor: BaseProcessor{name: "bgx"},
	}
}

// Process moves B to R, G stays, masks blue to zero.
func (p *BGXProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	return color.RGBA{
		R: originalColor.B, G: originalColor.G, B: 0, A: originalColor.A,
	}, nil
}

// RXGProcessor keeps R, masks green to zero, moves G to B.
type RXGProcessor struct {
	BaseProcessor
}

// NewRXGProcessor creates a new RXGProcessor instance.
func NewRXGProcessor() *RXGProcessor {
	return &RXGProcessor{
		BaseProcessor: BaseProcessor{name: "rxg"},
	}
}

// Process keeps R, masks green to zero, moves G to B.
func (p *RXGProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	return color.RGBA{
		R: originalColor.R, G: 0, B: originalColor.G, A: originalColor.A,
	}, nil
}

// GXRProcessor moves G to R, masks green to zero, R to B.
type GXRProcessor struct {
	BaseProcessor
}

// NewGXRProcessor creates a new GXRProcessor instance.
func NewGXRProcessor() *GXRProcessor {
	return &GXRProcessor{
		BaseProcessor: BaseProcessor{name: "gxr"},
	}
}

// Process moves G to R, masks green to zero, R to B.
func (p *GXRProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	return color.RGBA{
		R: originalColor.G, G: 0, B: originalColor.R, A: originalColor.A,
	}, nil
}

// GXBProcessor moves G to R, masks green to zero, B stays.
type GXBProcessor struct {
	BaseProcessor
}

// NewGXBProcessor creates a new GXBProcessor instance.
func NewGXBProcessor() *GXBProcessor {
	return &GXBProcessor{
		BaseProcessor: BaseProcessor{name: "gxb"},
	}
}

// Process moves G to R, masks green to zero, B stays.
func (p *GXBProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	return color.RGBA{
		R: originalColor.G, G: 0, B: originalColor.B, A: originalColor.A,
	}, nil
}

// BXRProcessor moves B to R, masks green to zero, R to B.
type BXRProcessor struct {
	BaseProcessor
}

// NewBXRProcessor creates a new BXRProcessor instance.
func NewBXRProcessor() *BXRProcessor {
	return &BXRProcessor{
		BaseProcessor: BaseProcessor{name: "bxr"},
	}
}

// Process moves B to R, masks green to zero, R to B.
func (p *BXRProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	return color.RGBA{
		R: originalColor.B, G: 0, B: originalColor.R, A: originalColor.A,
	}, nil
}

// BXGProcessor moves B to R, masks green to zero, G to B.
type BXGProcessor struct {
	BaseProcessor
}

// NewBXGProcessor creates a new BXGProcessor instance.
func NewBXGProcessor() *BXGProcessor {
	return &BXGProcessor{
		BaseProcessor: BaseProcessor{name: "bxg"},
	}
}

// Process moves B to R, masks green to zero, G to B.
func (p *BXGProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	return color.RGBA{
		R: originalColor.B, G: 0, B: originalColor.G, A: originalColor.A,
	}, nil
}

// XRGProcessor masks red to zero, R to G, G to B.
type XRGProcessor struct {
	BaseProcessor
}

// NewXRGProcessor creates a new XRGProcessor instance.
func NewXRGProcessor() *XRGProcessor {
	return &XRGProcessor{
		BaseProcessor: BaseProcessor{name: "xrg"},
	}
}

// Process masks red to zero, R to G, G to B.
func (p *XRGProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	return color.RGBA{
		R: 0, G: originalColor.R, B: originalColor.G, A: originalColor.A,
	}, nil
}

// XRBProcessor masks red to zero, R to G, B stays.
type XRBProcessor struct {
	BaseProcessor
}

// NewXRBProcessor creates a new XRBProcessor instance.
func NewXRBProcessor() *XRBProcessor {
	return &XRBProcessor{
		BaseProcessor: BaseProcessor{name: "xrb"},
	}
}

// Process masks red to zero, R to G, B stays.
func (p *XRBProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	return color.RGBA{
		R: 0, G: originalColor.R, B: originalColor.B, A: originalColor.A,
	}, nil
}

// XGRProcessor masks red to zero, G stays, R to B.
type XGRProcessor struct {
	BaseProcessor
}

// NewXGRProcessor creates a new XGRProcessor instance.
func NewXGRProcessor() *XGRProcessor {
	return &XGRProcessor{
		BaseProcessor: BaseProcessor{name: "xgr"},
	}
}

// Process masks red to zero, G stays, R to B.
func (p *XGRProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	return color.RGBA{
		R: 0, G: originalColor.G, B: originalColor.R, A: originalColor.A,
	}, nil
}

// XBRProcessor masks red to zero, B to G, R to B.
type XBRProcessor struct {
	BaseProcessor
}

// NewXBRProcessor creates a new XBRProcessor instance.
func NewXBRProcessor() *XBRProcessor {
	return &XBRProcessor{
		BaseProcessor: BaseProcessor{name: "xbr"},
	}
}

// Process masks red to zero, B to G, R to B.
func (p *XBRProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	return color.RGBA{
		R: 0, G: originalColor.B, B: originalColor.R, A: originalColor.A,
	}, nil
}

// XBGProcessor masks red to zero, B to G, G to B.
type XBGProcessor struct {
	BaseProcessor
}

// NewXBGProcessor creates a new XBGProcessor instance.
func NewXBGProcessor() *XBGProcessor {
	return &XBGProcessor{
		BaseProcessor: BaseProcessor{name: "xbg"},
	}
}

// Process masks red to zero, B to G, G to B.
func (p *XBGProcessor) Process(pixel color.Color) (color.Color, error) {
	originalColor, err := p.ToRGBA(pixel)
	if err != nil {
		return nil, err
	}

	return color.RGBA{
		R: 0, G: originalColor.B, B: originalColor.G, A: originalColor.A,
	}, nil
}
