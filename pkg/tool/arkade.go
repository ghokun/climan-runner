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
	generateToolSpecificFiles("arkade", arkade.Latest, getArkadeVersions, generateArkadeVersion)
}

func getArkade() (arkade Tool, err error) {
	return getLatestReleaseFromGithub("alexellis", "arkade", "arkade", "Open Source Kubernetes Marketplace",
		"darwin_amd64",
		"linux_amd64",
		"linux_arm",
		"linux_arm64",
		"windows_amd64")
}

func getArkadeVersions() (toolVersions []string, err error) {
	releases, err := getReleasesFromGithub("alexellis", "arkade", "arkade")
	if err != nil {
		return nil, err
	}
	for _, release := range releases {
		toolVersions = append(toolVersions, *release.TagName)
	}
	return toolVersions, nil
}

func generateArkadeVersion(version string) (toolVersion ToolVersion) {
	baseUrl := "https://github.com/alexellis/arkade/releases/download/" + version
	return ToolVersion{
		Version: version,
		Platforms: map[string]ToolDownload{
			"darwin_amd64": {
				Url:      baseUrl + "/arkade-darwin",
				Checksum: baseUrl + "/arkade-darwin.sha256",
				Alg:      "sha256",
			},
			"linux_amd64": {
				Url:      baseUrl + "/arkade",
				Checksum: baseUrl + "/arkade.sha256",
				Alg:      "sha256",
			},
			"linux_arm": {
				Url:      baseUrl + "/arkade-armhf",
				Checksum: baseUrl + "/arkade-armhf.sha256",
				Alg:      "sha256",
			},
			"linux_arm64": {
				Url:      baseUrl + "/arkade-arm64",
				Checksum: baseUrl + "/arkade-arm64.sha256",
				Alg:      "sha256",
			},
			"windows_amd64": {
				Url:      baseUrl + "/arkade.exe",
				Checksum: baseUrl + "/arkade.exe.sha256",
				Alg:      "sha256",
			},
		},
	}
}
