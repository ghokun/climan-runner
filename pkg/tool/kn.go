package tool

import (
	"log"
)

func init() {
	kn, err := getKn()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, kn)
}

func getKn() (kn Tool, err error) {
	owner := "knative"
	repo := "client"
	name := "kn"
	kn, err = getLatestTagFromGithub(owner, repo, name, "Knative cli",
		"darwin_amd64",
		"linux_amd64",
		"linux_arm64",
		"linux_ppc64le",
		"linux_s390x",
		"windows_amd64")
	kn.GetVersions = func() (toolVersions []string, err error) {
		tags, err := getTagsFromGithub(owner, repo, name)
		if err != nil {
			return nil, err
		}
		for _, tag := range tags {
			toolVersions = append(toolVersions, *tag.Name)
		}
		return toolVersions, nil
	}
	kn.GenerateVersion = func(version string) (toolVersion ToolVersion) {
		baseUrl := "https://github.com/knative/client/releases/download/" + version
		checksum := baseUrl + "/checksums.txt"
		alg := "sha256"
		return ToolVersion{
			Version: version,
			Platforms: map[string]ToolDownload{
				"darwin_amd64": {
					Url:      baseUrl + "/kn-darwin-amd64",
					Checksum: checksum,
					Alg:      alg,
				},
				"linux_amd64": {
					Url:      baseUrl + "/kn-linux-amd64",
					Checksum: checksum,
					Alg:      alg,
				},
				"linux_arm64": {
					Url:      baseUrl + "/kn-linux-arm64",
					Checksum: checksum,
					Alg:      alg,
				},
				"linux_ppc64le": {
					Url:      baseUrl + "/kn-linux-ppc64le",
					Checksum: checksum,
					Alg:      alg,
				},
				"linux_s390x": {
					Url:      baseUrl + "/kn-linux-s390x",
					Checksum: checksum,
					Alg:      alg,
				},
				"windows_amd64": {
					Url:      baseUrl + "/kn-windows-amd64.exe",
					Checksum: checksum,
					Alg:      alg,
				},
			},
		}
	}
	return kn, err
}
