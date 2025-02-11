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

type Processor struct {
	suffix         string
	colorProcessor processor.ColorProcessor
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func process(imgPath string) {
	var processors [36]Processor
	processors[0] = Processor{suffix: "gray", colorProcessor: processor.GrayColorProcessor}
	processors[1] = Processor{suffix: "rbg", colorProcessor: processor.RBGColorProcessor}
	processors[2] = Processor{suffix: "gbr", colorProcessor: processor.GBRColorProcessor}
	processors[3] = Processor{suffix: "grb", colorProcessor: processor.GRBColorProcessor}
	processors[4] = Processor{suffix: "brg", colorProcessor: processor.BRGColorProcessor}
	processors[5] = Processor{suffix: "bgr", colorProcessor: processor.BGRColorProcessor}
	processors[6] = Processor{suffix: "red", colorProcessor: processor.RXXColorProcessor}
	processors[7] = Processor{suffix: "green", colorProcessor: processor.XGXColorProcessor}
	processors[8] = Processor{suffix: "blue", colorProcessor: processor.XXBColorProcessor}
	processors[9] = Processor{suffix: "rnd", colorProcessor: processor.RandomColorProcessor}
	processors[10] = Processor{suffix: "photometric_grayscale", colorProcessor: processor.PhotometricGrayscaleColorProcessor}
	processors[11] = Processor{suffix: "photometric_graychrome", colorProcessor: processor.PhotometricGraychromeColorProcessor}
	processors[12] = Processor{suffix: "photometric_graychrome_negative", colorProcessor: processor.PhotometricGraychromeNegativeColorProcessor}
	processors[13] = Processor{suffix: "negative", colorProcessor: processor.NegativeColorProcessor}
	processors[14] = Processor{suffix: "redscale", colorProcessor: processor.PhotometricRedscaleColorProcessor}
	processors[15] = Processor{suffix: "greenscale", colorProcessor: processor.PhotometricGreenscaleColorProcessor}
	processors[16] = Processor{suffix: "bluescale", colorProcessor: processor.PhotometricBluescaleColorProcessor}
	processors[17] = Processor{suffix: "sepia", colorProcessor: processor.SepiaColorProcessor}

	// dualchrome
	processors[18] = Processor{suffix: "xgb", colorProcessor: processor.XGBColorProcessor}
	processors[19] = Processor{suffix: "rxb", colorProcessor: processor.RXBColorProcessor}
	processors[20] = Processor{suffix: "rgx", colorProcessor: processor.RGXColorProcessor}
	processors[21] = Processor{suffix: "rbx", colorProcessor: processor.RBXColorProcessor}
	processors[22] = Processor{suffix: "grx", colorProcessor: processor.GRXColorProcessor}
	processors[23] = Processor{suffix: "gbx", colorProcessor: processor.GBXColorProcessor}
	processors[24] = Processor{suffix: "brx", colorProcessor: processor.BRXColorProcessor}
	processors[25] = Processor{suffix: "bgx", colorProcessor: processor.BGXColorProcessor}
	processors[26] = Processor{suffix: "rxg", colorProcessor: processor.RXGColorProcessor}
	processors[27] = Processor{suffix: "gxr", colorProcessor: processor.GXRColorProcessor}
	processors[28] = Processor{suffix: "gxb", colorProcessor: processor.GXBColorProcessor}
	processors[29] = Processor{suffix: "bxr", colorProcessor: processor.BXRColorProcessor}
	processors[30] = Processor{suffix: "bxg", colorProcessor: processor.BXGColorProcessor}
	processors[31] = Processor{suffix: "xrg", colorProcessor: processor.XRGColorProcessor}
	processors[32] = Processor{suffix: "xrb", colorProcessor: processor.XRBColorProcessor}
	processors[33] = Processor{suffix: "xgr", colorProcessor: processor.XGRColorProcessor}
	processors[34] = Processor{suffix: "xbr", colorProcessor: processor.XBRColorProcessor}
	processors[35] = Processor{suffix: "xbg", colorProcessor: processor.XBGColorProcessor}

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
			processImage(img, imgPath, processors[k].colorProcessor, processors[k].suffix)
		}()
	}

	wg.Wait()
	fmt.Printf("Processing finished\n")
}

func processImage(img image.Image, imgPath string, colorProcessor processor.ColorProcessor, resultSuffix string) {
	wImg := processPixels(img, colorProcessor)
	newImgPath := newFileName(imgPath, resultSuffix)
	export(newImgPath, wImg)
}

func processPixels(img image.Image, colorProcessor processor.ColorProcessor) image.Image {
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

			c, err := colorProcessor(originalColor)
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
	fg, err := os.Create(newImagePath)
	check(err)
	defer fg.Close()
	err = jpeg.Encode(fg, wImg, &jpeg.Options{Quality: 100})
	check(err)
}
