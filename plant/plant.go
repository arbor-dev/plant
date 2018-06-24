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

	err := helper.WriteToFile(buf.String(), project, "main.go")

	if (err != nil) {
		fmt.Println(err)
		os.Exit(-1)
	}
}
