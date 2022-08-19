package bifrost

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/KyleBanks/depth"
	"github.com/jedib0t/go-pretty/v6/table"
)

type Bifrost struct {
	overrides []Deps
}

type Deps struct {
	PkgPath    string
	StructName string
	Type       string
}

var dependecies = []Deps{}
var t = table.NewWriter()

func parseDependecies(obj ...any) {

	for _, objv := range obj {

		switch objv.(type) {
		case reflect.Type:
			//fmt.Println("TRUE TYPE")

			//fmt.Println("TYPE", objv.(reflect.Type))
			if objv.(reflect.Type).String() == "string" || objv.(reflect.Type).String() == "int64" {
				fmt.Println("PRIMITIVE")
			} else {
				e := objv.(reflect.Type)

				var requiredDeps []string
				for x := 0; x < e.NumField(); x++ {
					requiredDeps = append(requiredDeps, e.Field(x).Type.String())
				}

				t.AppendRow(table.Row{e.PkgPath(), e.Name(), objv.(reflect.Type), e.NumField(), strings.Join(requiredDeps, " ")})

				dependecies = append(dependecies, Deps{
					PkgPath:    e.PkgPath(),
					StructName: e.Name(),
					Type:       objv.(reflect.Type).String(),
				})

				for x := 0; x < e.NumField(); x++ {
					// fmt.Println("=================")
					// fmt.Println(e.Name())

					// fmt.Println(e.Field(x).Type)

					// fmt.Println("=================")

					parseDependecies(e.Field(x).Type)

				}
			}
		default:
			//fmt.Println("TYPE", reflect.TypeOf(objv).String())
			//fmt.Println("DEFAULT")
			if reflect.TypeOf(objv).String() == "string" {
				fmt.Println("PRIMITIVE")
			} else {
				e := reflect.TypeOf(objv)

				var requiredDeps []string
				for x := 0; x < e.NumField(); x++ {
					requiredDeps = append(requiredDeps, e.Field(x).Type.String())
				}

				t.AppendRow(table.Row{e.PkgPath(), e.Name(), e, e.NumField(), strings.Join(requiredDeps, " ")})

				dependecies = append(dependecies, Deps{
					PkgPath:    e.PkgPath(),
					StructName: e.Name(),
					Type:       e.String(),
				})

				for x := 0; x < e.NumField(); x++ {
					// fmt.Println("=================")
					// fmt.Println(e.Name())
					// fmt.Println(e.Field(x).Type)

					// fmt.Println("=================")

					parseDependecies(e.Field(x).Type)

				}
			}
		}

	}
}
func writePkgJSON(w io.Writer, p depth.Pkg) {
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	e.Encode(p)
}

func (bf *Bifrost) Gen(obj any, ovr ...any) {
	t.ResetRows()
	t.ResetHeaders()

	bf.overrides = []Deps{}
	dependecies = []Deps{}

	for _, ov := range ovr {
		if reflect.TypeOf(ov).String() == "string" {
			fmt.Println("PRIMITIVE")
		} else {
			e := reflect.TypeOf(ov)

			bf.overrides = append(bf.overrides, Deps{
				PkgPath:    e.PkgPath(),
				StructName: e.Name(),
				Type:       e.String(),
			})

		}
	}

	parseDependecies(obj)

	e := reflect.TypeOf(obj)

	fmt.Println(fmt.Sprintf("========= %s DEPENDENCIES ===========", e))
	tTemp := table.Table{}
	tTemp.Render()

	t.AppendHeader(table.Row{"pkg", "struct", "type", "total field", "reuired"})

	fmt.Println(t.Render())

	fmt.Println("DEPENDENCIES", dependecies)

	fmt.Println("OVERRIDERS", bf.overrides)

}
