package tool

import (
	"log"
	"os"
	"path/filepath"
)

func init() {
	plumber, err := getPlumber()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, plumber)
	// Generate plumber specific directory
	folder := filepath.Join(".", "docs", plumber.Name)
	os.Mkdir(folder, os.ModePerm)
	// Generate versions.json
	toolVersions, err := getPlumberVersions()
	if err != nil {
		log.Fatal(err)
	}
	err = writeJson(folder, "versions.json", toolVersions)
	if err != nil {
		log.Fatal(err)
	}
	// Generate template.json
	template := generatePlumberVersion("{{.Version}}")
	err = writeJson(folder, "template.json", template)
	if err != nil {
		log.Fatal(err)
	}
	// Generate latest.json
	latest := generatePlumberVersion(plumber.Latest)
	err = writeJson(folder, "latest.json", latest)
	if err != nil {
		log.Fatal(err)
	}
}

func getPlumber() (plumber Tool, err error) {
	return getLatestReleaseFromGithub("batchcorp", "plumber", "plumber", "Cli for messaging systems",
		"darwin_amd64",
		"linux_amd64",
		"windows_amd64")
}

func getPlumberVersions() (toolVersions []string, err error) {
	releases, err := getReleasesFromGithub("batchcorp", "plumber", "plumber")
	if err != nil {
		return nil, err
	}
	for _, release := range releases {
		toolVersions = append(toolVersions, *release.TagName)
	}
	return toolVersions, nil
}

func generatePlumberVersion(version string) (toolVersion ToolVersion) {
	baseUrl := "https://github.com/batchcorp/plumber/releases/download/" + version
	return ToolVersion{
		Version: version,
		Platforms: map[string]ToolDownload{
			"darwin_amd64": {
				Url: baseUrl + "/plumber-darwin",
			},
			"linux_amd64": {
				Url: baseUrl + "/plumber-linux",
			},
			"windows_amd64": {
				Url: baseUrl + "/plumber-windows.exe",
			},
		},
	}
}
