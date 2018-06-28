/**
* Copyright © 2018, ACM@UIUC
*
* This file is part of the Arbor Project.
*
* The Arbor Project is open source software, released under the University of
* Illinois/NCSA Open Source License. You should have received a copy of
* this license in a file with the distribution.
**/

package plant

import (
	"bytes"
	"text/template"
	"strconv"
	"fmt"
	"os"
	
	"github.com/arbor-dev/seedling/helper"
)


func CreateMainFile(port int, project string, root string) {
	mainTemplate :=  `package main

import (
	// Add proper imports
    //"github.com/{{ .root }}/{{ .project }}/config"
	//"github.com/{{ .root }}/{{ .project }}/services"
	"github.com/arbor-dev/arbor"
)

func main() {
	config.LoadArborConfig()
	Routes := services.RegisterAPIs()
	arbor.Boot(Routes, "0.0.0.0", {{ .port }})
}`

	data := make(map[string]interface{})
	data["root"] = root
	data["project"] = project
	data["port"] = strconv.Itoa(port)

	tmpl, _ := template.New("").Parse(mainTemplate)

	buf := new(bytes.Buffer)
	tmpl.Execute(buf, data)

	err := helper.WriteToFile(buf.String(), project, "main.go", true)

	if (err != nil) {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func CreateServicesFiles(project string) {
	servicesTemplate :=  `package services

import (
	"fmt"
	"net/http"

	"github.com/arbor-dev/arbor"
)

var Routes = arbor.RouteCollection{
	arbor.Route{
		"Index",
		"GET",
		"/",
		Index,
	},
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{{ .project }} Project!\n")
}

func RegisterAPIs() arbor.RouteCollection {
	Routes = append(Routes, ExampleRoutes...)
	return Routes
}`

	data := make(map[string]interface{})
	data["project"] = project

	tmpl, _ := template.New("").Parse(servicesTemplate)

	buf := new(bytes.Buffer)
	tmpl.Execute(buf, data)

	err := helper.WriteToFile(buf.String(), project + "/services", "services.go", true)

	if (err != nil) {
		fmt.Println(err)
		os.Exit(-1)
	}

	createExampleServiceFile(project, project + "/services")
}

func createExampleServiceFile(project string, dir string) {
	exampleServiceTemplate := `package services

import (
	"net/http"

	"github.com/acm-uiuc/{{ .project }}/config"
	"github.com/arbor-dev/arbor"
)

const ExampleURL string = config.ExampleURL

const ExampleFormat string = "JSON"

//API Interface
var ExampleRoutes = arbor.RouteCollection{
	arbor.Route{
		"NewExample",
		"POST",
		"/example",
		NewExample,
	},
	arbor.Route{
		"DeleteExample",
		"DELETE",
		"/example?id={id}",
		DeleteExample,
	},
	arbor.Route{
		"GetExample",
		"GET",
		"/example/{id}",
		GetExample,
	},
	arbor.Route{
		"UpdateExample",
		"PUT",
		"/example/{id}",
		UpdateExample,
	},
}

// arbor.Route handler
// w = writer, r = reader
func NewExample(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, ExampleURL+r.URL.String(), ExampleFormat, "", r)
}

func DeleteExample(w http.ResponseWriter, r *http.Request) {
	arbor.DELETE(w, ExampleURL+r.URL.String(), ExampleFormat, "", r)
}

func GetExample(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, ExampleURL+r.URL.String(), ExampleFormat, "", r)
}

func UpdateExample(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, ExampleURL+r.URL.String(), ExampleFormat, "", r)
}`

	data := make(map[string]interface{})
	data["project"] = project

	tmpl, _ := template.New("").Parse(exampleServiceTemplate)

	buf := new(bytes.Buffer)
	tmpl.Execute(buf, data)

	err := helper.WriteToFile(buf.String(), project + "/services", "exampleservice.go", false)

	if (err != nil) {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func CreateConfigFile(project string) {
	configTemplate := `package config

import (
    "github.com/arbor-dev/arbor/proxy"
    "github.com/arbor-dev/arbor/security"
)

// {{ .project }} Config

// Example service URL
const ExampleURL = "http://localhost:5656"

//Arbor configurations
func LoadArborConfig() {
    security.AccessLogLocation = "log/access.log"
    security.ClientRegistryLocation = "clients.db"
    proxy.AccessControlPolicy = "*"
}`

	data := make(map[string]interface{})
	data["project"] = project

	tmpl, _ := template.New("").Parse(configTemplate)

	buf := new(bytes.Buffer)
	tmpl.Execute(buf, data)

	err := helper.WriteToFile(buf.String(), project + "/config", "config.go.template", true)

	if (err != nil) {
		fmt.Println(err)
		os.Exit(-1)
	}
}
