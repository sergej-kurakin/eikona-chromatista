package cmd

import (
	"image/jpeg"
	"os"
	"sync"

	"github.com/sergej-kurakin/eikona-chromatista/processor"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(gbrCmd)
	rootCmd.AddCommand(grbCmd)
	rootCmd.AddCommand(brgCmd)
	rootCmd.AddCommand(bgrCmd)
	rootCmd.AddCommand(rbgCmd)
	rootCmd.AddCommand(rgbAllCmd)
}

var gbrCmd = &cobra.Command{
	Use:   "gbr",
	Short: "Swap colors from RGB to GBR",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.GBRColorProcessor, "gbr")
	},
}

var grbCmd = &cobra.Command{
	Use:   "grb",
	Short: "Swap colors from RGB to GRB",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.GRBColorProcessor, "grb")
	},
}

var brgCmd = &cobra.Command{
	Use:   "brg",
	Short: "Swap colors from RGB to BRG",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.BRGColorProcessor, "brg")
	},
}

var bgrCmd = &cobra.Command{
	Use:   "bgr",
	Short: "Swap colors from RGB to BGR",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.BGRColorProcessor, "bgr")
	},
}

var rbgCmd = &cobra.Command{
	Use:   "rbg",
	Short: "Swap colors from RGB to RBG",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.RBGColorProcessor, "rbg")
	},
}

var rgbAllCmd = &cobra.Command{
	Use:   "rgbAll",
	Short: "Swap colors from RGB to different combinations",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		var processors [5]Processor
		processors[0] = Processor{suffix: "rbg", color_processor: processor.RBGColorProcessor}
		processors[1] = Processor{suffix: "gbr", color_processor: processor.GBRColorProcessor}
		processors[2] = Processor{suffix: "grb", color_processor: processor.GRBColorProcessor}
		processors[3] = Processor{suffix: "brg", color_processor: processor.BRGColorProcessor}
		processors[4] = Processor{suffix: "bgr", color_processor: processor.BGRColorProcessor}

		f, err := os.Open(args[0])
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
				process_image(img, args[0], processors[k].color_processor, processors[k].suffix)
			}()
		}

		wg.Wait()
	},
}

func rgbProcess(imgPath string, processorFunc processor.ColorProcessor, suffix string) {
	f, err := os.Open(imgPath)
	check(err)
	defer f.Close()

	img, err := jpeg.Decode(f)

	check(err)

	process_image(img, imgPath, processorFunc, suffix)
}
