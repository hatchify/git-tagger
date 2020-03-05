package main

import (
	"flag"
	"log"
)

const (
	flagName        = "type"
	flagDefault     = "patch"
	flagDescription = "Version type to increment by, supported values are: patch, minor, and major."
)

func main() {
	var (
		// Version type set by flag, defaults to "patch"
		versionType string

		tag *Tag
		err error
	)

	flag.StringVar(&versionType, flagName, flagDefault, flagDescription)
	flag.Parse()

	// Get current git tag as a parsed SEMVER tag
	if tag, err = getCurrentTag(); err != nil {
		log.Fatalf("error getting current tag: %v", err)

	}

	// Increment SEMVER tag by provided version type
	if err = tag.Increment(versionType); err != nil {
		log.Fatalf("error incrementing tag: %v", err)
	}

	// Push new tag locally and to origin
	if err = pushTag(tag); err != nil {
		log.Fatalf("error updating tag: %v", err)
	}
}
