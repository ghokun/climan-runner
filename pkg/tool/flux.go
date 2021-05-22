package tool

import (
	"log"
	"strings"
)

func init() {
	flux, err := getFlux()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, flux)
}

func getFlux() (flux Tool, err error) {
	owner := "fluxcd"
	repo := "flux2"
	name := "flux"
	flux, err = getLatestReleaseFromGithub(owner, repo, name, "Cli for Flux",
		"darwin_amd64",
		"darwin_arm64",
		"linux_amd64",
		"linux_arm",
		"linux_arm64",
		"windows_386",
		"windows_amd64")
	flux.GetVersions = func() (toolVersions []string, err error) {
		releases, err := getReleasesFromGithub(owner, repo, name)
		if err != nil {
			return nil, err
		}
		for _, release := range releases {
			toolVersions = append(toolVersions, *release.TagName)
		}
		return toolVersions, nil
	}
	flux.GenerateVersion = func(version string) (toolVersion ToolVersion) {
		withOutV := strings.TrimPrefix(version, "v")
		baseUrl := "https://github.com/fluxcd/flux2/releases/download/" + version + "/flux_" + withOutV
		checksum := baseUrl + "_checksums.txt"
		alg := "sha256"
		return ToolVersion{
			Version: version,
			Platforms: map[string]ToolDownload{
				"darwin_amd64": {
					Url:      baseUrl + "_darwin_amd64.tar.gz",
					Checksum: checksum,
					Alg:      alg,
				},
				"darwin_arm64": {
					Url:      baseUrl + "_darwin_arm64.tar.gz",
					Checksum: checksum,
					Alg:      alg,
				},
				"linux_amd64": {
					Url:      baseUrl + "_linux_amd64.tar.gz",
					Checksum: checksum,
					Alg:      alg,
				},
				"linux_arm": {
					Url:      baseUrl + "_linux_arm.tar.gz",
					Checksum: checksum,
					Alg:      alg,
				},
				"linux_arm64": {
					Url:      baseUrl + "_linux_arm64.tar.gz",
					Checksum: checksum,
					Alg:      alg,
				},
				"windows_amd64": {
					Url:      baseUrl + "_windows_amd64.tar.gz",
					Checksum: checksum,
					Alg:      alg,
				},
				"windows_386": {
					Url:      baseUrl + "_windows_386.tar.gz",
					Checksum: checksum,
					Alg:      alg,
				},
			},
		}
	}
	return flux, err
}
