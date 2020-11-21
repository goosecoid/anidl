package cmd

import (
	"fmt"
	"os"
	
	"github.com/spf13/cobra"

)

var rootCmd = &cobra.Command{
	Use:   "anidl",
	Short: "Anime website scraping tool",
}

func init () {
	rootCmd.SetHelpCommand(&cobra.Command{
		Use: "no-help",
		Hidden: true,
	})
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
