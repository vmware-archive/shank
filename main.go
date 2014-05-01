package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/codegangsta/cli"

	warden "github.com/cloudfoundry-incubator/garden/protocol"
	"github.com/cloudfoundry-incubator/gordon"
	"github.com/cloudfoundry-incubator/gordon/connection"
)

var shankRCFile = filepath.Join(os.Getenv("HOME"), ".shankrc")

func main() {
	app := cli.NewApp()
	app.Name = "shank"
	app.Usage = "Warden server CLI"

	app.Flags = []cli.Flag{
		cli.StringFlag{"network", "unix", "server network type (tcp, unix)"},
		cli.StringFlag{"addr", "/tmp/warden.sock", "server network address"},
	}

	app.Commands = []cli.Command{
		{
			Name:  "target",
			Usage: "Save -network and -addr to ~/.shankrc.",
			Action: func(c *cli.Context) {
				file, err := os.OpenFile(shankRCFile, os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic("cannot write to " + shankRCFile)
				}

				encoder := json.NewEncoder(file)
				encoder.Encode(map[string]string{
					"network": c.GlobalString("network"),
					"addr":    c.GlobalString("addr"),
				})
			},
		},

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
		generateCommand(reflect.ValueOf(&warden.ListRequest{})),
		generateCommand(reflect.ValueOf(&warden.NetInRequest{})),
		generateCommand(reflect.ValueOf(&warden.NetOutRequest{})),
		generateCommand(reflect.ValueOf(&warden.PingRequest{})),
		generateCommand(reflect.ValueOf(&warden.RunRequest{})),
		generateCommand(reflect.ValueOf(&warden.AttachRequest{})),
		generateCommand(reflect.ValueOf(&warden.StopRequest{})),
		generateCommand(reflect.ValueOf(&warden.CapacityRequest{})),
	}

	app.Run(os.Args)
}

func generateCommand(request reflect.Value) cli.Command {
	typ := request.Elem().Type()

	commandName := lowercase(strings.TrimSuffix(typ.Name(), "Request"))

	usage := USAGE[commandName]

	flags := []cli.Flag{}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		flag, ok := flagForField(field, usage.Flags)
		if ok {
			flags = append(flags, flag)
		}
	}

	return cli.Command{
		Name:        commandName,
		Flags:       flags,
		Usage:       usage.Usage,
		Description: usage.Description,
		Action: func(c *cli.Context) {
			cp := connectionInfo(c)

			conn, err := cp.ProvideConnection()
			if err != nil {
				fmt.Println("failed to connect to warden:", err)
				os.Exit(1)
			}

			request := requestFromInput(request, flags, c)

			response := warden.ResponseMessageForType(warden.TypeForMessage(request))

			encoder := json.NewEncoder(os.Stdout)

			if commandName == "attach" {
				err := conn.SendMessage(request)
				if err != nil {
					fmt.Println("request-response failed:", err)
					os.Exit(1)
				}

				streamProcessPayloads(conn, encoder)

				return
			}

			res, err := conn.RoundTrip(request, response)
			if err != nil {
				fmt.Println("request-response failed:", err)
				os.Exit(1)
			}

			encoder.Encode(res)

			if commandName == "run" {
				streamProcessPayloads(conn, encoder)
			}
		},
	}
}

func connectionInfo(c *cli.Context) gordon.ConnectionProvider {
	config := map[string]string{
		"network": c.GlobalString("network"),
		"addr":    c.GlobalString("addr"),
	}

	file, err := os.Open(shankRCFile)
	if err == nil {
		decoder := json.NewDecoder(file)
		err := decoder.Decode(&config)
		if err != nil {
			panic("cannot decode " + shankRCFile + ": " + err.Error())
		}
	}

	return &gordon.ConnectionInfo{
		Network: config["network"],
		Addr:    config["addr"],
	}
}

func streamProcessPayloads(conn *connection.Connection, encoder *json.Encoder) {
	for {
		payload := &warden.ProcessPayload{}

		res, err := conn.ReadResponse(payload)
		if err != nil {
			fmt.Println("stream failed:", err)
			os.Exit(1)
		}

		encoder.Encode(res)

		if payload.ExitStatus != nil {
			break
		}
	}
}
