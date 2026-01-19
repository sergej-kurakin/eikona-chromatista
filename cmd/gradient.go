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
	Run: func(_ *cobra.Command, args []string) {
		fmt.Println("Creating gradients")

		size := image.Point{X: 1024, Y: 1024}
		rect := image.Rect(0, 0, size.X, size.Y)
		wImg := image.NewRGBA(rect)

		c := color.RGBA{
			R: 130, G: 130, B: 130, A: 255,
		}

		fillWithGray(*wImg, c, size.X, size.Y)

		drawBtoWVerticalGradient(*wImg, 0, 256, 0, 256)

		rg1 := color.RGBA{
			R: 0, G: 0, B: 0, A: 255,
		}
		rg2 := color.RGBA{
			R: 255, G: 0, B: 0, A: 255,
		}

		fillRectangle(*wImg, func(_, y int) color.RGBA {
			code := uint8(y - 256)
			steps := 256

			ratio := float32(code) / float32(steps)

			r := float32(rg1.R)*ratio + float32(rg2.R)*(1-ratio)
			g := float32(rg1.G)*ratio + float32(rg2.G)*(1-ratio)
			b := float32(rg1.B)*ratio + float32(rg2.B)*(1-ratio)

			return color.RGBA{
				R: uint8(r), G: uint8(g), B: uint8(b), A: 255,
			}
		}, 0, 256, 256, 512)

		gg1 := color.RGBA{
			R: 0, G: 0, B: 0, A: 255,
		}
		gg2 := color.RGBA{
			R: 0, G: 255, B: 0, A: 255,
		}

		fillRectangle(*wImg, func(_, y int) color.RGBA {
			code := uint8(y - 512)
			steps := 256

			ratio := float32(code) / float32(steps)

			r := float32(gg1.R)*ratio + float32(gg2.R)*(1-ratio)
			g := float32(gg1.G)*ratio + float32(gg2.G)*(1-ratio)
			b := float32(gg1.B)*ratio + float32(gg2.B)*(1-ratio)

			return color.RGBA{
				R: uint8(r), G: uint8(g), B: uint8(b), A: 255,
			}
		}, 0, 256, 512, 768)

		bg1 := color.RGBA{
			R: 0, G: 0, B: 0, A: 255,
		}
		bg2 := color.RGBA{
			R: 0, G: 0, B: 255, A: 255,
		}

		fillRectangle(*wImg, func(_, y int) color.RGBA {
			code := uint8(y - 768)
			steps := 256

			ratio := float32(code) / float32(steps)

			r := float32(bg1.R)*ratio + float32(bg2.R)*(1-ratio)
			g := float32(bg1.G)*ratio + float32(bg2.G)*(1-ratio)
			b := float32(bg1.B)*ratio + float32(bg2.B)*(1-ratio)

			return color.RGBA{
				R: uint8(r), G: uint8(g), B: uint8(b), A: 255,
			}
		}, 0, 256, 768, 1024)

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

		fillRectangle(*wImg, func(_, y int) color.RGBA {
			ratio := float32(y) / float32(steps)

			r := float32(firstColor.R)*ratio + float32(secondColor.R)*(1-ratio)
			g := float32(firstColor.G)*ratio + float32(secondColor.G)*(1-ratio)
			b := float32(firstColor.B)*ratio + float32(secondColor.B)*(1-ratio)

			return color.RGBA{
				R: uint8(r), G: uint8(g), B: uint8(b), A: 255,
			}
		}, 256, 512, 0, size.Y)

		fillRectangle(*wImg, func(x, _ int) color.RGBA {
			ratio := float32(x-512) / float32(256)

			r := float32(firstColor.R)*ratio + float32(secondColor.R)*(1-ratio)
			g := float32(firstColor.G)*ratio + float32(secondColor.G)*(1-ratio)
			b := float32(firstColor.B)*ratio + float32(secondColor.B)*(1-ratio)

			return color.RGBA{
				R: uint8(r), G: uint8(g), B: uint8(b), A: 255,
			}
		}, 512, 768, 0, 256)

		fillRectangle(*wImg, func(x, _ int) color.RGBA {
			ratio := float32(x-512) / float32(256)

			r := float32(firstColor.R)*ratio + float32(secondColor.R)*(1-ratio)
			g := float32(firstColor.G)*ratio + float32(secondColor.G)*(1-ratio)
			b := float32(firstColor.B)*ratio + float32(secondColor.B)*(1-ratio)

			return color.RGBA{
				R: uint8(r), G: uint8(g), B: uint8(b), A: 255,
			}
		}, 512, 768, 512, 768)

		for x := 768; x < 1024; x++ {
			for y := 0; y < 256; y++ {
				bg := color.RGBA{
					R: 255, G: 0, B: 0, A: 255,
				}
				fg := color.RGBA{
					R: 0, G: 255, B: 0, A: 255,
				}

				fgA := float32(fg.A)
				bgA := float32(bg.A)

				fgR := float32(fg.R)
				fgG := float32(fg.G)
				fgB := float32(fg.B)

				bgR := float32(bg.R)
				bgG := float32(bg.G)
				bgB := float32(bg.B)

				a := (1-fgA/255)*(bgA/255) + (fgA / 255)

				r := ((1-fgA/255)*(bgA/255)*(bgR/255) + fgA/255*fgR/255) / a
				g := ((1-fgA/255)*(bgA/255)*(bgG/255) + fgA/255*fgG/255) / a
				b := ((1-fgA/255)*(bgA/255)*(bgB/255) + fgA/255*fgB/255) / a

				a *= 255
				r *= 255
				g *= 255
				b *= 255

				c := color.RGBA{
					R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a),
				}

				wImg.Set(x, y, c)
			}
		}

		fg, err := os.Create(args[0])
		check(err)
		defer fg.Close() //nolint:errcheck // ignore error for defer close
		err = jpeg.Encode(fg, wImg, &jpeg.Options{Quality: 100})
		check(err)
	},
}

func fillWithGray(wImg image.RGBA, bgColor color.RGBA, X, Y int) {
	for x := 0; x < X; x++ {
		for y := 0; y < Y; y++ {
			wImg.Set(x, y, bgColor)
		}
	}
}

func drawBtoWVerticalGradient(wImg image.RGBA, x1, x2, y1, y2 int) {
	for x := x1; x < x2; x++ {
		// and now loop thorough all of this x's y
		for y := y1; y < y2; y++ {
			c := color.RGBA{
				R: uint8(y), G: uint8(y), B: uint8(y), A: 255,
			}
			wImg.Set(x, y, c)
		}
	}
}

func fillRectangle(wImg image.RGBA, colorCalc func(x, y int) color.RGBA, x1, x2, y1, y2 int) {
	for x := x1; x < x2; x++ {
		for y := y1; y < y2; y++ {
			c := colorCalc(x, y)
			wImg.Set(x, y, c)
		}
	}
}
