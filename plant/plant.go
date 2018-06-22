package plant

import (
	"bytes"
	"fmt"
	//"os"
	//"path"
	//"path/filepath"
	"text/template"
	"strconv"
	//"strings"
)


func CreateMainFile(port int, project string, root string) {
	mainTemplate :=  `package main

import (
	"github.com/{{ .root }}/{{ .project }}/config"
	"github.com/{{ .root }}/{{ .project }}/services"
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
	fmt.Println(buf.String())
}
