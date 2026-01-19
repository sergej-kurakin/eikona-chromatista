package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Eikona Chromatista",
	Long:  `All software has versions. This is Eikona Chromatista's`,
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Eikona Chromatista (Εικόνα Χρωματιστά) Image Color Manipulation v0.0.0")
	},
}
