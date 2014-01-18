package main

import (
	"reflect"

	"github.com/codegangsta/cli"
)

type JSONFlag struct {
	Type reflect.Type

	cli.StringFlag
}

type Flag struct {
	Field string

	Name     string
	Required bool

	cli.Flag
}
