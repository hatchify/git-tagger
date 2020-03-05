package main

import (
	"flag"
	"fmt"
	"log"
)

const (
	versionDescription = "Version type to increment by, supported values are: patch, minor, and major."
	actionDescription  = "Action to take place, supported values are: increment and get."
)

func main() {
	var (
		// Version type set by flag, defaults to "patch"
		versionType string
		// Action to take
		action string

		tag *Tag
		err error
	)

	flag.StringVar(&versionType, "type", "patch", versionDescription)
	flag.StringVar(&action, "action", "increment", actionDescription)
	flag.Parse()

	// Get current git tag as a parsed SEMVER tag
	if tag, err = getCurrentTag(); err != nil {
		log.Fatalf("error getting current tag: %v", err)
	}

	switch action {
	case "increment":
		increment(tag, versionType)
	case "get":
		print(tag)
	}
}

func print(tag *Tag) {
	fmt.Println(tag.String())
}

func increment(tag *Tag, versionType string) (err error) {
	// Increment SEMVER tag by provided version type
	if err = tag.Increment(versionType); err != nil {
		return fmt.Errorf("error incrementing tag: %v", err)
	}

	// Set new tag locally
	if err = setTag(tag); err != nil {
		return fmt.Errorf("error setting tag: %v", err)
	}

	// Push new tag to origin
	if err = pushTag(tag); err != nil {
		return fmt.Errorf("error updating tag: %v", err)
	}

	return
}
