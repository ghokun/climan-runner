package tool

import (
	"log"
	"strings"
)

func init() {
	pscale, err := getPscale()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, pscale)
}

func getPscale() (pscale Tool, err error) {
	owner := "planetscale"
	repo := "cli"
	name := "pscale"
	pscale, err = getLatestReleaseFromGithub(owner, repo, name, "Cli for PlanetScale Database ",
		"darwin_amd64",
		"linux_386",
		"linux_amd64",
		"linux_arm64",
		"windows_386",
		"windows_amd64")
	pscale.GetVersions = func() (toolVersions []string, err error) {
		releases, err := getReleasesFromGithub(owner, repo, name)
		if err != nil {
			return nil, err
		}
		for _, release := range releases {
			toolVersions = append(toolVersions, *release.TagName)
		}
		return toolVersions, nil
	}
	pscale.GenerateVersion = func(version string) (toolVersion ToolVersion) {
		withOuthV := strings.TrimPrefix(version, "v")
		baseUrl := "https://github.com/planetscale/cli/releases/download/" + version + "/pscale_" + withOuthV
		checksum := baseUrl + "_checksums.txt"
		alg := "sha256"
		return ToolVersion{
			Version: version,
			Platforms: map[string]ToolDownload{
				"darwin_amd64": {
					Url:      baseUrl + "_macOS_amd64.tar.gz",
					Checksum: checksum,
					Alg:      alg,
				},
				"linux_386": {
					Url:      baseUrl + "_linux_386.tar.gz",
					Checksum: checksum,
					Alg:      alg,
				},
				"linux_amd64": {
					Url:      baseUrl + "_linux_amd64.tar.gz",
					Checksum: checksum,
					Alg:      alg},
				"linux_arm64": {
					Url:      baseUrl + "_linux_arm64.tar.gz",
					Checksum: checksum,
					Alg:      alg},
				"windows_386": {
					Url:      baseUrl + "_windows_386.zip",
					Checksum: checksum,
					Alg:      alg},
				"windows_amd64": {
					Url:      baseUrl + "_windows_amd64.zip",
					Checksum: checksum,
					Alg:      alg},
			},
		}
	}
	return pscale, err
}
