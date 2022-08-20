package sherlock

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Deps struct {
	PkgPath    string
	StructName string
	Type       string
}

type Sherlock struct {
	overriders      map[string][]Deps
	Dependencies    map[string][]Deps
	tempDependecies []Deps
}

func New() *Sherlock {
	return &Sherlock{
		overriders:      map[string][]Deps{},
		Dependencies:    map[string][]Deps{},
		tempDependecies: []Deps{},
	}
}

func (sr *Sherlock) Gen() error {

	fmt.Println("GENERATING WIRE FILE")

	wireFile1 := wireHead
	wireFile2 := wireInit
	imports := ""
	wireFile3 := ""

	for rootDepsName, deps := range sr.Dependencies {

		t := table.NewWriter()
		t.AppendHeader(table.Row{"PKG", "Name", "Type", "Is Overrider"})
		t.SetTitle(rootDepsName)

		wireBuild := ""
		ovveriderParam := []string{}

		for _, d := range deps {

			imports += fmt.Sprintf(`"%s"
			`, d.PkgPath)

			wireBuild += fmt.Sprintf(`wire.Struct(new(%s), "*"),
			`, d.Type)

			t.AppendRow(table.Row{d.PkgPath, d.StructName, d.Type, false})

		}

		for _, d := range sr.overriders[rootDepsName] {
			ovveriderParam = append(ovveriderParam, fmt.Sprintf(`%s %s`, strings.ToLower(d.StructName), d.Type))
			t.AppendRow(table.Row{d.PkgPath, d.StructName, d.Type, true})

		}

		wireFile3 += fmt.Sprintf(initializerTemplate, rootDepsName, strings.Join(ovveriderParam, ","), deps[0].Type, wireBuild, deps[0].Type)

		fmt.Println(t.Render())
	}

	os.WriteFile("wire.go", []byte(wireFile1+imports+wireFile2+wireFile3), os.ModePerm)

	_, err := exec.Command("go", "fmt", "wire.go").Output()
	if err != nil {
		fmt.Println("Formating error make sure gofmt installed", err.Error())
	}
	_, err = exec.Command("wire").Output()
	if err != nil {
		fmt.Println("Wiring error make sure google wire installed", err.Error())
	}
	return nil
}

func (sr *Sherlock) Add(obj any, ovr ...any) {

	ParseDependecies(&sr.tempDependecies, obj)

	e := reflect.TypeOf(obj)

	for _, ov := range ovr {
		if reflect.TypeOf(ov).String() == "string" {
			fmt.Println("PRIMITIVE")
		} else {
			el := reflect.TypeOf(ov)

			sr.overriders[e.Name()] = append(sr.overriders[e.Name()], Deps{
				PkgPath:    el.PkgPath(),
				StructName: el.Name(),
				Type:       el.String(),
			})
		}
	}

	filteredDependecies := []Deps{}

	for _, d := range sr.tempDependecies {
		if !IsOverrided(d.Type, sr.overriders[e.Name()]) {
			filteredDependecies = append(filteredDependecies, d)
		}
	}

	sr.Dependencies[e.Name()] = filteredDependecies

	sr.tempDependecies = []Deps{}

}
