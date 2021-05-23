package tool

import (
	"log"
)

func init() {
	kubevirt, err := getVirtctl()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, kubevirt)
}

func getVirtctl() (kubevirt Tool, err error) {
	owner := "kubevirt"
	repo := "kubevirt"
	name := "virtctl"
	kubevirt, err = getLatestReleaseFromGithub(owner, repo, name, "Cli for Kubevirt",
		"darwin_amd64",
		"linux_amd64",
		"windows_amd64")
	kubevirt.GetVersions = func() (toolVersions []string, err error) {
		releases, err := getReleasesFromGithub(owner, repo, name)
		if err != nil {
			return nil, err
		}
		for _, release := range releases {
			toolVersions = append(toolVersions, *release.TagName)
		}
		return toolVersions, nil
	}
	kubevirt.GenerateVersion = func(version string) (toolVersion ToolVersion) {
		baseUrl := "https://github.com/kubevirt/kubevirt/releases/download/" + version + "/virtctl-" + version
		return ToolVersion{
			Version: version,
			Platforms: map[string]ToolDownload{
				"darwin_amd64": {
					Url: baseUrl + "-darwin-amd64",
				},
				"linux_amd64": {
					Url: baseUrl + "-linux-amd64",
				},
				"windows_amd64": {
					Url: baseUrl + "-windows-amd64.exe",
				},
			},
		}
	}
	return kubevirt, err
}
