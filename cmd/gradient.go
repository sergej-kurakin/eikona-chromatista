package cmd

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "gradient",
	Short: "Gradient some sample colors",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating gradients")

		size := image.Point{X: 1024, Y: 1024}
		rect := image.Rect(0, 0, size.X, size.Y)
		wImg := image.NewRGBA(rect)

		for x := 0; x < size.X; x++ {
			// and now loop thorough all of this x's y
			for y := 0; y < size.Y; y++ {
				c := color.RGBA{
					R: 130, G: 130, B: 130, A: 255,
				}
				wImg.Set(x, y, c)
			}
		}

		for x := 0; x < 256; x++ {
			// and now loop thorough all of this x's y
			for y := 0; y < 256; y++ {
				c := color.RGBA{
					R: uint8(y), G: uint8(y), B: uint8(y), A: 255,
				}
				wImg.Set(x, y, c)
			}

			for y := 256; y < 512; y++ {
				code := uint8(y - 256)
				c := color.RGBA{
					R: code, G: 0, B: 0, A: 255,
				}
				wImg.Set(x, y, c)
			}

			for y := 512; y < 768; y++ {
				code := uint8(y - 512)
				c := color.RGBA{
					R: 0, G: code, B: 0, A: 255,
				}
				wImg.Set(x, y, c)
			}

			for y := 768; y < 1024; y++ {
				code := uint8(y - 768)
				c := color.RGBA{
					R: 0, G: 0, B: code, A: 255,
				}
				wImg.Set(x, y, c)
			}

		}

		// firstColor := color.RGBA{
		// 	R: 0, G: 0, B: 0, A: 255,
		// }

		// secondColor := color.RGBA{
		// 	R: 255, G: 255, B: 255, A: 255,
		// }

		firstColor := color.RGBA{
			R: 250, G: 250, B: 110, A: 255,
		}

		secondColor := color.RGBA{
			R: 42, G: 72, B: 88, A: 255,
		}

		steps := size.Y

		for x := 256; x < 512; x++ {
			for y := 0; y < size.Y; y++ {

				ratio := float32(y) / float32(steps)

				r := float32(firstColor.R)*ratio + float32(secondColor.R)*(1-ratio)
				g := float32(firstColor.G)*ratio + float32(secondColor.G)*(1-ratio)
				b := float32(firstColor.B)*ratio + float32(secondColor.B)*(1-ratio)

				c := color.RGBA{
					R: uint8(r), G: uint8(g), B: uint8(b), A: 255,
				}

				wImg.Set(x, y, c)
			}
		}

		fg, err := os.Create(args[0])
		defer fg.Close()
		check(err)
		err = jpeg.Encode(fg, wImg, &jpeg.Options{Quality: 100})
		check(err)
	},
}
