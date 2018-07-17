package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"html/template"

	"github.com/ghodss/yaml"
)

// App is the top level definition of the application being documented.
type App struct {
	Name         string     `json:"app_name"`
	OutputFolder string     `json:"output_folder"`
	Resources    []Resource `json:"resources,omitempty"`
}

// Page is an individual page that's generated.
type Page struct {
	AppName   string
	Resources []Resource // need this here to generate all the links
	Resource  Resource
}

// Resource is a logical grouping of endpoints.
type Resource struct {
	Name        string
	Description string
	Routes      []Route
	Template    *template.Template
}

// Route is an individual endpoint that can be interacted with.
type Route struct {
	Method      string
	Path        string
	Description string
	Headers     map[string]string
	Parameters  map[string]interface{}
	Responses   []Response
}

// Response is one of the possible results from interacting with a Route.
type Response struct {
	Status      int
	Description string
	Headers     map[string]string
	// Body is the YAML definition of what the response body can contain.
	Body     map[string]interface{}
	BodyJSON string `yaml:"-"`
}

// LoadApp reads "docky.yaml" in the current working directory and creates an
// app definition from it.
func LoadApp() (App, error) {
	var cfg App
	f, err := os.Open("docky.yaml")
	if err != nil {
		return App{}, err
	}

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return App{}, err
	}

	err = yaml.Unmarshal(buf, &cfg)
	if err != nil {
		return App{}, err
	}

	for _, resource := range cfg.Resources {
		for _, route := range resource.Routes {
			for i := range route.Responses {
				resp := route.Responses[i]
				schema, err := json.MarshalIndent(resp.Body, "", "  ")
				if err != nil {
					return App{}, err
				}

				resp.BodyJSON = string(schema)

				route.Responses[i] = resp
			}
		}
	}

	return cfg, nil
}

func (r Resource) GetLinkName() string {
	return strings.ToLower(r.Name)
}

func (r Route) GetLinkName() string {
	method := strings.ToLower(r.Method)
	path := ""

	for _, seg := range strings.Split(r.Path, "/")[1:] {
		path += seg + "-"
	}

	return fmt.Sprintf("%v-%v", method, path[:len(path)-1])
}

func (a App) GetYAMLDefinition() ([]byte, error) {
	return yaml.Marshal(a)
}
