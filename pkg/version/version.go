package version

import (
	"regexp"
	"strconv"
	"strings"
)

type Version struct {
	Major int
	Minor int
	Patch int
	Other string
}

var versionRegexp = regexp.MustCompile(`(\d.\d.\d.\w+)|(\d.\d.\d)|(\d.\d)`)

func New(ver []byte) Version {
	match := versionRegexp.Find(ver)
	if match == nil {
		return Version{}
	}
	parts := strings.SplitN(string(match), ".", 4)
	v := Version{}
	v.Major = parseInt(parts[0])
	if len(parts) > 1 {
		v.Minor = parseInt(parts[1])
	}
	if len(parts) > 2 {
		v.Patch = parseInt(parts[2])
	}
	if len(parts) > 3 {
		v.Other = parts[3]
	}
	return v
}

func parseInt(buf string) int {
	val, _ := strconv.Atoi(buf)
	return val
}

func (v Version) Greater(other Version) bool {
	if v.Major > other.Major {
		return true
	} else if v.Major == other.Major && v.Minor > other.Minor {
		return true
	} else if v.Minor == other.Minor && v.Patch > other.Patch {
		return true
	}
	return false
}

func (v Version) String() string {
	parts := []string{
		strconv.Itoa(v.Major),
		strconv.Itoa(v.Minor),
		strconv.Itoa(v.Patch),
	}
	if v.Other != "" {
		parts = append(parts, v.Other)
	}
	return strings.Join(parts, ".")
}
