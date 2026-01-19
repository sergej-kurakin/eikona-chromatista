package processor_test

import (
	"image/color"
	"testing"

	"github.com/sergej-kurakin/eikona-chromatista/processor"
)

// MockLuminosityCalculator for testing photometric processors.
type MockLuminosityCalculator struct {
	ReturnValue uint8
}

func (m *MockLuminosityCalculator) Calculate(_ color.RGBA) uint8 {
	return m.ReturnValue
}

func TestMonochromeProcessors_Name(t *testing.T) {
	tests := []struct {
		processor    processor.Processor
		expectedName string
	}{
		{processor.NewRXXProcessor(), "red"},
		{processor.NewXGXProcessor(), "green"},
		{processor.NewXXBProcessor(), "blue"},
		{processor.NewPhotometricRedscaleProcessor(nil), "redscale"},
		{processor.NewPhotometricGreenscaleProcessor(nil), "greenscale"},
		{processor.NewPhotometricBluescaleProcessor(nil), "bluescale"},
	}

	for _, tt := range tests {
		t.Run(tt.expectedName, func(t *testing.T) {
			if tt.processor.Name() != tt.expectedName {
				t.Errorf("expected name '%s', got '%s'", tt.expectedName, tt.processor.Name())
			}
		})
	}
}

