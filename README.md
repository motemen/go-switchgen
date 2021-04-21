# goswitchgen

A command line tool for Go that generates switch statement from type name.

## Installation

    % go install github.com/motemen/go-switchgen/cmd/goswitchgen@latest

## Usage

    % goswitchgen <package>.<type>

eg. [`reflect.Kind`](https://pkg.go.dev/reflect#Kind) (enums)

    % goswitchgen reflect.Kind
    switch _ {
    case reflect.Array:
    case reflect.Bool:
    [...]
    case reflect.Uintptr:
    case reflect.UnsafePointer:
    }

eg. [`go/ast.Node`](https://pkg.go.dev/go/ast#Node) (interfaces)

    % goswitchgen go/ast.Node
    switch _.(type) {
    case *ast.ArrayType:
    case *ast.AssignStmt:
    [...]
    case *ast.UnaryExpr:
    case *ast.ValueSpec:
    }

For non-standard modules, ones must be specified on go.mod.

    % cat go.mod | grep github.com/slack-go/slack
        github.com/slack-go/slack v0.8.2
    % goswitchgen github.com/slack-go/slack.MessageBlockType
    switch _ {
    case slack.MBTAction:
    [...]
    case slack.MBTSection:
    }
