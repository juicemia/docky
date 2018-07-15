package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/davecgh/go-spew/spew"

	"html/template"

	"github.com/ghodss/yaml"
)

var jsonSchema = `{
    "type": "object",
    "properties": {
        "status": {
            "type": "string",
            "enum": ["UP", "DOWN", "UNKNOWN"]
        }
    },
    "additionalProperties": false,
    "required": ["status"]
}`

const outFlags = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
const outMode = 0644

func main() {
	rootTpl, err := os.Open("./templates/root.tpl.html")
	if err != nil {
		panic(err)
	}

	buf, err := ioutil.ReadAll(rootTpl)
	if err != nil {
		panic(err)
	}

	t, err := template.New("root").Parse(string(buf))
	if err != nil {
		panic(err)
	}

	appCfg := loadConfig()
	for _, r := range appCfg.Resources {
		filename := fmt.Sprintf("%v.html", r.Name)
		out, err := os.OpenFile(filename, outFlags, outMode)
		if err != nil {
			panic(err)
		}

		cfg := struct {
			AppName         string `yaml:"app_name"`
			Resources       []resource
			CurrentResource resource
		}{
			AppName:         appCfg.AppName,
			Resources:       appCfg.Resources,
			CurrentResource: r,
		}

		err = t.Execute(out, cfg)
		if err != nil {
			panic(err)
		}
	}
}

type config struct {
	AppName         string `json:"app_name"`
	Resources       []resource
	CurrentResource resource
}

type resource struct {
	AppName string
	Name    string
	Routes  []route
}

type route struct {
	Method      string
	Path        string
	Description string
	Headers     map[string]string
	Parameters  map[string]interface{}
	Responses   []response
}

type header struct {
	Key    string
	Values []string
}

type parameter struct {
	Key  string
	Type string
}

type response struct {
	Status      int
	Description string
	Headers     map[string]string
	Body        map[string]interface{}
	BodySchema  string `yaml:"-"`
}

func loadConfig() config {
	var cfg config
	f, err := os.Open("docky.yaml")
	if err != nil {
		panic(err)
	}

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("buf: %s\n", buf)

	err = yaml.Unmarshal(buf, &cfg)
	if err != nil {
		panic(err)
	}

	for _, resource := range cfg.Resources {
		for _, route := range resource.Routes {
			for i := range route.Responses {
				resp := route.Responses[i]
				schema, err := json.MarshalIndent(resp.Body, "", "  ")
				if err != nil {
					panic(err)
				}

				resp.BodySchema = string(schema)

				spew.Dump(resp)
				route.Responses[i] = resp
			}
		}
	}

	spew.Dump(cfg)

	return cfg
}
