package cmd

import (
	"image/jpeg"
	"os"
	"sync"

	"github.com/sergej-kurakin/eikona-chromatista/processor"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(rgCmd)
	rootCmd.AddCommand(gbCmd)
	rootCmd.AddCommand(rbCmd)
	rootCmd.AddCommand(dualchromeAllCmd)
}

var rgCmd = &cobra.Command{
	Use:   "rg",
	Short: "Limit colors from RGB to RG",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.RGColorProcessor, "rg")
	},
}

var gbCmd = &cobra.Command{
	Use:   "gb",
	Short: "Limit colors from RGB to GB",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.GBColorProcessor, "gb")
	},
}

var rbCmd = &cobra.Command{
	Use:   "rb",
	Short: "Limit colors from RGB to RB",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rgbProcess(args[0], processor.RBColorProcessor, "rb")
	},
}

var dualchromeAllCmd = &cobra.Command{
	Use:   "dualchromeAll",
	Short: "Swap colors from RGB to different dualchrome combinations",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		var processors [3]Processor
		processors[0] = Processor{suffix: "rg", color_processor: processor.RGColorProcessor}
		processors[1] = Processor{suffix: "gb", color_processor: processor.GBColorProcessor}
		processors[2] = Processor{suffix: "rb", color_processor: processor.RBColorProcessor}

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
