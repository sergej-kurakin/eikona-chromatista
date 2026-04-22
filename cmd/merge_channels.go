package cmd

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"

	"github.com/sergej-kurakin/eikona-chromatista/processor"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(mergeChannelsCmd)
}

var mergeChannelsCmd = &cobra.Command{
	Use:   "mergeChannels <red.jpg> <green.jpg> <blue.jpg>",
	Short: "Merge red, green, and blue channel images into one JPEG",
	Args:  cobra.ExactArgs(3),
	Run: func(_ *cobra.Command, args []string) {
		mergeChannels(args[0], args[1], args[2])
	},
}

func mergeChannels(redPath, greenPath, bluePath string) {
	redImg := loadJPEG(redPath)
	greenImg := loadJPEG(greenPath)
	blueImg := loadJPEG(bluePath)

	merger := processor.NewChannelMerger()
	merged, err := merger.Merge(redImg, greenImg, blueImg)
	check(err)

	newImgPath := newFileName(redPath, "merged_rgb")
	export(newImgPath, merged)
	fmt.Printf("Merged image created: %s\n", newImgPath)
}

func loadJPEG(path string) image.Image {
	f, err := os.Open(filepath.Clean(path))
	check(err)
	defer f.Close() //nolint:errcheck // ignore error for defer close

	img, err := jpeg.Decode(f)
	check(err)

	return img
}