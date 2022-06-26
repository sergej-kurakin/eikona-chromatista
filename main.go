package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type Processor struct {
	suffix          string
	color_processor func(color.Color) color.Color
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	var processors [13]Processor
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

	if len(os.Args) < 2 {
		log.Fatalln("Image path is required")
	}
	imgPath := os.Args[1]

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
