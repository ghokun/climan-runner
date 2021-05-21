package tool

import (
	"log"
)

func init() {
	inletsctl, err := getInletsctl()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, inletsctl)
	generateToolSpecificFiles("inletsctl", inletsctl.Latest, getInletsctlVersions, generateInletsctlVersion)
}

func getInletsctl() (inletsctl Tool, err error) {
	return getLatestReleaseFromGithub("inlets", "inletsctl", "inletsctl", "The fastest way to create self-hosted exit-servers",
		"darwin_amd64",
		"linux_amd64",
		"linux_arm",
		"linux_arm64",
		"windows_amd64")
}

func getInletsctlVersions() (toolVersions []string, err error) {
	releases, err := getReleasesFromGithub("inlets", "inletsctl", "inletsctl")
	if err != nil {
		return nil, err
	}
	for _, release := range releases {
		toolVersions = append(toolVersions, *release.TagName)
	}
	return toolVersions, nil
}

func generateInletsctlVersion(version string) (toolVersion ToolVersion) {
	baseUrl := "https://github.com/inlets/inletsctl/releases/download/" + version
	return ToolVersion{
		Version: version,
		Platforms: map[string]ToolDownload{
			"darwin_amd64": {
				Url:      baseUrl + "/inletsctl-darwin.tgz",
				Checksum: baseUrl + "/inletsctl-darwin.sha256",
				Alg:      "sha256",
			},
			"linux_amd64": {
				Url:      baseUrl + "/inletsctl.tgz",
				Checksum: baseUrl + "/inletsctl.sha256",
				Alg:      "sha256",
			},
			"linux_arm": {
				Url:      baseUrl + "/inletsctl-armhf.tgz",
				Checksum: baseUrl + "/inletsctl-armhf.sha256",
				Alg:      "sha256",
			},
			"linux_arm64": {
				Url:      baseUrl + "/inletsctl-arm64.tgz",
				Checksum: baseUrl + "/inletsctl-arm64.sha256",
				Alg:      "sha256",
			},
			"windows_amd64": {
				Url:      baseUrl + "/inletsctl.exe.tgz",
				Checksum: baseUrl + "/inletsctl.exe.sha256",
				Alg:      "sha256",
			},
		},
	}
}
