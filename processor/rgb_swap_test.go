package processor_test

import (
	"image/color"
	"testing"

	"github.com/sergej-kurakin/eikona-chromatista/processor"
)

func TestGBRProcessor_Name(t *testing.T) {
	proc := processor.NewGBRProcessor()

	if proc.Name() != "gbr" {
		t.Errorf("expected name 'gbr', got '%s'", proc.Name())
	}
}

func TestGBRProcessor_Process(t *testing.T) {
	proc := processor.NewGBRProcessor()

	input := color.RGBA{R: 100, G: 150, B: 200, A: 255}
	result, err := proc.Process(input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	rgba, ok := result.(color.RGBA)
	if !ok {
		t.Fatal("result is not color.RGBA")
	}

	// GBR swaps: R=G, G=B, B=R
	// So input R=100, G=150, B=200 becomes R=150, G=200, B=100
	if rgba.R != 150 {
		t.Errorf("expected R=150, got R=%d", rgba.R)
	}
	if rgba.G != 200 {
		t.Errorf("expected G=200, got G=%d", rgba.G)
	}
	if rgba.B != 100 {
		t.Errorf("expected B=100, got B=%d", rgba.B)
	}
	if rgba.A != 255 {
		t.Errorf("expected A=255, got A=%d", rgba.A)
	}
}

func TestGBRProcessor_PreservesAlpha(t *testing.T) {
	proc := processor.NewGBRProcessor()

	input := color.RGBA{R: 50, G: 100, B: 150, A: 128}
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
}

func TestGRBProcessor_Name(t *testing.T) {
	proc := processor.NewGRBProcessor()

	if proc.Name() != "grb" {
		t.Errorf("expected name 'grb', got '%s'", proc.Name())
	}
}

func TestGRBProcessor_Process(t *testing.T) {
	proc := processor.NewGRBProcessor()

	input := color.RGBA{R: 100, G: 150, B: 200, A: 255}
	result, err := proc.Process(input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	rgba, ok := result.(color.RGBA)
	if !ok {
		t.Fatal("result is not color.RGBA")
	}

	// GRB swaps: R=G, G=R, B=B
	// So input R=100, G=150, B=200 becomes R=150, G=100, B=200
	if rgba.R != 150 {
		t.Errorf("expected R=150, got R=%d", rgba.R)
	}
	if rgba.G != 100 {
		t.Errorf("expected G=100, got G=%d", rgba.G)
	}
	if rgba.B != 200 {
		t.Errorf("expected B=200, got B=%d", rgba.B)
	}
	if rgba.A != 255 {
		t.Errorf("expected A=255, got A=%d", rgba.A)
	}
}

func TestGRBProcessor_PreservesAlpha(t *testing.T) {
	proc := processor.NewGRBProcessor()

	input := color.RGBA{R: 50, G: 100, B: 150, A: 128}
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
}

func TestBRGProcessor_Name(t *testing.T) {
	proc := processor.NewBRGProcessor()

	if proc.Name() != "brg" {
		t.Errorf("expected name 'brg', got '%s'", proc.Name())
	}
}

func TestBRGProcessor_Process(t *testing.T) {
	proc := processor.NewBRGProcessor()

	input := color.RGBA{R: 100, G: 150, B: 200, A: 255}
	result, err := proc.Process(input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	rgba, ok := result.(color.RGBA)
	if !ok {
		t.Fatal("result is not color.RGBA")
	}

	// BRG swaps: R=B, G=R, B=G
	// So input R=100, G=150, B=200 becomes R=200, G=100, B=150
	if rgba.R != 200 {
		t.Errorf("expected R=200, got R=%d", rgba.R)
	}
	if rgba.G != 100 {
		t.Errorf("expected G=100, got G=%d", rgba.G)
	}
	if rgba.B != 150 {
		t.Errorf("expected B=150, got B=%d", rgba.B)
	}
	if rgba.A != 255 {
		t.Errorf("expected A=255, got A=%d", rgba.A)
	}
}

func TestBRGProcessor_PreservesAlpha(t *testing.T) {
	proc := processor.NewBRGProcessor()

	input := color.RGBA{R: 50, G: 100, B: 150, A: 128}
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
}

func TestBGRProcessor_Name(t *testing.T) {
	proc := processor.NewBGRProcessor()

	if proc.Name() != "bgr" {
		t.Errorf("expected name 'bgr', got '%s'", proc.Name())
	}
}

func TestBGRProcessor_Process(t *testing.T) {
	proc := processor.NewBGRProcessor()

	input := color.RGBA{R: 100, G: 150, B: 200, A: 255}
	result, err := proc.Process(input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	rgba, ok := result.(color.RGBA)
	if !ok {
		t.Fatal("result is not color.RGBA")
	}

	// BGR swaps: R=B, G=G, B=R
	// So input R=100, G=150, B=200 becomes R=200, G=150, B=100
	if rgba.R != 200 {
		t.Errorf("expected R=200, got R=%d", rgba.R)
	}
	if rgba.G != 150 {
		t.Errorf("expected G=150, got G=%d", rgba.G)
	}
	if rgba.B != 100 {
		t.Errorf("expected B=100, got B=%d", rgba.B)
	}
	if rgba.A != 255 {
		t.Errorf("expected A=255, got A=%d", rgba.A)
	}
}

func TestBGRProcessor_PreservesAlpha(t *testing.T) {
	proc := processor.NewBGRProcessor()

	input := color.RGBA{R: 50, G: 100, B: 150, A: 128}
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
}

func TestRBGProcessor_Name(t *testing.T) {
	proc := processor.NewRBGProcessor()

	if proc.Name() != "rbg" {
		t.Errorf("expected name 'rbg', got '%s'", proc.Name())
	}
}

func TestRBGProcessor_Process(t *testing.T) {
	proc := processor.NewRBGProcessor()

	input := color.RGBA{R: 100, G: 150, B: 200, A: 255}
	result, err := proc.Process(input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	rgba, ok := result.(color.RGBA)
	if !ok {
		t.Fatal("result is not color.RGBA")
	}

	// RBG swaps: R=R, G=B, B=G
	// So input R=100, G=150, B=200 becomes R=100, G=200, B=150
	if rgba.R != 100 {
		t.Errorf("expected R=100, got R=%d", rgba.R)
	}
	if rgba.G != 200 {
		t.Errorf("expected G=200, got G=%d", rgba.G)
	}
	if rgba.B != 150 {
		t.Errorf("expected B=150, got B=%d", rgba.B)
	}
	if rgba.A != 255 {
		t.Errorf("expected A=255, got A=%d", rgba.A)
	}
}

func TestRBGProcessor_PreservesAlpha(t *testing.T) {
	proc := processor.NewRBGProcessor()

	input := color.RGBA{R: 50, G: 100, B: 150, A: 128}
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
}
