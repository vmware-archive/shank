package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"path/filepath"

	"github.com/pivotal-cf-experimental/shank/usage"
)

var protobufs = flag.String(
	"protobufs",
	"./warden-protocol",
	"Directory containing the .proto files.",
)

func main() {
	flag.Parse()

	usage := map[string]usage.Usage{
		"create":         usage.Parse(filepath.Join(*protobufs, "create.proto")),
		"destroy":        usage.Parse(filepath.Join(*protobufs, "destroy.proto")),
		"echo":           usage.Parse(filepath.Join(*protobufs, "echo.proto")),
		"info":           usage.Parse(filepath.Join(*protobufs, "info.proto")),
		"limitBandwidth": usage.Parse(filepath.Join(*protobufs, "limit_bandwidth.proto")),
		"limitCpu":       usage.Parse(filepath.Join(*protobufs, "limit_cpu.proto")),
		"limitDisk":      usage.Parse(filepath.Join(*protobufs, "limit_disk.proto")),
		"limitMemory":    usage.Parse(filepath.Join(*protobufs, "limit_memory.proto")),
		"list":           usage.Parse(filepath.Join(*protobufs, "list.proto")),
		"netIn":          usage.Parse(filepath.Join(*protobufs, "net_in.proto")),
		"netOut":         usage.Parse(filepath.Join(*protobufs, "net_out.proto")),
		"ping":           usage.Parse(filepath.Join(*protobufs, "ping.proto")),
		"run":            usage.Parse(filepath.Join(*protobufs, "run.proto")),
		"attach":         usage.Parse(filepath.Join(*protobufs, "attach.proto")),
		"stop":           usage.Parse(filepath.Join(*protobufs, "stop.proto")),
		"capacity":       usage.Parse(filepath.Join(*protobufs, "capacity.proto")),
		"streamIn":       usage.Parse(filepath.Join(*protobufs, "stream_in.proto")),
		"streamOut":      usage.Parse(filepath.Join(*protobufs, "stream_out.proto")),
	}

	enc, err := json.Marshal(usage)
	if err != nil {
		panic(err)
	}

	fmt.Println("package main")
	fmt.Println("")
	fmt.Println("import \"encoding/json\"")
	fmt.Println("import \"github.com/pivotal-cf-experimental/shank/usage\"")
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
