package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/codegangsta/cli"

	"github.com/vito/gordon/warden"
)

func main() {
	app := cli.NewApp()
	app.Name = "shank"
	app.Usage = "Warden server CLI"

	app.Commands = []cli.Command{
		generateCommand(reflect.ValueOf(&warden.CopyInRequest{})),
		generateCommand(reflect.ValueOf(&warden.CopyOutRequest{})),
		generateCommand(reflect.ValueOf(&warden.CreateRequest{})),
		generateCommand(reflect.ValueOf(&warden.DestroyRequest{})),
		generateCommand(reflect.ValueOf(&warden.EchoRequest{})),
		generateCommand(reflect.ValueOf(&warden.InfoRequest{})),
		generateCommand(reflect.ValueOf(&warden.LimitBandwidthRequest{})),
		generateCommand(reflect.ValueOf(&warden.LimitCpuRequest{})),
		generateCommand(reflect.ValueOf(&warden.LimitDiskRequest{})),
		generateCommand(reflect.ValueOf(&warden.LimitMemoryRequest{})),
		generateCommand(reflect.ValueOf(&warden.LinkRequest{})),
		generateCommand(reflect.ValueOf(&warden.ListRequest{})),
		generateCommand(reflect.ValueOf(&warden.NetInRequest{})),
		generateCommand(reflect.ValueOf(&warden.NetOutRequest{})),
		generateCommand(reflect.ValueOf(&warden.PingRequest{})),
		generateCommand(reflect.ValueOf(&warden.RunRequest{})),
		generateCommand(reflect.ValueOf(&warden.SpawnRequest{})),
		generateCommand(reflect.ValueOf(&warden.StopRequest{})),
		generateCommand(reflect.ValueOf(&warden.StreamRequest{})),
	}

	app.Run(os.Args)
}

func generateCommand(request reflect.Value) cli.Command {
	typ := request.Elem().Type()

	commandName := lowercase(strings.TrimSuffix(typ.Name(), "Request"))

	flags := []cli.Flag{}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		flag, ok := flagForField(field)
		if ok {
			flags = append(flags, flag)
		}
	}

	return cli.Command{
		Name:  commandName,
		Flags: flags,
		Action: func(c *cli.Context) {
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

				case cli.IntFlag:
					num := uint32(c.Int(flag.Name))
					field.Set(reflect.ValueOf(&num))

				case JSONFlag:
					val := reflect.New(flag.Flag.(JSONFlag).Type)

					str := c.String(flag.Name)
					err := json.Unmarshal([]byte(str), val.Interface())
					if err != nil {
						println(err.Error())
						os.Exit(1)
					}

					field.Set(reflect.Indirect(val))
				}
			}

			fmt.Println("request:", request.Interface())
		},
	}
}
