package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
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

	wg.Add(1)
	go func() {
		defer wg.Done()
		process_image(img, imgPath, make_gray_color, "gray")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		process_image(img, imgPath, make_RBG_color, "rbg")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		process_image(img, imgPath, make_GBR_color, "gbr")
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		process_image(img, imgPath, make_GRB_color, "grb")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		process_image(img, imgPath, make_BRG_color, "brg")
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		process_image(img, imgPath, make_BGR_color, "bgr")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		process_image(img, imgPath, make_GB_color, "gb")
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		process_image(img, imgPath, make_RB_color, "rb")
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		process_image(img, imgPath, make_RG_color, "rg")
	}()

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
