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
	rootCmd.AddCommand(redCmd)
	rootCmd.AddCommand(greenCmd)
	rootCmd.AddCommand(blueCmd)
	rootCmd.AddCommand(sepiaCmd)
}

var grayscaleCmd = &cobra.Command{
	Use:   "grayscale",
	Short: "Swap colors from RGB to grayscale",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.NewPhotometricGrayscaleProcessor(nil))
	},
}

var redscaleCmd = &cobra.Command{
	Use:   "redscale",
	Short: "Swap colors from RGB to redscale",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.NewPhotometricRedscaleProcessor(nil))
	},
}

var greenscaleCmd = &cobra.Command{
	Use:   "greenscale",
	Short: "Swap colors from RGB to greenscale",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.NewPhotometricGreenscaleProcessor(nil))
	},
}

var bluescaleCmd = &cobra.Command{
	Use:   "bluescale",
	Short: "Swap colors from RGB to bluescale",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.NewPhotometricBluescaleProcessor(nil))
	},
}

var redCmd = &cobra.Command{
	Use:   "red",
	Short: "Limit colors from RGB to reds",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.NewRXXProcessor())
	},
}

var greenCmd = &cobra.Command{
	Use:   "green",
	Short: "Limit colors from RGB to green",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.NewXGXProcessor())
	},
}

var blueCmd = &cobra.Command{
	Use:   "blue",
	Short: "Limit colors from RGB to blue",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.NewXXBProcessor())
	},
}

var sepiaCmd = &cobra.Command{
	Use:   "sepia",
	Short: "Swap colors from RGB to sepia",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.NewSepiaProcessor())
	},
}

var monochromeAllCmd = &cobra.Command{
	Use:   "monochromeAll",
	Short: "Swap colors from RGB to different monochrome combinations",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		processors := []processor.Processor{
			processor.NewPhotometricGrayscaleProcessor(nil),
			processor.NewPhotometricRedscaleProcessor(nil),
			processor.NewPhotometricGreenscaleProcessor(nil),
			processor.NewPhotometricBluescaleProcessor(nil),
			processor.NewRXXProcessor(),
			processor.NewXGXProcessor(),
			processor.NewXXBProcessor(),
			processor.NewSepiaProcessor(),
		}

		f, err := os.Open(args[0])
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
				processImage(img, args[0], processors[k])
			}()
		}

		wg.Wait()
	},
}
