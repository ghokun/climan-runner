package tool

import (
	"log"
	"strings"
)

func init() {
	tkn, err := getTkn()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, tkn)
	generateToolSpecificFiles("tkn", tkn.Latest, getTknVersions, generateTknVersion)
}

func getTkn() (tkn Tool, err error) {
	return getLatestReleaseFromGithub("tektoncd", "cli", "tkn", "Cli for interacting with Tekton",
		"darwin_amd64",
		"linux_amd64",
		"linux_arm64",
		"linux_ppc64le",
		"linux_s390x",
		"windows_amd64")
}

func getTknVersions() (toolVersions []string, err error) {
	releases, err := getReleasesFromGithub("tektoncd", "cli", "tkn")
	if err != nil {
		return nil, err
	}
	for _, release := range releases {
		toolVersions = append(toolVersions, *release.TagName)
	}
	return toolVersions, nil
}

func generateTknVersion(version string) (toolVersion ToolVersion) {
	baseUrl := "https://github.com/tektoncd/cli/releases/download/" + version
	withOutV := strings.TrimPrefix(version, "v")

	return ToolVersion{
		Version: version,
		Platforms: map[string]ToolDownload{
			"darwin_amd64": {
				Url:      baseUrl + "/tkn_" + withOutV + "_Darwin_x86_64.tar.gz",
				Checksum: baseUrl + "/checksums.txt",
				Alg:      "sha256",
			},
			"linux_amd64": {
				Url:      baseUrl + "/tkn_" + withOutV + "_Linux_x86_64.tar.gz",
				Checksum: baseUrl + "/checksums.txt",
				Alg:      "sha256"},
			"linux_arm64": {
				Url:      baseUrl + "/tkn_" + withOutV + "_Linux_arm64.tar.gz",
				Checksum: baseUrl + "/checksums.txt",
				Alg:      "sha256"},
			"linux_ppc64le": {
				Url:      baseUrl + "/tkn_" + withOutV + "_Linux_ppc64le.tar.gz",
				Checksum: baseUrl + "/checksums.txt",
				Alg:      "sha256"},
			"linux_s390x": {
				Url:      baseUrl + "/tkn_" + withOutV + "_Linux_s390x.tar.gz",
				Checksum: baseUrl + "/checksums.txt",
				Alg:      "sha256"},
			"windows_amd64": {
				Url:      baseUrl + "/tkn_" + withOutV + "_Windows_x86_64.zip",
				Checksum: baseUrl + "/checksums.txt",
				Alg:      "sha256"},
		},
	}
}
