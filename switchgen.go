// switchgen reflect.Kind

package switchgen

import (
	"fmt"
	"go/types"
	"strings"

	"golang.org/x/tools/go/packages"
)

type Cases []string

func (cc Cases) String() string {
	var b strings.Builder
	b.WriteString("switch {\n")
	for _, c := range cc {
		b.WriteString("case ")
		b.WriteString(c)
		b.WriteString(":\n")
	}
	b.WriteString("}\n")
	return b.String()
}

type Errors []packages.Error

func (ee Errors) Error() string {
	var b strings.Builder
	b.WriteString(ee[0].Error())
	if len(ee) > 1 {
		for i := 1; i < len(ee); i++ {
			b.WriteString("; ")
			b.WriteString(ee[i].Error())
		}
	}
	return b.String()
}

func Generate(pkgName, name string) (Cases, error) {
	conf := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedTypesInfo,
	}

	pp, err := packages.Load(conf, pkgName)
	if err != nil {
		return nil, fmt.Errorf("error while loading %q: %w", pkgName, err)
	}

	pkg := pp[0]

	if len(pkg.Errors) > 0 {
		return nil, Errors(pkg.Errors)
	}

	scope := pkg.Types.Scope()
	baseObj := scope.Lookup(name)

	_, ok := baseObj.(*types.TypeName)
	if !ok {
		return nil, fmt.Errorf("expected TypeName, got %T", baseObj)
	}

	baseInterface, isInterface := baseObj.Type().(*types.Named).Underlying().(*types.Interface)

	cases := Cases{}

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

	return cases, nil
}
