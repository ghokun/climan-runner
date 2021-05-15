package tool

import (
	"log"
)

func init() {
	kind, err := getKind()
	if err != nil {
		log.Fatal(err)
	}
	Tools[kind.Name] = kind
}

func getKind() (kind Tool, err error) {
	return getLatestReleaseFromGithub("kubernetes-sigs", "kind", "kind", "Kubernetes IN Docker - local clusters for testing Kubernetes", "darwin_amd64",
		"linux_amd64",
		"linux_arm64",
		"linux_ppc64le",
		"windows_amd64")
}
