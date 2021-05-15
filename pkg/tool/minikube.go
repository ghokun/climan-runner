package tool

import (
	"log"
)

func init() {
	minikube, err := getMinikube()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, minikube)
}

func getMinikube() (minikube Tool, err error) {
	return getLatestReleaseFromGithub("kubernetes", "minikube", "minikube", "Run Kubernetes locally", "darwin_amd64",
		"darwin_arm64",
		"linux_386",
		"linux_amd64",
		"linux_arm",
		"linux_arm64",
		"linux_ppc64le",
		"linux_s390x",
		"windows_amd64")
}
