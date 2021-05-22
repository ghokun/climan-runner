package tool

import (
	"log"
)

func init() {
	k3sup, err := getK3sup()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, k3sup)
}

func getK3sup() (k3sup Tool, err error) {
	owner := "alexellis"
	repo := "k3sup"
	name := "k3sup"
	k3sup, err = getLatestReleaseFromGithub(owner, repo, name, "Bootstrap Kubernetes with k3s",
		"darwin_amd64",
		"linux_amd64",
		"linux_arm",
		"linux_arm64",
		"windows_amd64")
	k3sup.GetVersions = func() (toolVersions []string, err error) {
		releases, err := getReleasesFromGithub(owner, repo, name)
		if err != nil {
			return nil, err
		}
		for _, release := range releases {
			toolVersions = append(toolVersions, *release.TagName)
		}
		return toolVersions, nil
	}
	k3sup.GenerateVersion = func(version string) (toolVersion ToolVersion) {
		baseUrl := "https://github.com/alexellis/k3sup/releases/download/" + version
		alg := "sha256"
		return ToolVersion{
			Version: version,
			Platforms: map[string]ToolDownload{
				"darwin_amd64": {
					Url:      baseUrl + "/k3sup-darwin",
					Checksum: baseUrl + "/k3sup-darwin.sha256",
					Alg:      alg,
				},
				"linux_amd64": {
					Url:      baseUrl + "/k3sup",
					Checksum: baseUrl + "/k3sup.sha256",
					Alg:      alg,
				},
				"linux_arm": {
					Url:      baseUrl + "/k3sup-armhf",
					Checksum: baseUrl + "/k3sup-armhf.sha256",
					Alg:      alg,
				},
				"linux_arm64": {
					Url:      baseUrl + "/k3sup-arm64",
					Checksum: baseUrl + "/k3sup-arm64.sha256",
					Alg:      alg,
				},
				"windows_amd64": {
					Url:      baseUrl + "/k3sup.exe",
					Checksum: baseUrl + "/k3sup.exe.sha256",
					Alg:      alg,
				},
			},
		}
	}
	return k3sup, err
}
