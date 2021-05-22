package tool

import (
	"log"
)

func init() {
	istioctl, err := getIstioctl()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, istioctl)
}

func getIstioctl() (istioctl Tool, err error) {
	owner := "istio"
	repo := "istio"
	name := "istioctl"
	istioctl, err = getLatestReleaseFromGithub(owner, repo, name, "Cli for Istio service mesh",
		"darwin_amd64",
		"darwin_arm64",
		"linux_amd64",
		"linux_arm",
		"linux_arm64",
		"windows_amd64")
	istioctl.GetVersions = func() (toolVersions []string, err error) {
		releases, err := getReleasesFromGithub(owner, repo, name)
		if err != nil {
			return nil, err
		}
		for _, release := range releases {
			toolVersions = append(toolVersions, *release.TagName)
		}
		return toolVersions, nil
	}
	istioctl.GenerateVersion = func(version string) (toolVersion ToolVersion) {
		baseUrl := "https://github.com/istio/istio/releases/download/" + version + "/istioctl-" + version
		alg := "sha256"
		return ToolVersion{
			Version: version,
			Platforms: map[string]ToolDownload{
				"darwin_amd64": {
					Url:      baseUrl + "-osx.tar.gz",
					Checksum: baseUrl + "-osx.tar.gz.sha256",
					Alg:      alg,
				},
				"darwin_arm64": {
					Url:      baseUrl + "-osx-arm64.tar.gz",
					Checksum: baseUrl + "-osx-arm64.tar.gz.sha256",
					Alg:      alg,
				},
				"linux_amd64": {
					Url:      baseUrl + "-linux-amd64.tar.gz",
					Checksum: baseUrl + "-linux-amd64.tar.gz.sha256",
					Alg:      alg,
				},
				"linux_arm": {
					Url:      baseUrl + "-linux-armv7.tar.gz",
					Checksum: baseUrl + "-linux-armv7.tar.gz.sha256",
					Alg:      alg,
				},
				"linux_arm64": {
					Url:      baseUrl + "-linux-arm64.tar.gz",
					Checksum: baseUrl + "-linux-arm64.tar.gz.sha256",
					Alg:      alg,
				},
				"windows_amd64": {
					Url:      baseUrl + "-win.zip",
					Checksum: baseUrl + "-win.zip.sha256",
					Alg:      alg,
				},
			},
		}
	}
	return istioctl, err
}
