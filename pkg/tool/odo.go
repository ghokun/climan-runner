package tool

import (
	"log"
)

func init() {
	odo, err := getOdo()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, odo)
}

func getOdo() (odo Tool, err error) {
	owner := "openshift"
	repo := "odo"
	name := "odo"
	odo, err = getLatestReleaseFromGithub(owner, repo, name, "Developer-focused cli for OpenShift",
		"darwin_amd64",
		"linux_amd64",
		"linux_arm64",
		"linux_ppc64le",
		"linux_s390x",
		"windows_amd64")
	odo.GetVersions = func() (toolVersions []string, err error) {
		releases, err := getReleasesFromGithub(owner, repo, name)
		if err != nil {
			return nil, err
		}
		for _, release := range releases {
			toolVersions = append(toolVersions, *release.TagName)
		}
		return toolVersions, nil
	}
	odo.GenerateVersion = func(version string) (toolVersion ToolVersion) {
		baseUrl := "https://mirror.openshift.com/pub/openshift-v4/clients/odo/" + version
		checksum := baseUrl + "/sha256sum.txt"
		alg := "sha256"
		return ToolVersion{
			Version: version,
			Platforms: map[string]ToolDownload{
				"darwin_amd64": {
					Url:      baseUrl + "/odo-darwin-amd64",
					Checksum: checksum,
					Alg:      alg,
				},
				"linux_amd64": {
					Url:      baseUrl + "/odo-linux-amd64",
					Checksum: checksum,
					Alg:      alg,
				},
				"linux_arm64": {
					Url:      baseUrl + "/odo-linux-arm64",
					Checksum: checksum,
					Alg:      alg,
				},
				"linux_ppc64le": {
					Url:      baseUrl + "/odo-linux-ppc64le",
					Checksum: checksum,
					Alg:      alg,
				},
				"linux_s390x": {
					Url:      baseUrl + "/odo-linux-s390x",
					Checksum: checksum,
					Alg:      alg,
				},
				"windows_amd64": {
					Url:      baseUrl + "/odo-windows-amd64.exe",
					Checksum: checksum,
					Alg:      alg,
				},
			},
		}
	}
	return odo, err
}
