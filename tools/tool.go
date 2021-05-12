package tools

import (
	"github.com/google/go-github/v34/github"
	"golang.org/x/mod/semver"
)

var (
	client = github.NewClient(nil) // github client
	Tools  = map[string]Tool{}     // holds supported tools for current platform
)

// Tool struct is used for command line tool definition and operations.
type Tool struct {
	Name           string                                        // Name of the tool, kubectl, helm, etc..
	Description    string                                        // Description of tool
	Support        int                                           // Supported OS and architectures
	GetLatest      func() (latest ToolVersion, err error)        // Get latest version of tool from remote source
	GetAll         func() (all ToolVersions, err error)          // Get all versions of tool from remote source
	GetApiVersion  func() (version string, err error)            // Get Api version of tool
	VerifyChecksum func(version string) (result bool, err error) // Verify checksum of downloaded tool
}

// ToolVersion struct holds version and binary download url for any tool.
type ToolVersion struct {
	Version string `json:"version,omitempty"`
	URL     string `json:"url,omitempty"`
}

type ToolVersions []ToolVersion

func (t ToolVersions) Len() int {
	return len(t)
}

func (t ToolVersions) Less(i, j int) bool {
	return semver.Compare(t[i].Version, t[j].Version) == 1
}

func (t ToolVersions) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

type Cache struct {
	Timestamp string            `json:"timestamp,omitempty"`
	Data      map[string]string `json:"data,omitempty"`
}
