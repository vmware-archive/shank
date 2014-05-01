package main

import "encoding/json"
import "github.com/vito/shank/usage"

var USAGE map[string]usage.Usage

func init() {
	err := json.Unmarshal([]byte("{\"attach\":{\"Usage\":\"Attach to a process and start streaming its output\",\"Description\":\"\",\"Flags\":{\"handle\":\"Container handle.\",\"processId\":\"Process ID.\"}},\"capacity\":{\"Usage\":\"Returns the physical capacity of the server's machine.\",\"Description\":\"\",\"Flags\":{}},\"copyIn\":{\"Usage\":\"Copies files into a container.\",\"Description\":\"\\nFile permissions and symbolic links are be preserved, while hard links\\nare materialized. If the source path contains a trailing `/`, only the\\ncontents of the directory will be copied. Otherwise, the outermost\\ndirectory, along with its contents, will be copied. The unprivileged\\nuser inside the container is made owner of the resulting files.\\n\",\"Flags\":{\"dstPath\":\"Path in the container to copy to.\",\"handle\":\"Container handle.\",\"srcPath\":\"Path on the host to copy from.\"}},\"copyOut\":{\"Usage\":\"Copies files out of a container.\",\"Description\":\"\\nFile permissions and symbolic links are be preserved, while hard links\\nare materialized. If the source path contains a trailing `/`, only the\\ncontents of the directory will be copied. Otherwise, the outermost\\ndirectory, along with its contents, will be copied.\\n\\nBy default, the files on the host will be owned by root.\\nIf the `owner` field in the request is specified (in the form of `USER:GROUP`),\\nthe resulting files and directories will be owned by this user and group.\\n\",\"Flags\":{\"dstPath\":\"Path on the host to copy to.\",\"handle\":\"Container handle.\",\"srcPath\":\"Path in the container to copy from.\"}},\"create\":{\"Usage\":\"Creates a new container.\",\"Description\":\"\",\"Flags\":{\"bindMounts\":\"Contains the paths that should be mounted in the container's filesystem. The `src_path` field for every bind mount holds the path as seen from the host, where the `dst_path` field holds the path as seem from the container.\",\"graceTime\":\"Can be used to specify how long a container can go unreferenced by any client connection. After this time, the container will automatically be destroyed. If not specified, the container will be subject to the globally configured grace time.\",\"handle\":\"If specified, its value must be used to refer to the container in future requests. If it is not specified, warden uses its internal container ID as the container handle.\",\"properties\":\"A sequence of string key/value pairs providing arbitrary data about the container. The keys are assumed to be unique but this is not enforced via the protocol.\"}},\"destroy\":{\"Usage\":\"Destroys a container.\",\"Description\":\"\\nWhen a container is destroyed, its resource allocations are released,\\nits filesystem is removed, and all references to its handle are removed.\\n\\nAll resources that have been acquired during the lifetime of the container are released.\\nExamples of these resources are its subnet, its UID, and ports that were redirected to the container.\\n\\n\\u003e **TODO** Link to list of resources that can be acquired during the lifetime of a container.\\n\",\"Flags\":{\"handle\":\"Container handle.\"}},\"echo\":{\"Usage\":\"Echoes a message.\",\"Description\":\"\",\"Flags\":{\"message\":\"Message to echo.\"}},\"info\":{\"Usage\":\"Returns information about a container.\",\"Description\":\"\",\"Flags\":{\"handle\":\"Container handle.\"}},\"limitBandwidth\":{\"Usage\":\"Limits the network bandwidth for a container.\",\"Description\":\"\",\"Flags\":{}},\"limitCpu\":{\"Usage\":\"Limits the cpu shares for a container.\",\"Description\":\"\",\"Flags\":{\"handle\":\"Container handle.\",\"limitInShares\":\"New cpu limit in shares.\"}},\"limitDisk\":{\"Usage\":\"Limits the disk usage for a container.\",\"Description\":\"\\nThe disk limits that are set by this command only have effect for the container's unprivileged user.\\nFiles/directories created by its privileged user are not subject to these limits.\\n\\n\\u003e **TODO** Link to page explaining how disk management works.\\n\",\"Flags\":{\"blockHard\":\"New hard block limit.\",\"blockSoft\":\"New soft block limit.\",\"byteHard\":\"New hard block limit specified in bytes. Only has effect when `block_hard` is not specified.\",\"byteSoft\":\"New soft block limit specified in bytes. Only has effect when `block_soft` is not specified.\",\"handle\":\"Container handle.\",\"inodeHard\":\"New hard inode limit.\",\"inodeSoft\":\"New soft inode limit.\"}},\"limitMemory\":{\"Usage\":\"Limits the memory usage for a container.\",\"Description\":\"\\nThe limit applies to all process in the container. When the limit is\\nexceeded, the container will be automatically stopped.\\n\\nIf no limit is given, the current value is returned, and no change is made.\\n\",\"Flags\":{\"handle\":\"Container handle.\",\"limitInBytes\":\"New memory usage limit in bytes.\"}},\"list\":{\"Usage\":\"Lists all containers.\",\"Description\":\"\",\"Flags\":{\"properties\":\"List of properties to filter by. Multiple properties are ANDed together.\"}},\"netIn\":{\"Usage\":\"Map a port on the host to a port in the container.\",\"Description\":\"\\nIf a host port is not given, a port will be acquired from the server's port\\npool.\\n\\nIf a container port is not given, the port will be the same as the\\ncontainer port.\\n\\nThe two resulting ports are returned in the response.\\n\",\"Flags\":{\"containerPort\":\"Port on the container's interface that traffic should be forwarded to.\",\"handle\":\"Container handle.\",\"hostPort\":\"External port to be mapped.\"}},\"netOut\":{\"Usage\":\"Whitelist outbound network traffic.\",\"Description\":\"\\nIf the configuration directive `deny_networks` is not used,\\nall networks are already whitelisted and this command is effectively a no-op.\\n\",\"Flags\":{\"handle\":\"Container handle.\",\"network\":\"Network to whitelist (in the form `1.2.3.4/8`).\",\"port\":\"Port to whitelist.\"}},\"ping\":{\"Usage\":\"Pings the server.\",\"Description\":\"\",\"Flags\":{}},\"run\":{\"Usage\":\"Run a script inside a container.\",\"Description\":\"\\nThis request is equivalent to atomically spawning a process and immediately\\nattaching to it.\\n\",\"Flags\":{\"env\":\"Environment Variables (see `EnvironmentVariable`).\",\"handle\":\"Container handle.\",\"privileged\":\"Whether to run the script as root or not.\",\"rlimits\":\"Resource limits (see `ResourceLimits`).\",\"script\":\"Script to execute.\"}},\"stop\":{\"Usage\":\"Stops a container.\",\"Description\":\"\\nOnce a container is stopped, warden does not allow spawning new processes inside the container.\\nIt is possible to copy files in to and out of a stopped container.\\nIt is only when a container is destroyed that its filesystem is cleaned up.\\n\",\"Flags\":{\"background\":\"Return a response immediately instead of waiting for the container to be stopped.\",\"handle\":\"Container handle.\",\"kill\":\"Send SIGKILL instead of SIGTERM.\"}}}"), &USAGE)
	if err != nil {
		panic(err)
	}
}
