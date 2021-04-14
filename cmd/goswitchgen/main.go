// switchgen reflect.Kind

package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/motemen/go-switchgen"
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

	cases, err := switchgen.Generate(pkgName, name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(cases.String())
}
