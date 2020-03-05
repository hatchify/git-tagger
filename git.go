package main

import "fmt"

func getCurrentTag() (tag *Tag, err error) {
	var str string
	if str, err = runCommand("git", "describe", "--tag"); err != nil {
		return
	}

	return newTag(str)
}

func pushTag(tag *Tag) (err error) {
	if _, err = runCommand("git", "tag", tag.String()); err != nil {
		return fmt.Errorf("error setting tag \"%s\":  %v", tag.String(), err)
	}

	if _, err = runCommand("git", "push", "origin", tag.String()); err != nil {
		return fmt.Errorf("error pushing tag \"%s\" to origin:  %v", tag.String(), err)
	}
	return
}
