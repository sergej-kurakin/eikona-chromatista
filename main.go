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
		gray(img, imgPath)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		rbg(img, imgPath)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		gbr(img, imgPath)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		grb(img, imgPath)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		brg(img, imgPath)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		bgr(img, imgPath)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		gb(img, imgPath)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		rb(img, imgPath)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		rg(img, imgPath)
	}()

	wg.Wait()
	fmt.Printf("Processing finished\n")
}

func gbr(img image.Image, imgPath string) {
	wImg := process_pixels(img, make_GBR_color)
	newImgPath := newFileName(imgPath, "gbr")
	export(newImgPath, wImg)
}

func grb(img image.Image, imgPath string) {
	wImg := process_pixels(img, make_GRB_color)
	newImgPath := newFileName(imgPath, "grb")
	export(newImgPath, wImg)
}

func brg(img image.Image, imgPath string) {
	wImg := process_pixels(img, make_BRG_color)
	newImgPath := newFileName(imgPath, "brg")
	export(newImgPath, wImg)
}

func bgr(img image.Image, imgPath string) {
	wImg := process_pixels(img, make_BGR_color)
	newImgPath := newFileName(imgPath, "bgr")
	export(newImgPath, wImg)
}

func rbg(img image.Image, imgPath string) {
	wImg := process_pixels(img, make_RBG_color)
	newImgPath := newFileName(imgPath, "rbg")
	export(newImgPath, wImg)
}

func gb(img image.Image, imgPath string) {
	wImg := process_pixels(img, make_GB_color)
	newImgPath := newFileName(imgPath, "gb")
	export(newImgPath, wImg)
}

func rb(img image.Image, imgPath string) {
	wImg := process_pixels(img, make_RB_color)
	newImgPath := newFileName(imgPath, "rb")
	export(newImgPath, wImg)
}

func rg(img image.Image, imgPath string) {
	wImg := process_pixels(img, make_RG_color)
	newImgPath := newFileName(imgPath, "rg")
	export(newImgPath, wImg)
}

func gray(img image.Image, imgPath string) {
	wImg := process_pixels(img, make_gray_color)
	newImgPath := newFileName(imgPath, "gray")
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
