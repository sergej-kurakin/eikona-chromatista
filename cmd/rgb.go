package cmd

import (
	"image/jpeg"
	"os"

	"github.com/sergej-kurakin/eikona-chromatista/processor"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(gbrCmd)
	rootCmd.AddCommand(grbCmd)
	rootCmd.AddCommand(brgCmd)
	rootCmd.AddCommand(bgrCmd)
	rootCmd.AddCommand(rbgCmd)
}

var gbrCmd = &cobra.Command{
	Use:   "gbr",
	Short: "Swap colors from RGB to GBR",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.GBRColorProcessor, "gbr")
	},
}

var grbCmd = &cobra.Command{
	Use:   "grb",
	Short: "Swap colors from RGB to GRB",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.GRBColorProcessor, "grb")
	},
}

var brgCmd = &cobra.Command{
	Use:   "brg",
	Short: "Swap colors from RGB to BRG",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.BRGColorProcessor, "brg")
	},
}

var bgrCmd = &cobra.Command{
	Use:   "bgr",
	Short: "Swap colors from RGB to BGR",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.BGRColorProcessor, "bgr")
	},
}

var rbgCmd = &cobra.Command{
	Use:   "rbg",
	Short: "Swap colors from RGB to RBG",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.RBGColorProcessor, "rbg")
	},
}

func rgbProcess(imgPath string, processorFunc processor.ProcessorFunc, suffix string) {
	f, err := os.Open(imgPath)
	check(err)
	defer f.Close()

	img, err := jpeg.Decode(f)

	check(err)

	process_image(img, imgPath, processorFunc, suffix)
}
