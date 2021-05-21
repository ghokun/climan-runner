package tool

import (
	"log"
)

func init() {
	k3d, err := getK3d()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, k3d)
	generateToolSpecificFiles("k3d", k3d.Latest, getK3dVersions, generateK3dVersion)
}

func getK3d() (k3d Tool, err error) {
	return getLatestReleaseFromGithub("rancher", "k3d", "k3d", "k3s in Docker",
		"darwin_amd64",
		"darwin_arm64",
		"linux_386",
		"linux_amd64",
		"linux_arm",
		"linux_arm64",
		"windows_amd64")
}

func getK3dVersions() (toolVersions []string, err error) {
	releases, err := getReleasesFromGithub("rancher", "k3d", "k3d")
	if err != nil {
		return nil, err
	}
	for _, release := range releases {
		toolVersions = append(toolVersions, *release.TagName)
	}
	return toolVersions, nil
}

func generateK3dVersion(version string) (toolVersion ToolVersion) {
	baseUrl := "https://github.com/rancher/k3d/releases/download/" + version
	return ToolVersion{
		Version: version,
		Platforms: map[string]ToolDownload{
			"darwin_amd64": {
				Url:      baseUrl + "/k3d-darwin-amd64",
				Checksum: baseUrl + "/sha256sum.txt",
				Alg:      "sha256",
			},
			"darwin_arm64": {
				Url:      baseUrl + "/k3d-darwin-arm64",
				Checksum: baseUrl + "/sha256sum.txt",
				Alg:      "sha256",
			},
			"linux_386": {
				Url:      baseUrl + "/k3d-linux-386",
				Checksum: baseUrl + "/sha256sum.txt",
				Alg:      "sha256",
			},
			"linux_amd64": {
				Url:      baseUrl + "/k3d-linux-amd64",
				Checksum: baseUrl + "/sha256sum.txt",
				Alg:      "sha256",
			},
			"linux_arm": {
				Url:      baseUrl + "/k3d-linux-arm",
				Checksum: baseUrl + "/sha256sum.txt",
				Alg:      "sha256",
			},
			"linux_arm64": {
				Url:      baseUrl + "/k3d-linux-arm64",
				Checksum: baseUrl + "/sha256sum.txt",
				Alg:      "sha256",
			},
			"windows_amd64": {
				Url:      baseUrl + "/k3d-windows-amd64.exe",
				Checksum: baseUrl + "/sha256sum.txt",
				Alg:      "sha256",
			},
		},
	}
}
