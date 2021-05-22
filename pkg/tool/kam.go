package tool

import (
	"log"
)

func init() {
	kam, err := getKam()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, kam)
}

func getKam() (kam Tool, err error) {
	owner := "redhat-developer"
	repo := "kam"
	name := "kam"
	kam, err = getLatestReleaseFromGithub(owner, repo, name, "GitOps Application Manager",
		"darwin_amd64",
		"linux_amd64",
		"windows_amd64")
	kam.GetVersions = func() (toolVersions []string, err error) {
		releases, err := getReleasesFromGithub(owner, repo, name)
		if err != nil {
			return nil, err
		}
		for _, release := range releases {
			toolVersions = append(toolVersions, *release.TagName)
		}
		return toolVersions, nil
	}
	kam.GenerateVersion = func(version string) (toolVersion ToolVersion) {
		baseUrl := "https://github.com/redhat-developer/kam/releases/download/" + version
		return ToolVersion{
			Version: version,
			Platforms: map[string]ToolDownload{
				"darwin_amd64": {
					Url: baseUrl + "/kam_darwin_amd64",
				},
				"linux_amd64": {
					Url: baseUrl + "/kam_linux_amd64",
				},
				"windows_amd64": {
					Url: baseUrl + "/kam_windows_amd64.exe",
				},
			},
		}
	}
	return kam, err
}
