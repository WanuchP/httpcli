package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "httpcli",
	Short: "Application for sending HTTP requests (DELETE, GET, POST, PUT) to a given URL with flags, headers, and query parameters.",
}
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Print help message")
	rootCmd.PersistentFlags().StringSliceP("query", "q", nil, "Query parameters")
	rootCmd.PersistentFlags().StringSliceP("header", "H", nil, "Headers")
}


