package tool

import (
	"log"
)

func init() {
	kn, err := getKn()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, kn)
}

func getKn() (kn Tool, err error) {
	return getLatestTagFromGithub("knative", "client", "kn", "Knative cli",
		"darwin_amd64",
		"linux_amd64",
		"linux_arm64",
		"linux_ppc64le",
		"linux_s390x",
		"windows_amd64")
}
