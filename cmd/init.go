package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/juicemia/docky/config"
	"github.com/spf13/cobra"
)

func newInitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Create a docky.yaml file in the current directory",
		Long:  `Create a docky.yaml file in the current directory`,
		Run: func(cmd *cobra.Command, args []string) {
			if _, err := os.Stat("docky.yaml"); err != nil {
				dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				segs := strings.Split(dir, "/")
				app := config.App{
					Name:         segs[len(segs)-1],
					OutputFolder: "public",
				}

				buf, err := app.GetYAMLDefinition()
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				err = ioutil.WriteFile("docky.yaml", buf, 0644)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				os.Exit(0)
			}

			fmt.Println("docky.yaml already exists")
			os.Exit(1)
		},
	}
}
