package tool

import (
	"log"
)

func init() {
	odo, err := getOdo()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, odo)
}

func getOdo() (odo Tool, err error) {
	return getLatestReleaseFromGithub("openshift", "odo", "odo", "Developer-focused cli for OpenShift", "darwin_amd64",
		"linux_amd64",
		"linux_arm64",
		"linux_ppc64le",
		"linux_s390x",
		"windows_amd64")
}
