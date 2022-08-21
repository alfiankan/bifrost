package sherlock

import (
	"fmt"
	"reflect"
)

// ParseDependecies walk to struct field and find dependencies resursively
// slice of deps must be pointer, object is variadic
func ParseDependecies(tempDeps *[]Deps, obj ...any) {

	for _, objv := range obj {

		switch objv.(type) {
		case reflect.Type:

			if IsPrimitive(objv.(reflect.Type).String()) {
				fmt.Println("PRIMITIVE")
			} else {
				e := objv.(reflect.Type)

				var requiredDeps []string
				for x := 0; x < e.NumField(); x++ {
					requiredDeps = append(requiredDeps, e.Field(x).Type.String())
				}

				*tempDeps = append(*tempDeps, Deps{
					PkgPath:    e.PkgPath(),
					StructName: e.Name(),
					Type:       objv.(reflect.Type).String(),
				})

				for x := 0; x < e.NumField(); x++ {
					ParseDependecies(tempDeps, e.Field(x).Type)
				}
			}
		default:
			if IsPrimitive(reflect.TypeOf(objv).String()) {
				fmt.Println("PRIMITIVE")
			} else {
				e := reflect.TypeOf(objv)

				var requiredDeps []string
				for x := 0; x < e.NumField(); x++ {
					requiredDeps = append(requiredDeps, e.Field(x).Type.String())
				}

				*tempDeps = append(*tempDeps, Deps{
					PkgPath:    e.PkgPath(),
					StructName: e.Name(),
					Type:       e.String(),
				})

				for x := 0; x < e.NumField(); x++ {
					ParseDependecies(tempDeps, e.Field(x).Type)
				}
			}
		}

	}
}

// IsOverrided Check If actual deps overriden
func IsOverrided(depsType string, overriders []Deps) bool {
	for _, d := range overriders {
		if d.Type == depsType {
			return true
		}
	}
	return false
}

// IsPrimitive Check if deps are primitives
func IsPrimitive(srcType string) bool {
	primitives := []string{
		"complex64",
		"complex128",
		"float32",
		"float64",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"int8",
		"int16",
		"int32",
		"int64",
		"uintptr",
		"error",
		"bool",
	}

	for _, prv := range primitives {
		if srcType == prv {
			return true
		}
	}
	return false
}
