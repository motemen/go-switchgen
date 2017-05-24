// switchgen reflect.Kind

package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"go/types"
	"golang.org/x/tools/go/loader"
)

func main() {
	log.SetPrefix("switchgen: ")
	log.SetFlags(0)

	var (
		pkgAndName = os.Args[1]
	)

	pair := strings.Split(pkgAndName, ".")
	if len(pair) != 2 {
		log.Fatal("usage: switchgen <pkg>.<name>")
	}

	var (
		pkgName = pair[0]
		name    = pair[1]
	)

	conf := loader.Config{}
	conf.FromArgs([]string{pkgName}, false)
	prog, err := conf.Load()
	if err != nil {
		log.Fatal(err)
	}

	pkg := prog.Package(pkgName)

	scope := pkg.Pkg.Scope()
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
