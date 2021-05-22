package tool

import (
	"log"
)

func init() {
	linkerd2, err := getLinkerd2()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, linkerd2)
}

func getLinkerd2() (linkerd2 Tool, err error) {
	owner := "linkerd"
	repo := "linkerd2"
	name := "linkerd2"
	linkerd2, err = getLatestReleaseFromGithub(owner, repo, name, "Ultralight, security-first service mesh",
		"darwin_amd64",
		"darwin_arm64",
		"linux_amd64",
		"linux_arm",
		"linux_arm64",
		"windows_amd64")
	linkerd2.GetVersions = func() (toolVersions []string, err error) {
		releases, err := getReleasesFromGithub(owner, repo, name)
		if err != nil {
			return nil, err
		}
		for _, release := range releases {
			toolVersions = append(toolVersions, *release.TagName)
		}
		return toolVersions, nil
	}
	linkerd2.GenerateVersion = func(version string) (toolVersion ToolVersion) {
		baseUrl := "https://github.com/linkerd/linkerd2/releases/download/" + version + "/linkerd2-cli-" + version
		alg := "sha256"
		return ToolVersion{
			Version: version,
			Platforms: map[string]ToolDownload{
				"darwin_amd64": {
					Url:      baseUrl + "-darwin",
					Checksum: baseUrl + "-darwin.sha256",
					Alg:      alg,
				},
				"darwin_arm64": {
					Url:      baseUrl + "-darwin-arm64",
					Checksum: baseUrl + "-darwin-arm64.sha256",
					Alg:      alg,
				},
				"linux_amd64": {
					Url:      baseUrl + "-linux-amd64",
					Checksum: baseUrl + "-linux-amd64.sha256",
					Alg:      alg,
				},
				"linux_arm": {
					Url:      baseUrl + "-linux-arm",
					Checksum: baseUrl + "-linux-arm.sha256",
					Alg:      alg,
				},
				"linux_arm64": {
					Url:      baseUrl + "-linux-arm64",
					Checksum: baseUrl + "-linux-arm64.sha256",
					Alg:      alg,
				},
				"windows_amd64": {
					Url:      baseUrl + "-windows.exe",
					Checksum: baseUrl + "-windows.exe.sha256",
					Alg:      alg,
				},
			},
		}
	}
	return linkerd2, err
}
