package tool

import (
	"log"
)

func init() {
	argocd, err := getArgocd()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, argocd)
}

func getArgocd() (argocd Tool, err error) {
	return getLatestReleaseFromGithub("argoproj", "argo-cd", "argocd", "Declarative continuous deployment for Kubernetes", "darwin_amd64",
		"linux_amd64",
		"windows_amd64")
}
