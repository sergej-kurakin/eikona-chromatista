package cmd

import (
	"image"
	"image/color"
	"log"

	"github.com/disintegration/imaging"
	"github.com/sergej-kurakin/eikona-chromatista/processor"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(infraredCmd)
	rootCmd.AddCommand(infrared2Cmd)
}

func swapRedBlue(img image.Image) *image.NRGBA {
	bounds := img.Bounds()
	dst := image.NewNRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			// Swap red and blue channels
			dst.Set(x, y, color.NRGBA{
				R: uint8(b >> 8),
				G: uint8(g >> 8),
				B: uint8(r >> 8),
				A: uint8(a >> 8),
			})
		}
	}

	return dst
}

var infraredCmd = &cobra.Command{
	Use:   "infrared",
	Short: "Infrared Experiments",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.InfraredProcessor, "infrared")
	},
}

var infrared2Cmd = &cobra.Command{
	Use:   "infrared2",
	Short: "Infrared 2 Experiments",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		src, err := imaging.Open(args[0])

		if err != nil {
			log.Fatalf("failed to open image: %v", err)
		}

		dst := swapRedBlue(src)

		newImgPath := newFileName(args[0], "infrared2")

		err = imaging.Save(dst, newImgPath)
		if err != nil {
			log.Fatalf("failed to save image: %v", err)
		}

		log.Println("Image processing complete")
	},
}
