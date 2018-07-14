package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"html/template"

	"github.com/BurntSushi/toml"
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

	// health := resource{
	// 	Name: "Health",
	// 	Routes: []route{
	// 		route{
	// 			Method:      http.MethodGet,
	// 			Path:        "/health",
	// 			Description: "Returns the health of the API.",
	// 			Headers: []header{
	// 				header{
	// 					Key:    "Accept",
	// 					Values: "application/json",
	// 				},
	// 			},
	// 			Responses: []response{
	// 				response{
	// 					Status: status{Code: 200, Description: "Success"},
	// 					Headers: []header{
	// 						header{
	// 							Key:    "Content-Type",
	// 							Values: "application/json",
	// 						},
	// 					},
	// 					Body: jsonSchema,
	// 				},
	// 			},
	// 		},
	// 	},
	// }

	// cfg := config{
	// 	AppName: "CoolApp",
	// 	Resources: []resource{
	// 		health,
	// 	},
	// 	CurrentResource: health,
	// }

	appCfg := loadConfig()
	for _, r := range appCfg.Resources {
		filename := fmt.Sprintf("%v.html", r.Name)
		out, err := os.OpenFile(filename, outFlags, outMode)
		if err != nil {
			panic(err)
		}

		cfg := struct {
			AppName         string
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
	AppName         string
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
	Headers     []header
	Parameters  []parameter
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
	Headers     []header
	Body        string
}

func loadConfig() config {
	var cfg config
	_, err := toml.DecodeFile("docky.toml", &cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}
