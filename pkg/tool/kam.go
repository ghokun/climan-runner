package tool

import (
	"log"
	"os"
	"path/filepath"
)

func init() {
	kam, err := getKam()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, kam)
	// Generate kam specific directory
	folder := filepath.Join(".", "docs", kam.Name)
	os.Mkdir(folder, os.ModePerm)
	// Generate versions.json
	toolVersions, err := getKamVersions()
	if err != nil {
		log.Fatal(err)
	}
	err = writeJson(folder, "versions.json", toolVersions)
	if err != nil {
		log.Fatal(err)
	}
	// Generate template.json
	template := generateKamVersion("{{.Version}}")
	err = writeJson(folder, "template.json", template)
	if err != nil {
		log.Fatal(err)
	}
	// Generate latest.json
	latest := generateKamVersion(kam.Latest)
	err = writeJson(folder, "latest.json", latest)
	if err != nil {
		log.Fatal(err)
	}
}

func getKam() (kam Tool, err error) {
	return getLatestReleaseFromGithub("redhat-developer", "kam", "kam", "GitOps Application Manager",
		"darwin_amd64",
		"linux_amd64",
		"windows_amd64")
}

func getKamVersions() (toolVersions []string, err error) {
	releases, err := getReleasesFromGithub("redhat-developer", "kam", "kam")
	if err != nil {
		return nil, err
	}
	for _, release := range releases {
		toolVersions = append(toolVersions, *release.TagName)
	}
	return toolVersions, nil
}

func generateKamVersion(version string) (toolVersion ToolVersion) {
	baseUrl := "https://github.com/redhat-developer/kam/releases/download/" + version
	return ToolVersion{
		Version: version,
		Platforms: map[string]ToolDownload{
			"darwin_amd64": {
				Url: baseUrl + "/kam_darwin_amd64",
			},
			"linux_amd64": {
				Url: baseUrl + "/kam_linux_amd64",
			},
			"windows_amd64": {
				Url: baseUrl + "/kam_windows_amd64.exe",
			},
		},
	}
}
