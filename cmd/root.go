package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "eiko",
	Short: "Eikona Chromatista is an Image Color manipulation software",
	Run: func(_ *cobra.Command, _ []string) {
		// Do Stuff Here
		fmt.Println("Hello, I'm Eikona Chromatista (Εικόνα Χρωματιστά)!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