func TestRXXProcessor_Process(t *testing.T) {
	proc := processor.NewRXXProcessor()

	input := color.RGBA{R: 100, G: 150, B: 200, A: 255}
	result, err := proc.Process(input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	rgba, ok := result.(color.RGBA)
	if !ok {
		t.Fatal("result is not color.RGBA")
	}

	// RXX keeps only red: R=R, G=0, B=0
	if rgba.R != 100 {
		t.Errorf("expected R=100, got R=%d", rgba.R)
	}
	if rgba.G != 0 {
		t.Errorf("expected G=0, got G=%d", rgba.G)
	}
	if rgba.B != 0 {
		t.Errorf("expected B=0, got B=%d", rgba.B)
	}
	if rgba.A != 255 {
		t.Errorf("expected A=255, got A=%d", rgba.A)
	}
}

func TestXGXProcessor_Process(t *testing.T) {
	proc := processor.NewXGXProcessor()

	input := color.RGBA{R: 100, G: 150, B: 200, A: 255}
	result, err := proc.Process(input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	rgba, ok := result.(color.RGBA)
	if !ok {
		t.Fatal("result is not color.RGBA")
	}

	// XGX keeps only green: R=0, G=G, B=0
	if rgba.R != 0 {
		t.Errorf("expected R=0, got R=%d", rgba.R)
	}
	if rgba.G != 150 {
		t.Errorf("expected G=150, got G=%d", rgba.G)
	}
	if rgba.B != 0 {
		t.Errorf("expected B=0, got B=%d", rgba.B)
	}
	if rgba.A != 255 {
		t.Errorf("expected A=255, got A=%d", rgba.A)
	}
}

func TestXXBProcessor_Process(t *testing.T) {
	proc := processor.NewXXBProcessor()

	input := color.RGBA{R: 100, G: 150, B: 200, A: 255}
	result, err := proc.Process(input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	rgba, ok := result.(color.RGBA)
	if !ok {
		t.Fatal("result is not color.RGBA")
	}

	// XXB keeps only blue: R=0, G=0, B=B
	if rgba.R != 0 {
		t.Errorf("expected R=0, got R=%d", rgba.R)
	}
	if rgba.G != 0 {
		t.Errorf("expected G=0, got G=%d", rgba.G)
	}
	if rgba.B != 200 {
		t.Errorf("expected B=200, got B=%d", rgba.B)
	}
	if rgba.A != 255 {
		t.Errorf("expected A=255, got A=%d", rgba.A)
	}
}

func TestPhotometricRedscaleProcessor_Process(t *testing.T) {
	mockCalc := &MockLuminosityCalculator{ReturnValue: 128}
	proc := processor.NewPhotometricRedscaleProcessor(mockCalc)

	input := color.RGBA{R: 100, G: 150, B: 200, A: 255}
	result, err := proc.Process(input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	rgba, ok := result.(color.RGBA)
	if !ok {
		t.Fatal("result is not color.RGBA")
	}

	// Redscale puts luminosity in R channel only
	if rgba.R != 128 {
		t.Errorf("expected R=128, got R=%d", rgba.R)
	}
	if rgba.G != 0 {
		t.Errorf("expected G=0, got G=%d", rgba.G)
	}
	if rgba.B != 0 {
		t.Errorf("expected B=0, got B=%d", rgba.B)
	}
	if rgba.A != 255 {
		t.Errorf("expected A=255, got A=%d", rgba.A)
	}
}

func TestPhotometricGreenscaleProcessor_Process(t *testing.T) {
	mockCalc := &MockLuminosityCalculator{ReturnValue: 128}
	proc := processor.NewPhotometricGreenscaleProcessor(mockCalc)

	input := color.RGBA{R: 100, G: 150, B: 200, A: 255}
	result, err := proc.Process(input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	rgba, ok := result.(color.RGBA)
	if !ok {
		t.Fatal("result is not color.RGBA")
	}

	// Greenscale puts luminosity in G channel only
	if rgba.R != 0 {
		t.Errorf("expected R=0, got R=%d", rgba.R)
	}
	if rgba.G != 128 {
		t.Errorf("expected G=128, got G=%d", rgba.G)
	}
	if rgba.B != 0 {
		t.Errorf("expected B=0, got B=%d", rgba.B)
	}
	if rgba.A != 255 {
		t.Errorf("expected A=255, got A=%d", rgba.A)
	}
}

func TestPhotometricBluescaleProcessor_Process(t *testing.T) {
	mockCalc := &MockLuminosityCalculator{ReturnValue: 128}
	proc := processor.NewPhotometricBluescaleProcessor(mockCalc)

	input := color.RGBA{R: 100, G: 150, B: 200, A: 255}
	result, err := proc.Process(input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	rgba, ok := result.(color.RGBA)
	if !ok {
		t.Fatal("result is not color.RGBA")
	}

	// Bluescale puts luminosity in B channel only
	if rgba.R != 0 {
		t.Errorf("expected R=0, got R=%d", rgba.R)
	}
	if rgba.G != 0 {
		t.Errorf("expected G=0, got G=%d", rgba.G)
	}
	if rgba.B != 128 {
		t.Errorf("expected B=128, got B=%d", rgba.B)
	}
	if rgba.A != 255 {
		t.Errorf("expected A=255, got A=%d", rgba.A)
	}
}

func TestPhotometricProcessors_DefaultLuminosityCalculator(t *testing.T) {
	// Test that nil uses DefaultLuminosityCalculator
	// Input: R=100, G=150, B=50
	// Luminosity = 0.2126*100 + 0.7152*150 + 0.0722*50 = 21.26 + 107.28 + 3.61 = 132.15 ≈ 132
	input := color.RGBA{R: 100, G: 150, B: 50, A: 255}
	expectedLuminosity := uint8(132)

	t.Run("Redscale", func(t *testing.T) {
		proc := processor.NewPhotometricRedscaleProcessor(nil)
		result, _ := proc.Process(input)
		rgba := result.(color.RGBA)
		if rgba.R != expectedLuminosity {
			t.Errorf("expected R=%d, got R=%d", expectedLuminosity, rgba.R)
		}
	})

	t.Run("Greenscale", func(t *testing.T) {
		proc := processor.NewPhotometricGreenscaleProcessor(nil)
		result, _ := proc.Process(input)
		rgba := result.(color.RGBA)
		if rgba.G != expectedLuminosity {
			t.Errorf("expected G=%d, got G=%d", expectedLuminosity, rgba.G)
		}
	})

	t.Run("Bluescale", func(t *testing.T) {
		proc := processor.NewPhotometricBluescaleProcessor(nil)
		result, _ := proc.Process(input)
		rgba := result.(color.RGBA)
		if rgba.B != expectedLuminosity {
			t.Errorf("expected B=%d, got B=%d", expectedLuminosity, rgba.B)
		}
	})
}

func TestMonochromeProcessors_PreservesAlpha(t *testing.T) {
	processors := []processor.Processor{
		processor.NewRXXProcessor(),
		processor.NewXGXProcessor(),
		processor.NewXXBProcessor(),
		processor.NewPhotometricRedscaleProcessor(nil),
		processor.NewPhotometricGreenscaleProcessor(nil),
		processor.NewPhotometricBluescaleProcessor(nil),
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
