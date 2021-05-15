package tool

import (
	"log"
)

func init() {
	arkade, err := getArkade()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, arkade)
}

func getArkade() (arkade Tool, err error) {
	return getLatestReleaseFromGithub("alexellis", "arkade", "arkade", "Open Source Kubernetes Marketplace", "darwin_amd64",
		"linux_amd64",
		"linux_arm",
		"linux_arm64",
		"windows_amd64")
}
