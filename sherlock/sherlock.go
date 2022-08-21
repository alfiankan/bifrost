package sherlock

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
)

// Deps deps contains pkg path, struct name and type (package.structname)
type Deps struct {
	PkgPath    string
	StructName string
	Type       string
}

// Sherlock contains Dependencies maps and overriders map
type Sherlock struct {
	overriders      map[string][]Deps
	Dependencies    map[string][]Deps
	GlobalOveriders []Deps
	tempDependecies []Deps
	path            string
	pkgName         string
}

// New instance new sherlock
func New() *Sherlock {
	return &Sherlock{
		overriders:      map[string][]Deps{},
		Dependencies:    map[string][]Deps{},
		tempDependecies: []Deps{},
		GlobalOveriders: []Deps{},
	}
}

// SetPath set save location for wire.go file adn wire gen output location path
func (sr *Sherlock) SetPath(path string) *Sherlock {
	sr.path = path
	return sr
}

// SetPkgName set pkg name for generated wire
func (sr *Sherlock) SetPkgName(name string) *Sherlock {
	sr.pkgName = name
	return sr
}

// SetGlobalInject inject or override globaly
func (sr *Sherlock) SetGlobalInject(obj any) *Sherlock {
	if reflect.TypeOf(obj).String() == "string" {
		fmt.Println("PRIMITIVE")
	} else {
		el := reflect.TypeOf(obj)

		for rootName, deps := range sr.Dependencies {
			sr.overriders[rootName] = append(sr.overriders[rootName], Deps{
				PkgPath:    el.PkgPath(),
				StructName: el.Name(),
				Type:       el.String(),
			})

			filteredDependecies := []Deps{}

			for _, d := range deps {
				if !IsOverrided(d.Type, sr.overriders[rootName]) {
					filteredDependecies = append(filteredDependecies, d)
				}
			}
			sr.Dependencies[rootName] = filteredDependecies
		}

	}
	return sr
}

// Gen generate wire.go file
func (sr *Sherlock) Gen() error {

	fmt.Println("GENERATING WIRE FILE")

	wireFile1 := fmt.Sprintf(wireHead, "main")
	wireFile2 := wireInit
	imports := ""
	wireFile3 := ""

	if sr.pkgName != "" {
		wireFile1 = fmt.Sprintf(wireHead, sr.pkgName)
	}

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

	wireFileName := filepath.Join("wire.go")
	if sr.path != "" {
		wireFileName = filepath.Join(sr.path, "wire.go")
	}

	os.WriteFile(wireFileName, []byte(wireFile1+imports+wireFile2+wireFile3), os.ModePerm)

	_, err := exec.Command("go", "fmt", wireFileName).Output()
	if err != nil {
		fmt.Println("Formating error make sure gofmt installed", err.Error())
	}
	_, err = exec.Command("wire", sr.path).Output()
	if err != nil {
		fmt.Println("Wiring error make sure google wire installed", err.Error())
	}
	return nil
}

// Add add new root deps for auto wiring
func (sr *Sherlock) Add(obj any, ovr ...any) *Sherlock {

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

	return sr
}
