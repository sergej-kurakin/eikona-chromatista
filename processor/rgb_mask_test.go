package processor_test

import (
	"image/color"
	"testing"

	"github.com/sergej-kurakin/eikona-chromatista/processor"
)

func TestRGBMaskProcessors(t *testing.T) {
	// Input: R=100, G=150, B=200, A=255
	input := color.RGBA{R: 100, G: 150, B: 200, A: 255}

	tests := []struct {
		name      string
		processor processor.Processor
		expectedR uint8
		expectedG uint8
		expectedB uint8
	}{
		// X in first position (R=0)
		{"XGBProcessor", processor.NewXGBProcessor(), 0, 150, 200},
		{"XRGProcessor", processor.NewXRGProcessor(), 0, 100, 150},
		{"XRBProcessor", processor.NewXRBProcessor(), 0, 100, 200},
		{"XGRProcessor", processor.NewXGRProcessor(), 0, 150, 100},
		{"XBRProcessor", processor.NewXBRProcessor(), 0, 200, 100},
		{"XBGProcessor", processor.NewXBGProcessor(), 0, 200, 150},

		// X in second position (G=0)
		{"RXBProcessor", processor.NewRXBProcessor(), 100, 0, 200},
		{"RXGProcessor", processor.NewRXGProcessor(), 100, 0, 150},
		{"GXRProcessor", processor.NewGXRProcessor(), 150, 0, 100},
		{"GXBProcessor", processor.NewGXBProcessor(), 150, 0, 200},
		{"BXRProcessor", processor.NewBXRProcessor(), 200, 0, 100},
		{"BXGProcessor", processor.NewBXGProcessor(), 200, 0, 150},

		// X in third position (B=0)
		{"RGXProcessor", processor.NewRGXProcessor(), 100, 150, 0},
		{"RBXProcessor", processor.NewRBXProcessor(), 100, 200, 0},
		{"GRXProcessor", processor.NewGRXProcessor(), 150, 100, 0},
		{"GBXProcessor", processor.NewGBXProcessor(), 150, 200, 0},
		{"BRXProcessor", processor.NewBRXProcessor(), 200, 100, 0},
		{"BGXProcessor", processor.NewBGXProcessor(), 200, 150, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.processor.Process(input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			rgba, ok := result.(color.RGBA)
			if !ok {
				t.Fatal("result is not color.RGBA")
			}

			if rgba.R != tt.expectedR {
				t.Errorf("expected R=%d, got R=%d", tt.expectedR, rgba.R)
			}
			if rgba.G != tt.expectedG {
				t.Errorf("expected G=%d, got G=%d", tt.expectedG, rgba.G)
			}
			if rgba.B != tt.expectedB {
				t.Errorf("expected B=%d, got B=%d", tt.expectedB, rgba.B)
			}
			if rgba.A != 255 {
				t.Errorf("expected A=255, got A=%d", rgba.A)
			}
		})
	}
}

func TestRGBMaskProcessors_Name(t *testing.T) {
	tests := []struct {
		processor    processor.Processor
		expectedName string
	}{
		{processor.NewXGBProcessor(), "xgb"},
		{processor.NewRXBProcessor(), "rxb"},
		{processor.NewRGXProcessor(), "rgx"},
		{processor.NewRBXProcessor(), "rbx"},
		{processor.NewGRXProcessor(), "grx"},
		{processor.NewGBXProcessor(), "gbx"},
		{processor.NewBRXProcessor(), "brx"},
		{processor.NewBGXProcessor(), "bgx"},
		{processor.NewRXGProcessor(), "rxg"},
		{processor.NewGXRProcessor(), "gxr"},
		{processor.NewGXBProcessor(), "gxb"},
		{processor.NewBXRProcessor(), "bxr"},
		{processor.NewBXGProcessor(), "bxg"},
		{processor.NewXRGProcessor(), "xrg"},
		{processor.NewXRBProcessor(), "xrb"},
		{processor.NewXGRProcessor(), "xgr"},
		{processor.NewXBRProcessor(), "xbr"},
		{processor.NewXBGProcessor(), "xbg"},
	}

	for _, tt := range tests {
		t.Run(tt.expectedName, func(t *testing.T) {
			if tt.processor.Name() != tt.expectedName {
				t.Errorf("expected name '%s', got '%s'", tt.expectedName, tt.processor.Name())
			}
		})
	}
}

func TestRGBMaskProcessors_PreservesAlpha(t *testing.T) {
	processors := []processor.Processor{
		processor.NewXGBProcessor(),
		processor.NewRXBProcessor(),
		processor.NewRGXProcessor(),
		processor.NewRBXProcessor(),
		processor.NewGRXProcessor(),
		processor.NewGBXProcessor(),
		processor.NewBRXProcessor(),
		processor.NewBGXProcessor(),
		processor.NewRXGProcessor(),
		processor.NewGXRProcessor(),
		processor.NewGXBProcessor(),
		processor.NewBXRProcessor(),
		processor.NewBXGProcessor(),
		processor.NewXRGProcessor(),
		processor.NewXRBProcessor(),
		processor.NewXGRProcessor(),
		processor.NewXBRProcessor(),
		processor.NewXBGProcessor(),
	}

	input := color.RGBA{R: 50, G: 100, B: 150, A: 128}

	for _, proc := range processors {
		t.Run(proc.Name(), func(t *testing.T) {
			result, err := proc.Process(input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			rgba, ok := result.(color.RGBA)
			if !ok {
				t.Fatal("result is not color.RGBA")
			}

			if rgba.A != 128 {
				t.Errorf("alpha channel not preserved: expected 128, got %d", rgba.A)
			}
		})
	}
}
