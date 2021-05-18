package tool

import (
	"log"
	"os"
	"path/filepath"
)

func init() {
	k3sup, err := getK3sup()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, k3sup)
	// Generate k3sup specific directory
	folder := filepath.Join(".", "docs", k3sup.Name)
	os.Mkdir(folder, os.ModePerm)
	// Generate versions.json
	toolVersions, err := getK3supVersions()
	if err != nil {
		log.Fatal(err)
	}
	err = writeJson(folder, "versions.json", toolVersions)
	if err != nil {
		log.Fatal(err)
	}
	// Generate template.json
	template := generateK3supVersion("{{.Version}}")
	err = writeJson(folder, "template.json", template)
	if err != nil {
		log.Fatal(err)
	}
	// Generate latest.json
	latest := generateK3supVersion(k3sup.Latest)
	err = writeJson(folder, "latest.json", latest)
	if err != nil {
		log.Fatal(err)
	}
}

func getK3sup() (k3sup Tool, err error) {
	return getLatestReleaseFromGithub("alexellis", "k3sup", "k3sup", "Bootstrap Kubernetes with k3s",
		"darwin_amd64",
		"linux_amd64",
		"linux_arm",
		"linux_arm64",
		"windows_amd64")
}

func getK3supVersions() (toolVersions []string, err error) {
	releases, err := getReleasesFromGithub("alexellis", "k3sup", "k3sup")
	if err != nil {
		return nil, err
	}
	for _, release := range releases {
		toolVersions = append(toolVersions, *release.TagName)
	}
	return toolVersions, nil
}

func generateK3supVersion(version string) (toolVersion ToolVersion) {
	baseUrl := "https://github.com/alexellis/k3sup/releases/download/" + version
	return ToolVersion{
		Version: version,
		Platforms: map[string]ToolDownload{
			"darwin_amd64": {
				Url:      baseUrl + "/k3sup-darwin",
				Checksum: baseUrl + "/k3sup-darwin.sha256",
				Alg:      "sha256",
			},
			"linux_amd64": {
				Url:      baseUrl + "/k3sup",
				Checksum: baseUrl + "/k3sup.sha256",
				Alg:      "sha256",
			},
			"linux_arm": {
				Url:      baseUrl + "/k3sup-armhf",
				Checksum: baseUrl + "/k3sup-armhf.sha256",
				Alg:      "sha256",
			},
			"linux_arm64": {
				Url:      baseUrl + "/k3sup-arm64",
				Checksum: baseUrl + "/k3sup-arm64.sha256",
				Alg:      "sha256",
			},
			"windows_amd64": {
				Url:      baseUrl + "/k3sup.exe",
				Checksum: baseUrl + "/k3sup.exe.sha256",
				Alg:      "sha256",
			},
		},
	}
}
