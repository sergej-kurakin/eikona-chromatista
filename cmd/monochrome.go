package cmd

import (
	"image/jpeg"
	"os"
	"sync"

	"github.com/sergej-kurakin/eikona-chromatista/processor"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(grayscaleCmd)
	rootCmd.AddCommand(redscaleCmd)
	rootCmd.AddCommand(greenscaleCmd)
	rootCmd.AddCommand(bluescaleCmd)
	rootCmd.AddCommand(monochromeAllCmd)
}

var grayscaleCmd = &cobra.Command{
	Use:   "grayscale",
	Short: "Swap colors from RGB to grayscale",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.PhotometricGrayscaleColorProcessor, "grayscale")
	},
}

var redscaleCmd = &cobra.Command{
	Use:   "redscale",
	Short: "Swap colors from RGB to redscale",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.PhotometricRedscaleColorProcessor, "redscale")
	},
}

var greenscaleCmd = &cobra.Command{
	Use:   "greenscale",
	Short: "Swap colors from RGB to greenscale",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.PhotometricGreenscaleColorProcessor, "greenscale")
	},
}

var bluescaleCmd = &cobra.Command{
	Use:   "bluescale",
	Short: "Swap colors from RGB to bluescale",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.PhotometricBluescaleColorProcessor, "bluescale")
	},
}

var redCmd = &cobra.Command{
	Use:   "red",
	Short: "Limit colors from RGB to reds",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.RXXColorProcessor, "red")
	},
}

var greenCmd = &cobra.Command{
	Use:   "green",
	Short: "Limit colors from RGB to green",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.XGXColorProcessor, "green")
	},
}

var blueCmd = &cobra.Command{
	Use:   "blue",
	Short: "Limit colors from RGB to blue",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.XXBColorProcessor, "blue")
	},
}

var sepiaCmd = &cobra.Command{
	Use:   "sepia",
	Short: "Swap colors from RGB to sepia",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.SepiaColorProcessor, "sepia")
	},
}

var monochromeAllCmd = &cobra.Command{
	Use:   "monochromeAll",
	Short: "Swap colors from RGB to different monochrome combinations",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		var processors [8]Processor
		processors[0] = Processor{suffix: "grayscale", color_processor: processor.PhotometricGrayscaleColorProcessor}
		processors[1] = Processor{suffix: "redscale", color_processor: processor.PhotometricRedscaleColorProcessor}
		processors[2] = Processor{suffix: "greenscale", color_processor: processor.PhotometricGreenscaleColorProcessor}
		processors[3] = Processor{suffix: "bluescale", color_processor: processor.PhotometricBluescaleColorProcessor}
		processors[4] = Processor{suffix: "red", color_processor: processor.RXXColorProcessor}
		processors[5] = Processor{suffix: "green", color_processor: processor.XGXColorProcessor}
		processors[6] = Processor{suffix: "blue", color_processor: processor.XXBColorProcessor}
		processors[7] = Processor{suffix: "sepia", color_processor: processor.SepiaColorProcessor}

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
