package tool

import (
	"log"
)

func init() {
	arkade, err := getArkade()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, arkade)
}

func getArkade() (arkade Tool, err error) {
	owner := "alexellis"
	repo := "arkade"
	name := "arkade"
	arkade, err = getLatestReleaseFromGithub(owner, repo, name, "Open Source Kubernetes Marketplace",
		"darwin_amd64",
		"linux_amd64",
		"linux_arm",
		"linux_arm64",
		"windows_amd64")
	arkade.GetVersions = func() (toolVersions []string, err error) {
		releases, err := getReleasesFromGithub(owner, repo, name)
		if err != nil {
			return nil, err
		}
		for _, release := range releases {
			toolVersions = append(toolVersions, *release.TagName)
		}
		return toolVersions, nil
	}
	arkade.GenerateVersion = func(version string) (toolVersion ToolVersion) {
		baseUrl := "https://github.com/alexellis/arkade/releases/download/" + version
		alg := "sha256"
		return ToolVersion{
			Version: version,
			Platforms: map[string]ToolDownload{
				"darwin_amd64": {
					Url:      baseUrl + "/arkade-darwin",
					Checksum: baseUrl + "/arkade-darwin.sha256",
					Alg:      alg,
				},
				"linux_amd64": {
					Url:      baseUrl + "/arkade",
					Checksum: baseUrl + "/arkade.sha256",
					Alg:      alg,
				},
				"linux_arm": {
					Url:      baseUrl + "/arkade-armhf",
					Checksum: baseUrl + "/arkade-armhf.sha256",
					Alg:      alg,
				},
				"linux_arm64": {
					Url:      baseUrl + "/arkade-arm64",
					Checksum: baseUrl + "/arkade-arm64.sha256",
					Alg:      alg,
				},
				"windows_amd64": {
					Url:      baseUrl + "/arkade.exe",
					Checksum: baseUrl + "/arkade.exe.sha256",
					Alg:      alg,
				},
			},
		}
	}
	return arkade, err
}
