package processor_test

import (
	"image"
	"image/color"
	"testing"

	"github.com/sergej-kurakin/eikona-chromatista/processor"
)

func TestChannelMerger_Merge(t *testing.T) {
	redImg := image.NewRGBA(image.Rect(0, 0, 2, 1))
	greenImg := image.NewRGBA(image.Rect(0, 0, 2, 1))
	blueImg := image.NewRGBA(image.Rect(0, 0, 2, 1))

	redImg.Set(0, 0, color.RGBA{R: 10, G: 99, B: 99, A: 255})
	greenImg.Set(0, 0, color.RGBA{R: 88, G: 20, B: 88, A: 255})
	blueImg.Set(0, 0, color.RGBA{R: 77, G: 77, B: 30, A: 255})

	redImg.Set(1, 0, color.RGBA{R: 40, G: 55, B: 55, A: 255})
	greenImg.Set(1, 0, color.RGBA{R: 66, G: 50, B: 66, A: 255})
	blueImg.Set(1, 0, color.RGBA{R: 77, G: 77, B: 60, A: 255})

	merger := processor.NewChannelMerger()
	result, err := merger.Merge(redImg, greenImg, blueImg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	px00, ok := color.RGBAModel.Convert(result.At(0, 0)).(color.RGBA)
	if !ok {
		t.Fatal("result pixel is not color.RGBA")
	}

	if px00.R != 10 || px00.G != 20 || px00.B != 30 || px00.A != 255 {
		t.Errorf("expected [10,20,30,255], got [%d,%d,%d,%d]", px00.R, px00.G, px00.B, px00.A)
	}

	px10, ok := color.RGBAModel.Convert(result.At(1, 0)).(color.RGBA)
	if !ok {
		t.Fatal("result pixel is not color.RGBA")
	}

	if px10.R != 40 || px10.G != 50 || px10.B != 60 || px10.A != 255 {
		t.Errorf("expected [40,50,60,255], got [%d,%d,%d,%d]", px10.R, px10.G, px10.B, px10.A)
	}
}

func TestChannelMerger_Merge_ErrOnMismatchedBounds(t *testing.T) {
	redImg := image.NewRGBA(image.Rect(0, 0, 2, 2))
	greenImg := image.NewRGBA(image.Rect(0, 0, 1, 2))
	blueImg := image.NewRGBA(image.Rect(0, 0, 2, 2))

	merger := processor.NewChannelMerger()
	_, err := merger.Merge(redImg, greenImg, blueImg)
	if err == nil {
		t.Fatal("expected error for mismatched bounds")
	}
}