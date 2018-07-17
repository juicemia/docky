package cmd

import (
	"fmt"
	"os"

	"github.com/juicemia/docky/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "docky",
	Short: "Docky is a static site generator geared towards documentation",
	Long: `A static site generator geared towards API documentation, with
support for validation.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	app, err := config.LoadApp()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rootCmd.AddCommand(newGenerateCmd(app))
}
