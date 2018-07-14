package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"html/template"
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

	rootOut, err := os.OpenFile("root.html", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}

	health := resource{
		Name: "Health",
		Routes: []route{
			route{
				Method:      http.MethodGet,
				Path:        "/health",
				Description: "Returns the health of the API.",
				Headers: []header{
					header{
						Key:    "Accept",
						Values: "application/json",
					},
				},
				Responses: []response{
					response{
						Status: status{Code: 200, Description: "Success"},
						Headers: []header{
							header{
								Key:    "Content-Type",
								Values: "application/json",
							},
						},
						Body: jsonSchema,
					},
				},
			},
		},
	}

	cfg := data{
		AppName: "CoolApp",
		Resources: []resource{
			health,
		},
		CurrentResource: health,
	}

	err = t.Execute(rootOut, cfg)
	if err != nil {
		panic(err)
	}
}

type data struct {
	AppName         string
	Resources       []resource
	CurrentResource resource
}

type resource struct {
	Name   string
	Routes []route
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
	Values string
}

type parameter struct {
	Key  string
	Type string
}

type response struct {
	Status  status
	Headers []header
	Body    string
}

type status struct {
	Code        int
	Description string
}
