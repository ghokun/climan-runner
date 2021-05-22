package tool

import (
	"log"
)

func init() {
	plumber, err := getPlumber()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, plumber)
}

func getPlumber() (plumber Tool, err error) {
	owner := "batchcorp"
	repo := "plumber"
	name := "plumber"
	plumber, err = getLatestReleaseFromGithub(owner, repo, name, "Cli for messaging systems",
		"darwin_amd64",
		"linux_amd64",
		"windows_amd64")
	plumber.GetVersions = func() (toolVersions []string, err error) {
		releases, err := getReleasesFromGithub(owner, repo, name)
		if err != nil {
			return nil, err
		}
		for _, release := range releases {
			toolVersions = append(toolVersions, *release.TagName)
		}
		return toolVersions, nil
	}
	plumber.GenerateVersion = func(version string) (toolVersion ToolVersion) {
		baseUrl := "https://github.com/batchcorp/plumber/releases/download/" + version
		return ToolVersion{
			Version: version,
			Platforms: map[string]ToolDownload{
				"darwin_amd64": {
					Url: baseUrl + "/plumber-darwin",
				},
				"linux_amd64": {
					Url: baseUrl + "/plumber-linux",
				},
				"windows_amd64": {
					Url: baseUrl + "/plumber-windows.exe",
				},
			},
		}
	}
	return plumber, err
}
