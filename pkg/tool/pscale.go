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
	generateToolSpecificFiles("pscale", pscale.Latest, getPscaleVersions, generatePscaleVersion)

}

func getPscale() (pscale Tool, err error) {
	return getLatestReleaseFromGithub("planetscale", "cli", "pscale", "Cli for PlanetScale Database ",
		"darwin_amd64",
		"linux_386",
		"linux_amd64",
		"linux_arm64",
		"windows_386",
		"windows_amd64")
}

func getPscaleVersions() (toolVersions []string, err error) {
	releases, err := getReleasesFromGithub("planetscale", "cli", "pscale")
	if err != nil {
		return nil, err
	}
	for _, release := range releases {
		toolVersions = append(toolVersions, *release.TagName)
	}
	return toolVersions, nil
}

func generatePscaleVersion(version string) (toolVersion ToolVersion) {
	withOuthV := strings.TrimPrefix(version, "v")
	baseUrl := "https://github.com/planetscale/cli/releases/download/" + version + "/pscale_" + withOuthV
	return ToolVersion{
		Version: version,
		Platforms: map[string]ToolDownload{
			"darwin_amd64": {
				Url:      baseUrl + "_macOS_amd64.tar.gz",
				Checksum: baseUrl + "_checksums.txt",
				Alg:      "sha256",
			},
			"linux_386": {
				Url:      baseUrl + "_linux_386.tar.gz",
				Checksum: baseUrl + "_checksums.txt",
				Alg:      "sha256",
			},
			"linux_amd64": {
				Url:      baseUrl + "_linux_amd64.tar.gz",
				Checksum: baseUrl + "_checksums.txt",
				Alg:      "sha256"},
			"linux_arm64": {
				Url:      baseUrl + "_linux_arm64.tar.gz",
				Checksum: baseUrl + "_checksums.txt",
				Alg:      "sha256"},
			"windows_386": {
				Url:      baseUrl + "_windows_386.zip",
				Checksum: baseUrl + "_checksums.txt",
				Alg:      "sha256"},
			"windows_amd64": {
				Url:      baseUrl + "_windows_amd64.zip",
				Checksum: baseUrl + "_checksums.txt",
				Alg:      "sha256"},
		},
	}
}
