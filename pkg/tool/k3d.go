package tool

import (
	"log"
)

func init() {
	k3d, err := getK3d()
	if err != nil {
		log.Fatal(err)
	}
	Tools[k3d.Name] = k3d
}

func getK3d() (k3d Tool, err error) {
	return getLatestReleaseFromGithub("rancher", "k3d", "k3d", "Little helper to run Rancher Lab's k3s in Docker", "darwin_amd64",
		"darwin_arm64",
		"linux_386",
		"linux_amd64",
		"linux_arm",
		"linux_arm64",
		"windows_amd64")
}
