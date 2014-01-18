package main

import (
	"reflect"
	"strings"

	"github.com/codegangsta/cli"
)

func flagForField(field reflect.StructField) (*Flag, bool) {
	if strings.HasPrefix(field.Name, "XXX") {
		return nil, false
	}

	flagName := lowercase(field.Name)

	flag, ok := flagForType(flagName, field.Type)
	if !ok {
		return nil, false
	}

	required := false

	protoTags := strings.Split(field.Tag.Get("protobuf"), ",")
	for _, tag := range protoTags {
		switch tag {
		case "opt":
			required = false
		case "req":
			required = true
		}
	}

	return &Flag{
		Field: field.Name,

		Name:     flagName,
		Required: required,

		Flag: flag,
	}, ok
}

func flagForType(name string, typ reflect.Type) (cli.Flag, bool) {
	switch typ.Kind() {
	case reflect.Ptr:
		return flagForType(name, typ.Elem())
	case reflect.String:
		return cli.StringFlag{name, "", "string value"}, true
	case reflect.Uint32:
		return cli.IntFlag{name, 0, "integer value"}, true
	default:
		return JSONFlag{typ, cli.StringFlag{name, "", "json value"}}, true
	}

	return nil, false
}

func lowercase(str string) string {
	return strings.ToLower(str[:1]) + str[1:]
}
