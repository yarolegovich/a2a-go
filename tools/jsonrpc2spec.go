// Sanitizes proto-gener

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	inPath := flag.String("in", "./internal/jsonrpc/spec.go", "path to json-schema file")
	outPath := flag.String("out", "./a2a/types.go", "path to .go file with type definitions")
	flag.Parse()

	inFile, err := os.Open(*inPath)
	if err != nil {
		fmt.Printf("Failed to open input file: %v\n", err)
		os.Exit(1)
	}
	defer inFile.Close()

	converted, err := protoToStruct(inFile)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile(*outPath, []byte(converted), 0644); err != nil {
		fmt.Printf("Failed to write converted string to %s: %v\n", *outPath, err)
		os.Exit(1)
	}

	fmt.Printf("Successfully converted %s to %s\n", *inPath, *outPath)
}

func protoToStruct(in io.Reader) (string, error) {
	s := bufio.NewScanner(in)

	var result strings.Builder

	for s.Scan() {
		line := s.Text()

		line = handleReplacements(line)

		if strings.HasPrefix(line, "import ") {
			continue
		}

		if strings.HasPrefix(line, "func ") || strings.HasPrefix(line, "var ") {
			if err := skipPkgLevelBlock(s); err != nil {
				return "", err
			}
			continue
		}

		if !strings.HasPrefix(line, "type ") {
			writeLine(&result, line)
			continue
		}

		handleStruct(s, line, &result)
	}

	return result.String(), nil
}

func handleStruct(s *bufio.Scanner, line string, result *strings.Builder) error {
	if strings.Contains(line, " = ") {
		return skipPkgLevelBlock(s)
	}

	if !strings.HasSuffix(line, "{") {
		writeLine(result, line)
		return nil
	}

	writeLine(result, line)

	if strings.Contains(line, "interface{}") {
		return nil
	}

	text, err := consumeStruct(s)
	if err != nil {
		return err
	}

	result.WriteString(text)
	return nil
}

func consumeStruct(s *bufio.Scanner) (string, error) {
	var b strings.Builder

	for s.Scan() {
		line := s.Text()

		if shouldSkipStructLine(line) {
			continue
		}

		line = handleReplacements(line)

		if line == "}" {
			writeLine(&b, line)
			return b.String(), nil
		}

		if tagStart := strings.IndexRune(line, '`'); tagStart >= 0 {
			writeLine(&b, line[:tagStart-1])
			continue
		}

		writeLine(&b, line)
	}
	return "", fmt.Errorf("unexpected EOF while skipping func")
}

func shouldSkipStructLine(line string) bool {
	prefixBlacklist := []string{
		"// The version of the JSON-RPC protocol. MUST be exactly",
		"Jsonrpc string `json:",
		"// The method name. Must be ",
		"Method string `json:",
		"// The identifier for this request.",
		"// The identifier established by the client.",
		"Id interface{}",
	}
	trimmed := strings.TrimSpace(line)
	for _, p := range prefixBlacklist {
		if strings.HasPrefix(trimmed, p) {
			return true
		}
	}
	return false
}

func handleReplacements(line string) string {
	replace := map[string]string{
		"Url":               "URL",
		"interface{}":       "any",
		"JSON-RPC response": "response",
		"JSON-RPC request":  "request",
		"// UnmarshalJSON implements json.Unmarshaler.": "",
	}
	for old, new := range replace {
		if strings.Contains(line, old) {
			return strings.Replace(line, old, new, 1)
		}
	}
	return line
}

func skipPkgLevelBlock(s *bufio.Scanner) error {
	for s.Scan() {
		if line := s.Text(); line == "}" || line == ")" {
			return nil
		}
	}
	return fmt.Errorf("unexpected EOF while skipping pkg level block")
}

func writeLine(sb *strings.Builder, line string) {
	sb.WriteString(line)
	sb.WriteByte('\n')
}
