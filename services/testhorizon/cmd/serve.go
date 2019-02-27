package cmd

import (
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "run testhorizon server",
	Long:  "serve initializes then starts the testhorizon HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		initApp().Serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
