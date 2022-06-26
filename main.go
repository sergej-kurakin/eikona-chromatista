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

	gray(img, imgPath)

	rbg(img, imgPath)

	gbr(img, imgPath)
	grb(img, imgPath)

	brg(img, imgPath)
	bgr(img, imgPath)
}

func gbr(img image.Image, imgPath string) {
	wImg := make_GBR(img)
	newImgPath := newFileName(imgPath, "gbr")
	export(newImgPath, wImg)
}

func grb(img image.Image, imgPath string) {
	wImg := make_GRB(img)
	newImgPath := newFileName(imgPath, "grb")
	export(newImgPath, wImg)
}

func brg(img image.Image, imgPath string) {
	wImg := make_BRG(img)
	newImgPath := newFileName(imgPath, "brg")
	export(newImgPath, wImg)
}

func bgr(img image.Image, imgPath string) {
	wImg := make_BGR(img)
	newImgPath := newFileName(imgPath, "bgr")
	export(newImgPath, wImg)
}

func rbg(img image.Image, imgPath string) {
	wImg := make_RBG(img)
	newImgPath := newFileName(imgPath, "rbg")
	export(newImgPath, wImg)
}

func gray(img image.Image, imgPath string) {
	wImg := make_gray(img)
	newImgPath := newFileName(imgPath, "gray")
	export(newImgPath, wImg)
}

func make_gray(img image.Image) image.Image {
	size := img.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	wImg := image.NewRGBA(rect)

	for x := 0; x < size.X; x++ {
		// and now loop thorough all of this x's y
		for y := 0; y < size.Y; y++ {
			pixel := img.At(x, y)
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
			wImg.Set(x, y, c)
		}
	}

	return wImg
}

func make_GBR(img image.Image) image.Image {
	size := img.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	wImg := image.NewRGBA(rect)

	for x := 0; x < size.X; x++ {
		// and now loop thorough all of this x's y
		for y := 0; y < size.Y; y++ {
			pixel := img.At(x, y)
			originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
			c := color.RGBA{
				R: originalColor.G, G: originalColor.B, B: originalColor.R, A: originalColor.A,
			}
			wImg.Set(x, y, c)
		}
	}

	return wImg
}

func make_GRB(img image.Image) image.Image {
	size := img.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	wImg := image.NewRGBA(rect)

	for x := 0; x < size.X; x++ {
		// and now loop thorough all of this x's y
		for y := 0; y < size.Y; y++ {
			pixel := img.At(x, y)
			originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
			c := color.RGBA{
				R: originalColor.G, G: originalColor.R, B: originalColor.B, A: originalColor.A,
			}
			wImg.Set(x, y, c)
		}
	}

	return wImg
}

func make_BRG(img image.Image) image.Image {
	size := img.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	wImg := image.NewRGBA(rect)

	for x := 0; x < size.X; x++ {
		// and now loop thorough all of this x's y
		for y := 0; y < size.Y; y++ {
			pixel := img.At(x, y)
			originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
			c := color.RGBA{
				R: originalColor.B, G: originalColor.R, B: originalColor.G, A: originalColor.A,
			}
			wImg.Set(x, y, c)
		}
	}

	return wImg
}

func make_BGR(img image.Image) image.Image {
	size := img.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	wImg := image.NewRGBA(rect)

	for x := 0; x < size.X; x++ {
		// and now loop thorough all of this x's y
		for y := 0; y < size.Y; y++ {
			pixel := img.At(x, y)
			originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
			c := color.RGBA{
				R: originalColor.B, G: originalColor.G, B: originalColor.R, A: originalColor.A,
			}
			wImg.Set(x, y, c)
		}
	}

	return wImg
}

func make_RBG(img image.Image) image.Image {
	size := img.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	wImg := image.NewRGBA(rect)

	for x := 0; x < size.X; x++ {
		// and now loop thorough all of this x's y
		for y := 0; y < size.Y; y++ {
			pixel := img.At(x, y)
			originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
			c := color.RGBA{
				R: originalColor.R, G: originalColor.B, B: originalColor.G, A: originalColor.A,
			}
			wImg.Set(x, y, c)
		}
	}

	return wImg
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
	err = jpeg.Encode(fg, wImg, &jpeg.Options{Quality: 90})
	check(err)
}
