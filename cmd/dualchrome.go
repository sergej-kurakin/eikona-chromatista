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
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.RGXColorProcessor, "rgx")
	},
}

var xgbCmd = &cobra.Command{
	Use:   "xgb",
	Short: "Limit colors from RGB to XGB",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.XGBColorProcessor, "xgb")
	},
}

var rxbCmd = &cobra.Command{
	Use:   "rxb",
	Short: "Limit colors from RGB to RXB",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.RXBColorProcessor, "rxb")
	},
}

var rbxCmd = &cobra.Command{
	Use:   "rbx",
	Short: "Limit & swap colors from RGB to RBX",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.RBXColorProcessor, "rbx")
	},
}

var grxCmd = &cobra.Command{
	Use:   "grx",
	Short: "Limit & swap colors from RGB to GRX",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.GRXColorProcessor, "grx")
	},
}

var gbxCmd = &cobra.Command{
	Use:   "gbx",
	Short: "Limit & swap colors from RGB to GBX",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.GBXColorProcessor, "gbx")
	},
}

var brxCmd = &cobra.Command{
	Use:   "rb",
	Short: "Limit & swap colors from RGB to BRX",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.BRXColorProcessor, "brx")
	},
}

var bgxCmd = &cobra.Command{
	Use:   "bgx",
	Short: "Limit & swap colors from RGB to BGX",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.BGXColorProcessor, "bgx")
	},
}

var rxgCmd = &cobra.Command{
	Use:   "rxg",
	Short: "Limit & swap colors from RGB to RXG",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.RXGColorProcessor, "rxg")
	},
}

var gxrCmd = &cobra.Command{
	Use:   "gxr",
	Short: "Limit & swap colors from RGB to GXR",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.GXRColorProcessor, "gxr")
	},
}

var gxbCmd = &cobra.Command{
	Use:   "gxb",
	Short: "Limit & swap colors from RGB to GXB",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.GXBColorProcessor, "gxb")
	},
}

var bxrCmd = &cobra.Command{
	Use:   "bxr",
	Short: "Limit & swap colors from RGB to BXR",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.BXRColorProcessor, "bxr")
	},
}

var bxgCmd = &cobra.Command{
	Use:   "bxg",
	Short: "Limit & swap colors from RGB to BXG",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.BXGColorProcessor, "bxg")
	},
}

var xrgCmd = &cobra.Command{
	Use:   "xrg",
	Short: "Limit & swap colors from RGB to XRG",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.XRGColorProcessor, "xrg")
	},
}

var xrbCmd = &cobra.Command{
	Use:   "xrb",
	Short: "Limit & swap colors from RGB to XRB",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.XRBColorProcessor, "xrb")
	},
}

var xgrCmd = &cobra.Command{
	Use:   "xgr",
	Short: "Limit & swap colors from RGB to XGR",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.XGRColorProcessor, "xgr")
	},
}

var xbrCmd = &cobra.Command{
	Use:   "xbr",
	Short: "Limit & swap colors from RGB to XBR",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.XBRColorProcessor, "xbr")
	},
}

var xbgCmd = &cobra.Command{
	Use:   "xbg",
	Short: "Limit & swap colors from RGB to XBG",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.XBGColorProcessor, "xbg")
	},
}

var dualchromeAllCmd = &cobra.Command{
	Use:   "dualchromeAll",
	Short: "Swap colors from RGB to different dualchrome combinations",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var processors [17]Processor
		processors[0] = Processor{suffix: "rgx", color_processor: processor.RGXColorProcessor}
		processors[1] = Processor{suffix: "xgb", color_processor: processor.XGBColorProcessor}
		processors[2] = Processor{suffix: "rxb", color_processor: processor.RXBColorProcessor}

		processors[3] = Processor{suffix: "rbx", color_processor: processor.RBXColorProcessor}
		processors[4] = Processor{suffix: "grx", color_processor: processor.GRXColorProcessor}
		processors[5] = Processor{suffix: "gbx", color_processor: processor.GBXColorProcessor}
		processors[6] = Processor{suffix: "brx", color_processor: processor.BRXColorProcessor}
		processors[7] = Processor{suffix: "bgx", color_processor: processor.BGXColorProcessor}

		processors[8] = Processor{suffix: "rxg", color_processor: processor.RXGColorProcessor}
		processors[9] = Processor{suffix: "gxr", color_processor: processor.GXRColorProcessor}
		processors[10] = Processor{suffix: "gxb", color_processor: processor.GXBColorProcessor}
		processors[11] = Processor{suffix: "bxr", color_processor: processor.BXRColorProcessor}
		processors[11] = Processor{suffix: "bxg", color_processor: processor.BXGColorProcessor}

		processors[12] = Processor{suffix: "xrg", color_processor: processor.XRGColorProcessor}
		processors[13] = Processor{suffix: "xrb", color_processor: processor.XRBColorProcessor}
		processors[14] = Processor{suffix: "xgr", color_processor: processor.XGRColorProcessor}
		processors[15] = Processor{suffix: "xbr", color_processor: processor.XBRColorProcessor}
		processors[16] = Processor{suffix: "xbg", color_processor: processor.XBGColorProcessor}

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
