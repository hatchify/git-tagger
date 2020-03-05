package main

import (
	"fmt"
	"strings"
)

func newTag(str string) (tp *Tag, err error) {
	var t Tag
	spl := strings.Split(str, ".")
	if len(spl) != 3 {
		err = fmt.Errorf("invalid number of tag parts, expected 3 and received %d", len(spl))
		return
	}

	if t.Major, err = newMajorVersion(spl[0]); err != nil {
		err = fmt.Errorf("error parsing major version: %v", err)
		return
	}

	if t.Minor, err = newMinorVersion(spl[1]); err != nil {
		err = fmt.Errorf("error parsing minor version: %v", err)
		return
	}

	if t.Patch, err = newPatchVersion(spl[2]); err != nil {
		err = fmt.Errorf("error parsing patch version: %v", err)
		return
	}

	tp = &t
	return
}

// Tag represents a SEMVER tag
type Tag struct {
	Major uint64
	Minor uint64
	Patch uint64
}

func (t *Tag) String() string {
	return fmt.Sprintf("v%d.%d.%d", t.Major, t.Minor, t.Patch)
}

// Increment will increment the SEMVER tag by 1 for a given version type
func (t *Tag) Increment(versionType string) (err error) {
	switch versionType {
	case "patch":
		t.Patch++
	case "minor":
		t.Minor++
		t.Patch = 0
	case "major":
		t.Major++
		t.Minor = 0
		t.Patch = 0

	default:
		return fmt.Errorf("invalid version type provided: %s", versionType)
	}

	return
}
