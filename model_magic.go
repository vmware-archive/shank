package main

import (
	"encoding/json"
	"os"
	"reflect"
	"strings"

	"code.google.com/p/gogoprotobuf/proto"
	"github.com/codegangsta/cli"
)

func flagForField(field reflect.StructField, usage map[string]string) (*Flag, bool) {
	if strings.HasPrefix(field.Name, "XXX") {
		return nil, false
	}

	flagName := lowercase(field.Name)

	flag, ok := flagForType(flagName, field.Type, usage[flagName])
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

func flagForType(name string, typ reflect.Type, usage string) (cli.Flag, bool) {
	switch typ.Kind() {
	case reflect.Ptr:
		return flagForType(name, typ.Elem(), usage)
	case reflect.String:
		return cli.StringFlag{name, "", usage}, true
	default:
		return JSONFlag{typ, cli.StringFlag{name, "", usage}}, true
	}

	return nil, false
}

func requestFromInput(request reflect.Value, flags []cli.Flag, c *cli.Context) proto.Message {
	for _, f := range flags {
		flag := f.(*Flag)

		if !c.IsSet(flag.Name) {
			if flag.Required {
				println("missing required flag '" + flag.Name + "'")
				os.Exit(1)
			}

			continue
		}

		field := request.Elem().FieldByName(flag.Field)

		switch flag.Flag.(type) {
		case cli.StringFlag:
			str := c.String(flag.Name)
			field.Set(reflect.ValueOf(&str))

		case JSONFlag:
			val := reflect.New(flag.Flag.(JSONFlag).Type)

			str := c.String(flag.Name)
			err := json.Unmarshal([]byte(str), val.Interface())
			if err != nil {
				println(err.Error())
				os.Exit(1)
			}

			if field.Type().Kind() == reflect.Ptr {
				field.Set(val)
			} else {
				field.Set(reflect.Indirect(val))
			}
		}
	}

	return request.Interface().(proto.Message)
}
