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
	generateToolSpecificFiles("odo", odo.Latest, getOdoVersions, generateOdoVersion)
}

func getOdo() (odo Tool, err error) {
	return getLatestReleaseFromGithub("openshift", "odo", "odo", "Developer-focused cli for OpenShift",
		"darwin_amd64",
		"linux_amd64",
		"linux_arm64",
		"linux_ppc64le",
		"linux_s390x",
		"windows_amd64")
}

func getOdoVersions() (toolVersions []string, err error) {
	releases, err := getReleasesFromGithub("openshift", "odo", "odo")
	if err != nil {
		return nil, err
	}
	for _, release := range releases {
		toolVersions = append(toolVersions, *release.TagName)
	}
	return toolVersions, nil
}

func generateOdoVersion(version string) (toolVersion ToolVersion) {
	baseUrl := "https://mirror.openshift.com/pub/openshift-v4/clients/odo/" + version
	return ToolVersion{
		Version: version,
		Platforms: map[string]ToolDownload{
			"darwin_amd64": {
				Url:      baseUrl + "/odo-darwin-amd64",
				Checksum: baseUrl + "/sha256sum.txt",
				Alg:      "sha256",
			},
			"linux_amd64": {
				Url:      baseUrl + "/odo-linux-amd64",
				Checksum: baseUrl + "/sha256sum.txt",
				Alg:      "sha256",
			},
			"linux_arm64": {
				Url:      baseUrl + "/odo-linux-arm64",
				Checksum: baseUrl + "/sha256sum.txt",
				Alg:      "sha256",
			},
			"linux_ppc64le": {
				Url:      baseUrl + "/odo-linux-ppc64le",
				Checksum: baseUrl + "/sha256sum.txt",
				Alg:      "sha256",
			},
			"linux_s390x": {
				Url:      baseUrl + "/odo-linux-s390x",
				Checksum: baseUrl + "/sha256sum.txt",
				Alg:      "sha256",
			},
			"windows_amd64": {
				Url:      baseUrl + "/odo-windows-amd64.exe",
				Checksum: baseUrl + "/sha256sum.txt",
				Alg:      "sha256",
			},
		},
	}
}
