package cmd

import (
	"image/jpeg"
	"os"
	"sync"

	"github.com/sergej-kurakin/eikona-chromatista/processor"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(rgxCmd)
	rootCmd.AddCommand(xgbCmd)
	rootCmd.AddCommand(rxbCmd)

	rootCmd.AddCommand(rbxCmd)
	rootCmd.AddCommand(grxCmd)
	rootCmd.AddCommand(gbxCmd)
	rootCmd.AddCommand(brxCmd)
	rootCmd.AddCommand(bgxCmd)

	rootCmd.AddCommand(rxgCmd)
	rootCmd.AddCommand(gxrCmd)
	rootCmd.AddCommand(gxbCmd)
	rootCmd.AddCommand(bxrCmd)
	rootCmd.AddCommand(bxgCmd)

	rootCmd.AddCommand(xrgCmd)
	rootCmd.AddCommand(xrbCmd)
	rootCmd.AddCommand(xgrCmd)
	rootCmd.AddCommand(xbrCmd)
	rootCmd.AddCommand(xbgCmd)

	rootCmd.AddCommand(dualchromeAllCmd)
}

var rgxCmd = &cobra.Command{
	Use:   "rgx",
	Short: "Limit colors from RGB to RGX",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.NewRGXProcessor())
	},
}

var xgbCmd = &cobra.Command{
	Use:   "xgb",
	Short: "Limit colors from RGB to XGB",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.NewXGBProcessor())
	},
}

var rxbCmd = &cobra.Command{
	Use:   "rxb",
	Short: "Limit colors from RGB to RXB",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.NewRXBProcessor())
	},
}

var rbxCmd = &cobra.Command{
	Use:   "rbx",
	Short: "Limit & swap colors from RGB to RBX",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.NewRBXProcessor())
	},
}

var grxCmd = &cobra.Command{
	Use:   "grx",
	Short: "Limit & swap colors from RGB to GRX",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.NewGRXProcessor())
	},
}

var gbxCmd = &cobra.Command{
	Use:   "gbx",
	Short: "Limit & swap colors from RGB to GBX",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.NewGBXProcessor())
	},
}

var brxCmd = &cobra.Command{
	Use:   "rb",
	Short: "Limit & swap colors from RGB to BRX",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.NewBRXProcessor())
	},
}

var bgxCmd = &cobra.Command{
	Use:   "bgx",
	Short: "Limit & swap colors from RGB to BGX",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.NewBGXProcessor())
	},
}

var rxgCmd = &cobra.Command{
	Use:   "rxg",
	Short: "Limit & swap colors from RGB to RXG",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.NewRXGProcessor())
	},
}

var gxrCmd = &cobra.Command{
	Use:   "gxr",
	Short: "Limit & swap colors from RGB to GXR",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.NewGXRProcessor())
	},
}

var gxbCmd = &cobra.Command{
	Use:   "gxb",
	Short: "Limit & swap colors from RGB to GXB",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.NewGXBProcessor())
	},
}

var bxrCmd = &cobra.Command{
	Use:   "bxr",
	Short: "Limit & swap colors from RGB to BXR",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.NewBXRProcessor())
	},
}

var bxgCmd = &cobra.Command{
	Use:   "bxg",
	Short: "Limit & swap colors from RGB to BXG",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.NewBXGProcessor())
	},
}

var xrgCmd = &cobra.Command{
	Use:   "xrg",
	Short: "Limit & swap colors from RGB to XRG",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.NewXRGProcessor())
	},
}

var xrbCmd = &cobra.Command{
	Use:   "xrb",
	Short: "Limit & swap colors from RGB to XRB",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.NewXRBProcessor())
	},
}

var xgrCmd = &cobra.Command{
	Use:   "xgr",
	Short: "Limit & swap colors from RGB to XGR",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.NewXGRProcessor())
	},
}

var xbrCmd = &cobra.Command{
	Use:   "xbr",
	Short: "Limit & swap colors from RGB to XBR",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.NewXBRProcessor())
	},
}

var xbgCmd = &cobra.Command{
	Use:   "xbg",
	Short: "Limit & swap colors from RGB to XBG",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		rgbProcess(args[0], processor.NewXBGProcessor())
	},
}

var dualchromeAllCmd = &cobra.Command{
	Use:   "dualchromeAll",
	Short: "Swap colors from RGB to different dualchrome combinations",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		processors := []processor.Processor{
			processor.NewRGXProcessor(),
			processor.NewXGBProcessor(),
			processor.NewRXBProcessor(),
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
