package tool

import (
	"log"
)

func init() {
	helm, err := getHelm()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, helm)
}

func getHelm() (helm Tool, err error) {
	return getLatestReleaseFromGithub("helm", "helm", "helm", "The Kubernetes Package Manager", "darwin_amd64",
		"darwin_arm64",
		"linux_386",
		"linux_amd64",
		"linux_arm",
		"linux_arm64",
		"linux_ppc64le",
		"linux_s390x",
		"windows_amd64")
}
