package tool

import (
	"log"
)

func init() {
	helm, err := getHelm()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, helm)
	generateToolSpecificFiles("helm", helm.Latest, getHelmVersions, generateHelmVersion)
}

func getHelm() (helm Tool, err error) {
	return getLatestReleaseFromGithub("helm", "helm", "helm", "The Kubernetes Package Manager",
		"darwin_amd64",
		"linux_386",
		"linux_amd64",
		"linux_arm",
		"linux_arm64",
		"linux_ppc64le",
		"linux_s390x",
		"windows_amd64")
}

func getHelmVersions() (toolVersions []string, err error) {
	releases, err := getReleasesFromGithub("helm", "helm", "helm")
	if err != nil {
		return nil, err
	}
	for _, release := range releases {
		toolVersions = append(toolVersions, *release.TagName)
	}
	return toolVersions, nil
}

func generateHelmVersion(version string) (toolVersion ToolVersion) {
	baseUrl := "https://get.helm.sh/helm-" + version
	return ToolVersion{
		Version: version,
		Platforms: map[string]ToolDownload{
			"darwin_amd64": {
				Url:      baseUrl + "-darwin-amd64.tar.gz",
				Checksum: baseUrl + "-darwin-amd64.tar.gz.sha256sum",
				Alg:      "sha256",
			},
			"linux_386": {
				Url:      baseUrl + "/helm-linux-386",
				Checksum: baseUrl + "/sha256sum.txt",
				Alg:      "sha256",
			},
			"linux_amd64": {
				Url:      baseUrl + "-linux-amd64.tar.gz",
				Checksum: baseUrl + "-linux-amd64.tar.gz.sha256sum",
				Alg:      "sha256",
			},
			"linux_arm": {
				Url:      baseUrl + "-linux-arm.tar.gz",
				Checksum: baseUrl + "-linux-arm.tar.gz.sha256sum",
				Alg:      "sha256",
			},
			"linux_arm64": {
				Url:      baseUrl + "-linux-arm64.tar.gz",
				Checksum: baseUrl + "-linux-arm64.tar.gz.sha256sum",
				Alg:      "sha256",
			},
			"linux_ppc64le": {
				Url:      baseUrl + "-linux-ppc64le.tar.gz",
				Checksum: baseUrl + "-linux-ppc64le.tar.gz.sha256sum",
				Alg:      "sha256",
			},
			"linux_s390x": {
				Url:      baseUrl + "-linux-s390x.tar.gz",
				Checksum: baseUrl + "-linux-s390x.tar.gz.sha256sum",
				Alg:      "sha256",
			},
			"windows_amd64": {
				Url:      baseUrl + "-windows-amd64.zip",
				Checksum: baseUrl + "-windows-amd64.zip.sha256sum",
				Alg:      "sha256",
			},
		},
	}
}
