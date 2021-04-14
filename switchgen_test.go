package switchgen

import "testing"

func TestGenerate(t *testing.T) {
	tests := []struct {
		pkg  string
		name string
		code string
	}{
		{
			"reflect", "Kind", `switch _ {
case reflect.Array:
case reflect.Bool:
case reflect.Chan:
case reflect.Complex128:
case reflect.Complex64:
case reflect.Float32:
case reflect.Float64:
case reflect.Func:
case reflect.Int:
case reflect.Int16:
case reflect.Int32:
case reflect.Int64:
case reflect.Int8:
case reflect.Interface:
case reflect.Invalid:
case reflect.Map:
case reflect.Ptr:
case reflect.Slice:
case reflect.String:
case reflect.Struct:
case reflect.Uint:
case reflect.Uint16:
case reflect.Uint32:
case reflect.Uint64:
case reflect.Uint8:
case reflect.Uintptr:
case reflect.UnsafePointer:
}
`,
		},
		{
			"go/ast", "Node", `switch _.(type) {
case *ast.ArrayType:
case *ast.AssignStmt:
case *ast.BadDecl:
case *ast.BadExpr:
case *ast.BadStmt:
case *ast.BasicLit:
case *ast.BinaryExpr:
case *ast.BlockStmt:
case *ast.BranchStmt:
case *ast.CallExpr:
case *ast.CaseClause:
case *ast.ChanType:
case *ast.CommClause:
case *ast.Comment:
case *ast.CommentGroup:
case *ast.CompositeLit:
case ast.Decl:
case *ast.DeclStmt:
case *ast.DeferStmt:
case *ast.Ellipsis:
case *ast.EmptyStmt:
case ast.Expr:
case *ast.ExprStmt:
case *ast.Field:
case *ast.FieldList:
case *ast.File:
case *ast.ForStmt:
case *ast.FuncDecl:
case *ast.FuncLit:
case *ast.FuncType:
case *ast.GenDecl:
case *ast.GoStmt:
case *ast.Ident:
case *ast.IfStmt:
case *ast.ImportSpec:
case *ast.IncDecStmt:
case *ast.IndexExpr:
case *ast.InterfaceType:
case *ast.KeyValueExpr:
case *ast.LabeledStmt:
case *ast.MapType:
case *ast.Package:
case *ast.ParenExpr:
case *ast.RangeStmt:
case *ast.ReturnStmt:
case *ast.SelectStmt:
case *ast.SelectorExpr:
case *ast.SendStmt:
case *ast.SliceExpr:
case ast.Spec:
case *ast.StarExpr:
case ast.Stmt:
case *ast.StructType:
case *ast.SwitchStmt:
case *ast.TypeAssertExpr:
case *ast.TypeSpec:
case *ast.TypeSwitchStmt:
case *ast.UnaryExpr:
case *ast.ValueSpec:
}
`,
		},
		{
			"golang.org/x/tools/go/packages", "LoadMode", `switch _ {
case packages.LoadAllSyntax:
case packages.LoadFiles:
case packages.LoadImports:
case packages.LoadSyntax:
case packages.LoadTypes:
case packages.NeedCompiledGoFiles:
case packages.NeedDeps:
case packages.NeedExportsFile:
case packages.NeedFiles:
case packages.NeedImports:
case packages.NeedModule:
case packages.NeedName:
case packages.NeedSyntax:
case packages.NeedTypes:
case packages.NeedTypesInfo:
case packages.NeedTypesSizes:
}
`,
		},
	}

	for _, test := range tests {
		cases, err := Generate(test.pkg, test.name)
		if err != nil {
			t.Fatalf("%s.%s: %s", test.pkg, test.name, err)
		}

		if expected, got := test.code, cases.String(); got != expected {
			t.Errorf("failed: %s.%s", test.pkg, test.name)
			t.Log(cases.String())
		}
	}

	_, err := Generate("xyz", "V")
	if err == nil {
		t.Fatalf("should err")
	}
}
