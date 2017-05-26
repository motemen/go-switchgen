// switchgen reflect.Kind

package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"go/importer"
	"go/types"
)

func main() {
	log.SetPrefix("switchgen: ")
	log.SetFlags(0)

	pkgAndName := os.Args[1]

	n := strings.LastIndex(pkgAndName, ".")
	if n == -1 {
		log.Fatal("usage: switchgen <pkg>.<name>")
	}

	var (
		pkgName = pkgAndName[0:n]
		name    = pkgAndName[n+1:]
	)

	pkg, err := importer.Default().Import(pkgName)
	if err != nil {
		log.Fatal(err)
	}

	scope := pkg.Scope()
	baseObj := scope.Lookup(name)

	_, ok := baseObj.(*types.TypeName)
	if !ok {
		log.Fatalf("got %T", baseObj)
	}

	baseInterface, isInterface := baseObj.Type().(*types.Named).Underlying().(*types.Interface)

	cases := []string{}

	typ := baseObj.Type()
	for _, name := range scope.Names() {
		obj := scope.Lookup(name)
		if obj == baseObj {
			continue
		}
		if !obj.Exported() {
			continue
		}

		if isInterface {
			if types.Implements(obj.Type(), baseInterface) {
				cases = append(cases, fmt.Sprintf("%s.%s", obj.Pkg().Name(), obj.Name()))
			} else if types.Implements(types.NewPointer(obj.Type()), baseInterface) {
				cases = append(cases, fmt.Sprintf("*%s.%s", obj.Pkg().Name(), obj.Name()))
			}
		} else {
			if obj.Type() == typ {
				cases = append(cases, fmt.Sprintf("%s.%s", obj.Pkg().Name(), obj.Name()))
			}
		}
	}

	fmt.Println("switch {")
	for _, c := range cases {
		fmt.Printf("case %s:\n", c)
	}
	fmt.Println("}")
}
