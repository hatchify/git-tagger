package main

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func newMajorVersion(str string) (major uint64, err error) {
	if str[0] != 'v' {
		err = fmt.Errorf("expected part to start with \"v\" and received \"%c\"", str[0])
		return
	}

	return strconv.ParseUint(str[1:], 10, 64)
}

func newMinorVersion(str string) (minor uint64, err error) {
	return strconv.ParseUint(str, 10, 64)
}

func newPatchVersion(str string) (patch uint64, err error) {
	str = strings.SplitN(str, "-", 2)[0]
	str = strings.Replace(str, "\n", "", -1)
	return strconv.ParseUint(str, 10, 64)
}

func runCommand(name string, args ...string) (stdout, stderr string, err error) {
	// Initialize buffers
	buf := bytes.NewBuffer(nil)
	errBuf := bytes.NewBuffer(nil)

	// Create and initialize command
	cmd := exec.Command(name, args...)
	cmd.Stdout = buf
	cmd.Stderr = errBuf

	// Run command
	if err = cmd.Run(); err != nil {
		err = errors.New(errBuf.String())
		return
	}

	// Return buffer string types
	stdout = buf.String()
	stderr = errBuf.String()
	return
}
