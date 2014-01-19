package usage

import (
	"bufio"
	"os"
	"strings"
)

type Usage struct {
	Usage       string
	Description string
	Flags       map[string]string
}

func Parse(proto string) Usage {
	file, err := os.Open(proto)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)

	usage := readCommentLine(reader)

	description := readUntilRequests(reader)

	skipUntilRequestField(reader)

	fields := readRequestFields(reader)

	return Usage{
		Usage:       usage,
		Description: description,
		Flags:       fields,
	}
}

func readCommentLine(reader *bufio.Reader) string {
	start := make([]byte, 2)

	_, err := reader.Read(start)
	if err != nil {
		panic("expected start of comment; error: " + err.Error())
	}

	if string(start) != "//" {
		panic("expected start of comment; got " + string(start))
	}

	line, err := reader.ReadBytes('\n')
	if err != nil {
		panic("expected line of text; error: " + err.Error())
	}

	return strings.Trim(string(line), "\n ")
}

func readUntilRequests(reader *bufio.Reader) string {
	description := ""

	for {
		line := readCommentLine(reader)

		if line == "### Request" {
			break
		}

		if line == "" && description == "" {
			continue
		}

		description = description + "\n" + line
	}

	return description
}

func skipUntilRequestField(reader *bufio.Reader) {
	for {
		start, err := reader.Peek(4)
		if err != nil {
			panic("expected request field")
		}

		if string(start) == "// *" {
			break
		}

		if string(start) == "// >" {
			break
		}

		line := readCommentLine(reader)
		if line == "Empty." {
			break
		}
	}
}

func readRequestFields(reader *bufio.Reader) map[string]string {
	fields := make(map[string]string)

	for {
		start, err := reader.Peek(4)
		if err != nil {
			println("error checking for next field: " + err.Error())
			break
		}

		if string(start) != "// *" {
			break
		}

		name := readFieldName(reader)
		description := readFieldDescription(reader)

		fields[name] = description
	}

	return fields
}

func readFieldName(reader *bufio.Reader) string {
	start, err := reader.ReadBytes('`')
	if err != nil {
		panic("expected field name open; error: " + err.Error())
	}

	if string(start) != "// * `" {
		panic("expected field name open; got: " + string(start))
	}

	name, err := reader.ReadBytes('`')
	if err != nil {
		panic("expected field name close; error: " + err.Error())
	}

	return camelize(string(name[:len(name)-1]))
}

func readFieldDescription(reader *bufio.Reader) string {
	colon := make([]byte, 2)

	_, err := reader.Read(colon)
	if err != nil {
		panic("expected description start; error: " + err.Error())
	}

	if string(colon) != ": " {
		panic("expected colon and space; got: " + string(colon))
	}

	firstLine, err := reader.ReadBytes('\n')
	if err != nil {
		panic("expected first line of description; error: " + err.Error())
	}

	description := string(firstLine[:len(firstLine)-1])

	for {
		start, err := reader.Peek(4)
		if err != nil {
			println("error checking for next field in description: " + err.Error())
			break
		}

		if string(start) == "// *" {
			break
		}

		line := readCommentLine(reader)
		if line == "" {
			break
		}

		description = description + " " + line
	}

	return description
}

func camelize(str string) string {
	words := strings.Split(str, "_")

	camel := words[0]
	for _, word := range words[1:] {
		camel = camel + capitalize(word)
	}

	return camel
}

func capitalize(str string) string {
	return strings.ToUpper(str[:1]) + str[1:]
}
