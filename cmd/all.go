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
	color_processor processor.ProcessorFunc
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func process(imgPath string) {

	var processors [36]Processor
	processors[0] = Processor{suffix: "gray", color_processor: processor.GrayColorProcessor}
	processors[1] = Processor{suffix: "rbg", color_processor: processor.RBGColorProcessor}
	processors[2] = Processor{suffix: "gbr", color_processor: processor.GBRColorProcessor}
	processors[3] = Processor{suffix: "grb", color_processor: processor.GRBColorProcessor}
	processors[4] = Processor{suffix: "brg", color_processor: processor.BRGColorProcessor}
	processors[5] = Processor{suffix: "bgr", color_processor: processor.BGRColorProcessor}
	processors[6] = Processor{suffix: "red", color_processor: processor.RXXColorProcessor}
	processors[7] = Processor{suffix: "green", color_processor: processor.XGXColorProcessor}
	processors[8] = Processor{suffix: "blue", color_processor: processor.XXBColorProcessor}
	processors[9] = Processor{suffix: "rnd", color_processor: processor.RandomColorProcessor}
	processors[10] = Processor{suffix: "photometric_grayscale", color_processor: processor.PhotometricGrayscaleColorProcessor}
	processors[11] = Processor{suffix: "photometric_graychrome", color_processor: processor.PhotometricGraychromeColorProcessor}
	processors[12] = Processor{suffix: "photometric_graychrome_negative", color_processor: processor.PhotometricGraychromeNegativeColorProcessor}
	processors[13] = Processor{suffix: "negative", color_processor: processor.NegativeColorProcessor}
	processors[14] = Processor{suffix: "redscale", color_processor: processor.PhotometricRedscaleColorProcessor}
	processors[15] = Processor{suffix: "greenscale", color_processor: processor.PhotometricGreenscaleColorProcessor}
	processors[16] = Processor{suffix: "bluescale", color_processor: processor.PhotometricBluescaleColorProcessor}
	processors[17] = Processor{suffix: "sepia", color_processor: processor.SepiaColorProcessor}

	// dualchrome
	processors[18] = Processor{suffix: "xgb", color_processor: processor.XGBColorProcessor}
	processors[19] = Processor{suffix: "rxb", color_processor: processor.RXBColorProcessor}
	processors[20] = Processor{suffix: "rgx", color_processor: processor.RGXColorProcessor}
	processors[21] = Processor{suffix: "rbx", color_processor: processor.RBXColorProcessor}
	processors[22] = Processor{suffix: "grx", color_processor: processor.GRXColorProcessor}
	processors[23] = Processor{suffix: "gbx", color_processor: processor.GBXColorProcessor}
	processors[24] = Processor{suffix: "brx", color_processor: processor.BRXColorProcessor}
	processors[25] = Processor{suffix: "bgx", color_processor: processor.BGXColorProcessor}
	processors[26] = Processor{suffix: "rxg", color_processor: processor.RXGColorProcessor}
	processors[27] = Processor{suffix: "gxr", color_processor: processor.GXRColorProcessor}
	processors[28] = Processor{suffix: "gxb", color_processor: processor.GXBColorProcessor}
	processors[29] = Processor{suffix: "bxr", color_processor: processor.BXRColorProcessor}
	processors[30] = Processor{suffix: "bxg", color_processor: processor.BXGColorProcessor}
	processors[31] = Processor{suffix: "xrg", color_processor: processor.XRGColorProcessor}
	processors[32] = Processor{suffix: "xrb", color_processor: processor.XRBColorProcessor}
	processors[33] = Processor{suffix: "xgr", color_processor: processor.XGRColorProcessor}
	processors[34] = Processor{suffix: "xbr", color_processor: processor.XBRColorProcessor}
	processors[35] = Processor{suffix: "xbg", color_processor: processor.XBGColorProcessor}

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
