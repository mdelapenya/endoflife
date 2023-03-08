package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "eol",
	Short: "eol is a CLI to query the endoflife.date API",
	Long: `EndOfLife (eol) is a CLI to query the endoflife.date API built with
				  love by mdelapenya and friends in Go`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
