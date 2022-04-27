package main

import (
	"fmt"
	"os"
	"text/template"
)

type Manifest struct {
	ClusterDomain string
	MasterIps     []string
	NodeIps       map[string]string
	Name string
	Name2 string
	Name3 string
}

var manifestTpl = `
[masters]
{{range $index, $value := .MasterIps}}
   master{{$index}}.{{$.ClusterDomain}} {{$value}}{{end}}
[nodes]
{{range $index, $value := .NodeIps}} 
   out.{{$index}} = {{$value}}
{{end}}
{{if and .Name .Name3}} 
123
{{else if .Name2}}
Anonymous
{{else if .Name3}}
youzu
{{end}}!
`

func main() {
	manifest := Manifest{
		ClusterDomain: "okd.local",
		MasterIps:     []string{"10.10.13.2", "10.10.13.3"},
		NodeIps: map[string]string{
			"bushi":"mouhouzi",
			"bushi1":"mouhouzi1",
			"bushi2":"mouhouzi2",
		},
		//Name: "nihao",
		Name3: "hah",
	}
	manifestTpl, _ := template.New("manifest").Parse(manifestTpl)
	manifestTpl.Execute(os.Stdout, manifest)
	fmt.Println("")
}