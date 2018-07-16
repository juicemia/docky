package main

import (
	"fmt"
	"html/template"
	"os"
)

const outFlags = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
const outMode = 0744

var templates = []string{
	"./templates/root.tpl.html",
	"./templates/route.tpl.html",
}

func main() {
	app, err := LoadApp()
	if err != nil {
		panic(err)
	}

	if app.OutputFolder == "" {
		app.OutputFolder = "public"
	}

	if _, err := os.Stat(app.OutputFolder); os.IsNotExist(err) {
		if err := os.MkdirAll(app.OutputFolder, os.ModeDir|0755); err != nil {
			panic(err)
		}
	}

	t, err := template.New("out").Parse(root + route)
	if err != nil {
		panic(err)
	}

	for _, r := range app.Resources {
		filename := fmt.Sprintf("%v/%v.html", app.OutputFolder, r.Name)
		out, err := os.OpenFile(filename, outFlags, outMode)
		if err != nil {
			panic(err)
		}

		pg := Page{AppName: app.Name, Resources: app.Resources, Resource: r}

		err = t.Execute(out, pg)
		if err != nil {
			panic(err)
		}
	}
}
