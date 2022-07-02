package cmd

import (
	"fmt"

	"image"
	"image/color"
	"image/jpeg"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/spf13/cobra"
)

var Source string

func init() {
	rootCmd.AddCommand(allCmd)
	allCmd.Flags().StringVarP(&Source, "source", "s", "", "Source image file to read from")
	allCmd.MarkFlagRequired("source")
}

var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Make all Image Color Manipulations",
	Long:  `Make all Image Color Manipulations, bla... bla... bla...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("All")
		fmt.Println("Source:", Source)
		process(Source)
	},
}

type Processor struct {
	suffix          string
	color_processor func(color.Color) color.Color
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func process(imgPath string) {

	var processors [20]Processor
	processors[0] = Processor{suffix: "gray", color_processor: make_gray_color}
	processors[1] = Processor{suffix: "rbg", color_processor: make_RBG_color}
	processors[2] = Processor{suffix: "gbr", color_processor: make_GBR_color}
	processors[3] = Processor{suffix: "grb", color_processor: make_GRB_color}
	processors[4] = Processor{suffix: "brg", color_processor: make_BRG_color}
	processors[5] = Processor{suffix: "bgr", color_processor: make_BGR_color}
	processors[6] = Processor{suffix: "gb", color_processor: make_GB_color}
	processors[7] = Processor{suffix: "rb", color_processor: make_RB_color}
	processors[8] = Processor{suffix: "rg", color_processor: make_RG_color}
	processors[9] = Processor{suffix: "r", color_processor: make_R_color}
	processors[10] = Processor{suffix: "g", color_processor: make_G_color}
	processors[11] = Processor{suffix: "b", color_processor: make_B_color}
	processors[12] = Processor{suffix: "rnd", color_processor: make_random_color}
	processors[13] = Processor{suffix: "photometric_grayscale", color_processor: make_photometric_grayscale}
	processors[14] = Processor{suffix: "photometric_graychrome", color_processor: make_photometric_graychrome}
	processors[15] = Processor{suffix: "photometric_graychrome_negative", color_processor: make_photometric_graychrome_negative}
	processors[16] = Processor{suffix: "negative", color_processor: make_negative}
	processors[17] = Processor{suffix: "redscale", color_processor: make_photometric_redscale}
	processors[18] = Processor{suffix: "greenscale", color_processor: make_photometric_greenscale}
	processors[19] = Processor{suffix: "bluescale", color_processor: make_photometric_bluescale}

	f, err := os.Open(imgPath)
	check(err)
	defer f.Close()

	img, err := jpeg.Decode(f)

	check(err)

	var wg sync.WaitGroup

	for i := 0; i < len(processors); i++ {
		wg.Add(1)
		k := i
		go func() {
			defer wg.Done()
			process_image(img, imgPath, processors[k].color_processor, processors[k].suffix)
		}()
	}

	wg.Wait()
	fmt.Printf("Processing finished\n")
}

func process_image(img image.Image, imgPath string, color_processor func(color.Color) color.Color, resultSuffix string) {
	wImg := process_pixels(img, color_processor)
	newImgPath := newFileName(imgPath, resultSuffix)
	export(newImgPath, wImg)
}

func process_pixels(img image.Image, color_processor func(color.Color) color.Color) image.Image {
	size := img.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	wImg := image.NewRGBA(rect)

	for x := 0; x < size.X; x++ {
		// and now loop thorough all of this x's y
		for y := 0; y < size.Y; y++ {
			pixel := img.At(x, y)
			c := color_processor(color.RGBAModel.Convert(pixel).(color.RGBA))
			wImg.Set(x, y, c)
		}
	}

	return wImg
}

func make_gray_color(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	// Offset colors a little, adjust it to your taste
	r := float64(originalColor.R) * 0.92126
	g := float64(originalColor.G) * 0.97152
	b := float64(originalColor.B) * 0.90722
	// average
	grey := uint8((r + g + b) / 3)
	c := color.RGBA{
		R: grey, G: grey, B: grey, A: originalColor.A,
	}
	return c
}

func make_GBR_color(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: originalColor.G, G: originalColor.B, B: originalColor.R, A: originalColor.A,
	}
	return c
}

func make_GRB_color(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: originalColor.G, G: originalColor.R, B: originalColor.B, A: originalColor.A,
	}
	return c
}

func make_BRG_color(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: originalColor.B, G: originalColor.R, B: originalColor.G, A: originalColor.A,
	}
	return c
}

func make_BGR_color(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: originalColor.B, G: originalColor.G, B: originalColor.R, A: originalColor.A,
	}
	return c
}

func make_RBG_color(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: originalColor.R, G: originalColor.B, B: originalColor.G, A: originalColor.A,
	}
	return c
}

func make_GB_color(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: 0, G: originalColor.G, B: originalColor.B, A: originalColor.A,
	}
	return c
}

func make_RB_color(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: originalColor.R, G: 0, B: originalColor.B, A: originalColor.A,
	}
	return c
}

func make_RG_color(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: originalColor.R, G: originalColor.G, B: 0, A: originalColor.A,
	}
	return c
}

func make_R_color(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: originalColor.R, G: 0, B: 0, A: originalColor.A,
	}
	return c
}

func make_G_color(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: 0, G: originalColor.G, B: 0, A: originalColor.A,
	}
	return c
}

func make_B_color(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: 0, G: 0, B: originalColor.B, A: originalColor.A,
	}
	return c
}

func make_random_color(pixel color.Color) color.Color {

	var processors [12]func(color.Color) color.Color
	processors[0] = make_gray_color
	processors[1] = make_RBG_color
	processors[2] = make_GBR_color
	processors[3] = make_GRB_color
	processors[4] = make_BRG_color
	processors[5] = make_BGR_color
	processors[6] = make_GB_color
	processors[7] = make_RB_color
	processors[8] = make_RG_color
	processors[9] = make_R_color
	processors[10] = make_G_color
	processors[11] = make_B_color

	processor_num := rand.Intn(len(processors))

	return processors[processor_num](pixel)
}

func make_photometric_grayscale(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	// Y = 0.2126 R + 0.7152 G + 0.0722 B
	Y := 0.2126*float32(originalColor.R) + 0.7152*float32(originalColor.G) + 0.0722*float32(originalColor.B)
	c := color.RGBA{
		R: uint8(Y), G: uint8(Y), B: uint8(Y), A: originalColor.A,
	}

	return c
}

func make_photometric_redscale(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	// Y = 0.2126 R + 0.7152 G + 0.0722 B
	Y := 0.2126*float32(originalColor.R) + 0.7152*float32(originalColor.G) + 0.0722*float32(originalColor.B)
	c := color.RGBA{
		R: uint8(Y), G: 0, B: 0, A: originalColor.A,
	}

	return c
}

func make_photometric_greenscale(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	// Y = 0.2126 R + 0.7152 G + 0.0722 B
	Y := 0.2126*float32(originalColor.R) + 0.7152*float32(originalColor.G) + 0.0722*float32(originalColor.B)
	c := color.RGBA{
		R: 0, G: uint8(Y), B: 0, A: originalColor.A,
	}

	return c
}

func make_photometric_bluescale(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	// Y = 0.2126 R + 0.7152 G + 0.0722 B
	Y := 0.2126*float32(originalColor.R) + 0.7152*float32(originalColor.G) + 0.0722*float32(originalColor.B)
	c := color.RGBA{
		R: 0, G: 0, B: uint8(Y), A: originalColor.A,
	}

	return c
}

func make_photometric_graychrome(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	// Y = 0.2126 R + 0.7152 G + 0.0722 B
	Y := 0.2126*float32(originalColor.R) + 0.7152*float32(originalColor.G) + 0.0722*float32(originalColor.B)

	k := uint8(Y)
	var c color.RGBA

	// 0 < x <= 64 -- 0
	// 64 < x <= 128 -- 85
	// 128 < x <= 192 -- 180
	// 192 < x <= 255 -- 255

	if k <= 64 {
		c = color.RGBA{
			R: 0, G: 0, B: 0, A: originalColor.A,
		}
	} else if k <= 128 {
		c = color.RGBA{
			R: 85, G: 85, B: 85, A: originalColor.A,
		}
	} else if k <= 192 {
		c = color.RGBA{
			R: 180, G: 180, B: 180, A: originalColor.A,
		}
	} else {
		c = color.RGBA{
			R: 255, G: 255, B: 255, A: originalColor.A,
		}
	}

	return c
}

func make_negative(pixel color.Color) color.Color {
	originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
	c := color.RGBA{
		R: 255 - originalColor.R, G: 255 - originalColor.G, B: 255 - originalColor.B, A: originalColor.A,
	}
	return c
}

func make_photometric_graychrome_negative(pixel color.Color) color.Color {
	graychrome_pixel := make_photometric_graychrome(pixel)
	return make_negative(graychrome_pixel)
}

func newFileName(imgPath string, suffix string) string {
	ext := filepath.Ext(imgPath)
	name := strings.TrimSuffix(filepath.Base(imgPath), ext)
	newImagePath := fmt.Sprintf("%s/%s_%s%s", filepath.Dir(imgPath), name, suffix, ext)
	return newImagePath
}

func export(newImagePath string, wImg image.Image) {
	fg, err := os.Create(newImagePath)
	defer fg.Close()
	check(err)
	err = jpeg.Encode(fg, wImg, &jpeg.Options{Quality: 100})
	check(err)
}
