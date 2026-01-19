package cmd

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/sergej-kurakin/eikona-chromatista/processor"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(allCmd)
}

var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Make all Image Color Manipulations",
	Long:  `Make all Image Color Manipulations, bla... bla... bla...`,
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		process(args[0])
	},
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func process(imgPath string) {
	processors := []processor.Processor{
		processor.NewGrayProcessor(),
		processor.NewRBGProcessor(),
		processor.NewGBRProcessor(),
		processor.NewGRBProcessor(),
		processor.NewBRGProcessor(),
		processor.NewBGRProcessor(),
		processor.NewRXXProcessor(),
		processor.NewXGXProcessor(),
		processor.NewXXBProcessor(),
		processor.NewRandomProcessor(nil),
		processor.NewPhotometricGrayscaleProcessor(nil),
		processor.NewPhotometricGraychromeProcessor(nil),
		processor.NewPhotometricGraychromeNegativeProcessor(nil, nil),
		processor.NewNegativeProcessor(),
		processor.NewPhotometricRedscaleProcessor(nil),
		processor.NewPhotometricGreenscaleProcessor(nil),
		processor.NewPhotometricBluescaleProcessor(nil),
		processor.NewSepiaProcessor(),
		// dualchrome
		processor.NewXGBProcessor(),
		processor.NewRXBProcessor(),
		processor.NewRGXProcessor(),
		processor.NewRBXProcessor(),
		processor.NewGRXProcessor(),
		processor.NewGBXProcessor(),
		processor.NewBRXProcessor(),
		processor.NewBGXProcessor(),
		processor.NewRXGProcessor(),
		processor.NewGXRProcessor(),
		processor.NewGXBProcessor(),
		processor.NewBXRProcessor(),
		processor.NewBXGProcessor(),
		processor.NewXRGProcessor(),
		processor.NewXRBProcessor(),
		processor.NewXGRProcessor(),
		processor.NewXBRProcessor(),
		processor.NewXBGProcessor(),
	}

	f, err := os.Open(filepath.Clean(imgPath))
	check(err)
	defer f.Close() //nolint:errcheck // ignore error for defer close

	img, err := jpeg.Decode(f)

	check(err)

	var wg sync.WaitGroup

	for i := 0; i < len(processors); i++ {
		wg.Add(1)
		k := i
		go func() {
			defer wg.Done()
			processImage(img, imgPath, processors[k])
		}()
	}

	wg.Wait()
	fmt.Printf("Processing finished\n")
}

func processImage(img image.Image, imgPath string, proc processor.Processor) {
	wImg := processPixels(img, proc)
	newImgPath := newFileName(imgPath, proc.Name())
	export(newImgPath, wImg)
}

func processPixels(img image.Image, proc processor.Processor) image.Image {
	size := img.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	wImg := image.NewRGBA(rect)

	for x := 0; x < size.X; x++ {
		// and now loop thorough all of this x's y
		for y := 0; y < size.Y; y++ {
			pixel := img.At(x, y)
			originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)

			if !ok {
				panic("Failed to convert color")
			}

			c, err := proc.Process(originalColor)
			check(err)

			wImg.Set(x, y, c)
		}
	}

	return wImg
}

func newFileName(imgPath, suffix string) string {
	ext := filepath.Ext(imgPath)
	name := strings.TrimSuffix(filepath.Base(imgPath), ext)
	newImagePath := fmt.Sprintf("%s/%s_%s%s", filepath.Dir(imgPath), name, suffix, ext)
	return newImagePath
}

func export(newImagePath string, wImg image.Image) {
	fg, err := os.Create(filepath.Clean(newImagePath))
	check(err)
	defer fg.Close() //nolint:errcheck // ignore error for defer close
	err = jpeg.Encode(fg, wImg, &jpeg.Options{Quality: 100})
	check(err)
}
