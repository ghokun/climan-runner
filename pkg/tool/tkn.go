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
}

func getTkn() (tkn Tool, err error) {
	owner := "tektoncd"
	repo := "cli"
	name := "tkn"
	tkn, err = getLatestReleaseFromGithub(owner, repo, name, "Cli for interacting with Tekton",
		"darwin_amd64",
		"linux_amd64",
		"linux_arm64",
		"linux_ppc64le",
		"linux_s390x",
		"windows_amd64")
	tkn.GetVersions = func() (toolVersions []string, err error) {
		releases, err := getReleasesFromGithub(owner, repo, name)
		if err != nil {
			return nil, err
		}
		for _, release := range releases {
			toolVersions = append(toolVersions, *release.TagName)
		}
		return toolVersions, nil
	}
	tkn.GenerateVersion = func(version string) (toolVersion ToolVersion) {
		baseUrl := "https://github.com/tektoncd/cli/releases/download/" + version
		withOutV := strings.TrimPrefix(version, "v")
		checksum := baseUrl + "/checksums.txt"
		alg := "sha256"
		return ToolVersion{
			Version: version,
			Platforms: map[string]ToolDownload{
				"darwin_amd64": {
					Url:      baseUrl + "/tkn_" + withOutV + "_Darwin_x86_64.tar.gz",
					Checksum: checksum,
					Alg:      alg,
				},
				"linux_amd64": {
					Url:      baseUrl + "/tkn_" + withOutV + "_Linux_x86_64.tar.gz",
					Checksum: checksum,
					Alg:      alg,
				},
				"linux_arm64": {
					Url:      baseUrl + "/tkn_" + withOutV + "_Linux_arm64.tar.gz",
					Checksum: checksum,
					Alg:      alg,
				},
				"linux_ppc64le": {
					Url:      baseUrl + "/tkn_" + withOutV + "_Linux_ppc64le.tar.gz",
					Checksum: checksum,
					Alg:      alg,
				},
				"linux_s390x": {
					Url:      baseUrl + "/tkn_" + withOutV + "_Linux_s390x.tar.gz",
					Checksum: checksum,
					Alg:      alg,
				},
				"windows_amd64": {
					Url:      baseUrl + "/tkn_" + withOutV + "_Windows_x86_64.zip",
					Checksum: checksum,
					Alg:      alg,
				},
			},
		}
	}
	return tkn, err
}
