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
}

func getK3d() (k3d Tool, err error) {
	owner := "rancher"
	repo := "k3d"
	name := "k3d"
	k3d, err = getLatestReleaseFromGithub(owner, repo, name, "k3s in Docker",
		"darwin_amd64",
		"darwin_arm64",
		"linux_386",
		"linux_amd64",
		"linux_arm",
		"linux_arm64",
		"windows_amd64")
	k3d.GetVersions = func() (toolVersions []string, err error) {
		releases, err := getReleasesFromGithub(owner, repo, name)
		if err != nil {
			return nil, err
		}
		for _, release := range releases {
			toolVersions = append(toolVersions, *release.TagName)
		}
		return toolVersions, nil
	}
	k3d.GenerateVersion = func(version string) (toolVersion ToolVersion) {
		baseUrl := "https://github.com/rancher/k3d/releases/download/" + version
		checksum := baseUrl + "/sha256sum.txt"
		alg := "sha256"
		return ToolVersion{
			Version: version,
			Platforms: map[string]ToolDownload{
				"darwin_amd64": {
					Url:      baseUrl + "/k3d-darwin-amd64",
					Checksum: checksum,
					Alg:      alg,
				},
				"darwin_arm64": {
					Url:      baseUrl + "/k3d-darwin-arm64",
					Checksum: checksum,
					Alg:      alg,
				},
				"linux_386": {
					Url:      baseUrl + "/k3d-linux-386",
					Checksum: checksum,
					Alg:      alg,
				},
				"linux_amd64": {
					Url:      baseUrl + "/k3d-linux-amd64",
					Checksum: checksum,
					Alg:      alg,
				},
				"linux_arm": {
					Url:      baseUrl + "/k3d-linux-arm",
					Checksum: checksum,
					Alg:      alg,
				},
				"linux_arm64": {
					Url:      baseUrl + "/k3d-linux-arm64",
					Checksum: checksum,
					Alg:      alg,
				},
				"windows_amd64": {
					Url:      baseUrl + "/k3d-windows-amd64.exe",
					Checksum: checksum,
					Alg:      alg,
				},
			},
		}
	}
	return k3d, err
}
