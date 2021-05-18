package tool

import (
	"log"
	"os"
	"path/filepath"
)

func init() {
	istioctl, err := getIstioctl()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, istioctl)
	// Generate istioctl specific directory
	folder := filepath.Join(".", "docs", istioctl.Name)
	os.Mkdir(folder, os.ModePerm)
	// Generate versions.json
	toolVersions, err := getIstioctlVersions()
	if err != nil {
		log.Fatal(err)
	}
	err = writeJson(folder, "versions.json", toolVersions)
	if err != nil {
		log.Fatal(err)
	}
	// Generate template.json
	template := generateIstioctlVersion("{{.Version}}")
	err = writeJson(folder, "template.json", template)
	if err != nil {
		log.Fatal(err)
	}
	// Generate latest.json
	latest := generateIstioctlVersion(istioctl.Latest)
	err = writeJson(folder, "latest.json", latest)
	if err != nil {
		log.Fatal(err)
	}
}

func getIstioctl() (istioctl Tool, err error) {
	return getLatestReleaseFromGithub("istio", "istio", "istioctl", "Cli for Istio service mesh",
		"darwin_amd64",
		"darwin_arm64",
		"linux_amd64",
		"linux_arm",
		"linux_arm64",
		"windows_amd64")
}

func getIstioctlVersions() (toolVersions []string, err error) {
	releases, err := getReleasesFromGithub("istio", "istio", "istioctl")
	if err != nil {
		return nil, err
	}
	for _, release := range releases {
		toolVersions = append(toolVersions, *release.TagName)
	}
	return toolVersions, nil
}

func generateIstioctlVersion(version string) (toolVersion ToolVersion) {
	baseUrl := "https://github.com/istio/istio/releases/download/" + version + "/istioctl-" + version
	return ToolVersion{
		Version: version,
		Platforms: map[string]ToolDownload{
			"darwin_amd64": {
				Url:      baseUrl + "-osx.tar.gz",
				Checksum: baseUrl + "-osx.tar.gz.sha256",
				Alg:      "sha256",
			},
			"darwin_arm64": {
				Url:      baseUrl + "-osx-arm64.tar.gz",
				Checksum: baseUrl + "-osx-arm64.tar.gz.sha256",
				Alg:      "sha256",
			},
			"linux_amd64": {
				Url:      baseUrl + "-linux-amd64.tar.gz",
				Checksum: baseUrl + "-linux-amd64.tar.gz.sha256",
				Alg:      "sha256",
			},
			"linux_arm": {
				Url:      baseUrl + "-linux-armv7.tar.gz",
				Checksum: baseUrl + "-linux-armv7.tar.gz.sha256",
				Alg:      "sha256",
			},
			"linux_arm64": {
				Url:      baseUrl + "-linux-arm64.tar.gz",
				Checksum: baseUrl + "-linux-arm64.tar.gz.sha256",
				Alg:      "sha256",
			},
			"windows_amd64": {
				Url:      baseUrl + "-win.zip",
				Checksum: baseUrl + "-win.zip.sha256",
				Alg:      "sha256",
			},
		},
	}
}
