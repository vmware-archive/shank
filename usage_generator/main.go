package main

import (
	"flag"
	"path/filepath"
	"fmt"
	"encoding/json"

	"github.com/vito/shank/usage"
)

var protobufs = flag.String(
	"protobufs",
	"./warden-protocol",
	"Directory containing the .proto files.",
)

func main() {
	flag.Parse()

	usage := map[string]usage.Usage{
		"copyIn":         usage.Parse(filepath.Join(*protobufs, "copy_in.proto")),
		"copyOut":        usage.Parse(filepath.Join(*protobufs, "copy_out.proto")),
		"create":         usage.Parse(filepath.Join(*protobufs, "create.proto")),
		"destroy":        usage.Parse(filepath.Join(*protobufs, "destroy.proto")),
		"echo":           usage.Parse(filepath.Join(*protobufs, "echo.proto")),
		"info":           usage.Parse(filepath.Join(*protobufs, "info.proto")),
		"limitBandwidth": usage.Parse(filepath.Join(*protobufs, "limit_bandwidth.proto")),
		"limitCpu":       usage.Parse(filepath.Join(*protobufs, "limit_cpu.proto")),
		"limitDisk":      usage.Parse(filepath.Join(*protobufs, "limit_disk.proto")),
		"limitMemory":    usage.Parse(filepath.Join(*protobufs, "limit_memory.proto")),
		"link":           usage.Parse(filepath.Join(*protobufs, "link.proto")),
		"list":           usage.Parse(filepath.Join(*protobufs, "list.proto")),
		"netIn":          usage.Parse(filepath.Join(*protobufs, "net_in.proto")),
		"netOut":         usage.Parse(filepath.Join(*protobufs, "net_out.proto")),
		"ping":           usage.Parse(filepath.Join(*protobufs, "ping.proto")),
		//"run": usage.Parse(filepath.Join(*protobufs, "run.proto")),
		"spawn":  usage.Parse(filepath.Join(*protobufs, "spawn.proto")),
		"stop":   usage.Parse(filepath.Join(*protobufs, "stop.proto")),
		"stream": usage.Parse(filepath.Join(*protobufs, "stream.proto")),
	}

	enc, err := json.Marshal(usage)
	if err != nil {
		panic(err)
	}

	fmt.Println("package main")
	fmt.Println("")
	fmt.Println("import \"encoding/json\"")
	fmt.Println("import \"github.com/vito/shank/usage\"")
	fmt.Println("")
	fmt.Println("var USAGE map[string]usage.Usage")
	fmt.Println("")
	fmt.Println("func init() {")
	fmt.Printf("\terr := json.Unmarshal([]byte(%#v), &USAGE)\n", string(enc))
	fmt.Println("\tif err != nil {")
	fmt.Println("\t\tpanic(err)")
	fmt.Println("\t}")
	fmt.Println("}")
}
