package tool

import (
	"log"
)

func init() {
	tkn, err := getTkn()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, tkn)
}

func getTkn() (tkn Tool, err error) {
	return getLatestReleaseFromGithub("tektoncd", "cli", "tkn", "Cli for interacting with Tekton", "darwin_amd64",
		"linux_amd64",
		"linux_arm64",
		"linux_ppc64le",
		"linux_s390x",
		"windows_amd64")
}
