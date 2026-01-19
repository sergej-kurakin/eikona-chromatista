package cmd

import (
	"image/jpeg"
	"os"
	"path/filepath"
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
		rgbProcess(args[0], processor.NewGBRProcessor())
	},
}

var grbCmd = &cobra.Command{
	Use:   "grb",
	Short: "Swap colors from RGB to GRB",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.NewGRBProcessor())
	},
}

var brgCmd = &cobra.Command{
	Use:   "brg",
	Short: "Swap colors from RGB to BRG",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.NewBRGProcessor())
	},
}

var bgrCmd = &cobra.Command{
	Use:   "bgr",
	Short: "Swap colors from RGB to BGR",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.NewBGRProcessor())
	},
}

var rbgCmd = &cobra.Command{
	Use:   "rbg",
	Short: "Swap colors from RGB to RBG",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.NewRBGProcessor())
	},
}

var rgbAllCmd = &cobra.Command{
	Use:   "rgbAll",
	Short: "Swap colors from RGB to different combinations",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		processors := []processor.Processor{
			processor.NewRBGProcessor(),
			processor.NewGBRProcessor(),
			processor.NewGRBProcessor(),
			processor.NewBRGProcessor(),
			processor.NewBGRProcessor(),
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

func rgbProcess(imgPath string, proc processor.Processor) {
	f, err := os.Open(filepath.Clean(imgPath))
	check(err)
	defer f.Close() //nolint:errcheck // ignore error for defer close

	img, err := jpeg.Decode(f)

	check(err)

	processImage(img, imgPath, proc)
}
