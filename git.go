package main

import (
	"errors"
	"fmt"
)

func getCurrentTag() (tag *Tag, err error) {
	var stdout, stderr string
	if stdout, stderr, err = runCommand("git", "describe", "--tag"); err != nil {
		return
	}

	if len(stderr) > 0 {
		err = errors.New(stderr)
		return
	}

	return newTag(stdout)
}

func setTag(tag *Tag) (err error) {
	var stderr string
	fmt.Println("Setting tag", tag.String())
	if _, stderr, err = runCommand("git", "tag", tag.String()); err != nil {
		return
	}

	if len(stderr) > 0 {
		return errors.New(stderr)
	}

	return
}

func pushTag(tag *Tag) (err error) {
	var stdout string
	// For some reason, git push pipes the output to stderr. Because of this, we set the second
	// argument as the stdout
	if _, stdout, err = runCommand("git", "push", "origin", tag.String()); err != nil {
		return fmt.Errorf("error pushing tag \"%s\" to origin:  %v", tag.String(), err)
	}

	fmt.Print(stdout)
	return
}
