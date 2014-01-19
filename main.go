package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/codegangsta/cli"

	"github.com/vito/gordon"
	"github.com/vito/gordon/warden"
)

func main() {
	app := cli.NewApp()
	app.Name = "shank"
	app.Usage = "Warden server CLI"

	app.Flags = []cli.Flag{
		cli.StringFlag{"network", "unix", "server network type (tcp, unix)"},
		cli.StringFlag{"addr", "/tmp/warden.sock", "server network address"},
	}

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
			cp := &gordon.ConnectionInfo{
				Network: c.GlobalString("network"),
				Addr: c.GlobalString("addr"),
			}

			conn, err := cp.ProvideConnection()
			if err != nil {
				fmt.Println("failed to connect to warden:", err)
				os.Exit(1)
			}

			request := requestFromInput(request, flags, c)

			response := warden.ResponseMessageForType(warden.TypeForMessage(request))

			res, err := conn.RoundTrip(request, response)
			if err != nil {
				fmt.Println("request-response failed:", err)
				os.Exit(1)
			}

			fmt.Println("request:", request)
			fmt.Println("response:", res)
		},
	}
}
