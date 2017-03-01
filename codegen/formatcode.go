package codegen

import (
	"bytes"
	"os/exec"
	"strings"
)

func FormatCode(code string) ([]byte, error) {
	fmtCmd := exec.Command("gofmt", "-s")
	fmtCmd.Stdin = strings.NewReader(code)
	var out bytes.Buffer
	fmtCmd.Stdout = &out
	if err := fmtCmd.Run(); err != nil {
		return []byte(code), err
	}

	if out.Len() == 0 {
		return []byte(code), nil
	}

	return out.Bytes(), nil
}
