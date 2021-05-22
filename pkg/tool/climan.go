package tool

import (
	"log"
)

func init() {
	climan, err := getCliman()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, climan)
}

func getCliman() (climan Tool, err error) {
	owner := "ghokun"
	repo := "climan"
	name := "climan"
	climan, err = getLatestReleaseFromGithub(owner, repo, name, "Cloud tools cli manager",
		"darwin_amd64",
		"darwin_arm64",
		"linux_386",
		"linux_amd64",
		"linux_arm",
		"linux_arm64",
		"linux_ppc64le",
		"linux_s390x",
		"windows_386",
		"windows_amd64",
		"windows_arm")
	climan.GetVersions = func() (toolVersions []string, err error) {
		releases, err := getReleasesFromGithub(owner, repo, name)
		if err != nil {
			return nil, err
		}
		for _, release := range releases {
			toolVersions = append(toolVersions, *release.TagName)
		}
		return toolVersions, nil
	}
	climan.GenerateVersion = func(version string) (toolVersion ToolVersion) {
		baseUrl := "https://github.com/ghokun/climan/releases/download/" + version
		checksum := baseUrl + "/climan-checksums.txt"
		alg := "sha256"
		return ToolVersion{
			Version: version,
			Platforms: map[string]ToolDownload{
				"darwin_amd64": {
					Url:      baseUrl + "/climan-darwin-amd64",
					Checksum: checksum,
					Alg:      alg,
				},
				"darwin_arm64": {
					Url:      baseUrl + "/climan-darwin-arm64",
					Checksum: checksum,
					Alg:      alg,
				},
				"linux_386": {
					Url:      baseUrl + "/climan-linux-386",
					Checksum: checksum,
					Alg:      alg,
				},
				"linux_amd64": {
					Url:      baseUrl + "/climan-linux-amd64",
					Checksum: checksum,
					Alg:      alg,
				},
				"linux_arm": {
					Url:      baseUrl + "/climan-linux-arm",
					Checksum: checksum,
					Alg:      alg,
				},
				"linux_arm64": {
					Url:      baseUrl + "/climan-linux-arm64",
					Checksum: checksum,
					Alg:      alg,
				},
				"linux_ppc64le": {
					Url:      baseUrl + "/climan-linux-ppc64le",
					Checksum: checksum,
					Alg:      alg,
				},
				"linux_s390x": {
					Url:      baseUrl + "/climan-linux-s390x",
					Checksum: checksum,
					Alg:      alg,
				},
				"windows_386": {
					Url:      baseUrl + "/climan-windows-386.exe",
					Checksum: checksum,
					Alg:      alg,
				},
				"windows_amd64": {
					Url:      baseUrl + "/climan-windows-amd64.exe",
					Checksum: checksum,
					Alg:      alg,
				},
				"windows_arm": {
					Url:      baseUrl + "/climan-windows-arm.exe",
					Checksum: checksum,
					Alg:      alg,
				},
			},
		}
	}
	return climan, err
}
