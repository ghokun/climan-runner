package tool

import (
	"log"
	"strings"
)

func init() {
	kamel, err := getKamel()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, kamel)
}

func getKamel() (kamel Tool, err error) {
	owner := "apache"
	repo := "camel-k"
	name := "kamel"
	kamel, err = getLatestReleaseFromGithub(owner, repo, name, "Cli for Apacke Camel-K",
		"darwin_amd64",
		"linux_amd64",
		"windows_amd64")
	kamel.GetVersions = func() (toolVersions []string, err error) {
		releases, err := getReleasesFromGithub(owner, repo, name)
		if err != nil {
			return nil, err
		}
		for _, release := range releases {
			toolVersions = append(toolVersions, *release.TagName)
		}
		return toolVersions, nil
	}
	kamel.GenerateVersion = func(version string) (toolVersion ToolVersion) {
		baseUrl := "https://github.com/apache/camel-k/releases/download/" + version
		withOutV := strings.TrimPrefix(version, "v")
		return ToolVersion{
			Version: version,
			Platforms: map[string]ToolDownload{
				"darwin_amd64": {
					Url: baseUrl + "/camel-k-client-" + withOutV + "-mac-64bit.tar.gz",
				},
				"linux_amd64": {
					Url: baseUrl + "/camel-k-client-" + withOutV + "-linux-64bit.tar.gz",
				},
				"windows_amd64": {
					Url: baseUrl + "/camel-k-client-" + withOutV + "-windows-64bit.tar.gz",
				},
			},
		}
	}
	return kamel, err
}
