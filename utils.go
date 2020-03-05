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

func runCommand(name string, args ...string) (output string, err error) {
	// Initialize buffers
	buf := bytes.NewBuffer(nil)
	errBuf := bytes.NewBuffer(nil)

	// Create and initialize command
	cmd := exec.Command(name, args...)
	cmd.Stdout = buf
	cmd.Stderr = errBuf

	// Run command
	if err = cmd.Run(); err != nil {
		return
	}

	// Check to see if anything was pushed to Stderr
	if errBuf.Len() > 0 {
		// Error buffer was populated, set and return error
		err = errors.New(errBuf.String())
		return
	}

	// Set output as buffer converted to a string
	output = buf.String()
	return
}
