package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/alecthomas/template"
	"github.com/juicemia/docky/config"
	"github.com/spf13/cobra"
)

func newGenerateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "generate",
		Short: "Generate API documentation",
		Long: `Generates a static site for API documentation using "docky.yaml" for
	the application definition.`,
		Run: func(cmd *cobra.Command, args []string) {
			app, err := config.LoadApp()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			err = runGenerate(app)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
}

func runGenerate(app config.App) error {
	outFlags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	outMode := os.FileMode(0744)

	if app.OutputFolder == "" {
		app.OutputFolder = "public"
	}

	if _, err := os.Stat(app.OutputFolder); os.IsNotExist(err) {
		if err := os.MkdirAll(app.OutputFolder, os.ModeDir|0755); err != nil {
			return err
		}
	}

	itpl, err := template.New("index").Parse(root + index)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%v/index.html", app.OutputFolder)
	out, err := os.OpenFile(filename, outFlags, outMode)
	if err != nil {
		return err
	}

	err = itpl.Execute(out, config.Page{AppName: app.Name, Resources: app.Resources})
	if err != nil {
		return err
	}

	t, err := template.New("route").Parse(root + route)
	if err != nil {
		return err
	}

	for _, r := range app.Resources {
		filename := fmt.Sprintf("%v/%v.html", app.OutputFolder, strings.ToLower(r.Name))
		out, err := os.OpenFile(filename, outFlags, outMode)
		if err != nil {
			return err
		}

		pg := config.Page{AppName: app.Name, Resources: app.Resources, Resource: r}

		err = t.Execute(out, pg)
		if err != nil {
			return err
		}
	}

	return nil
}
