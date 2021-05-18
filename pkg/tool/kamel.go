package tool

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	kamel, err := getKamel()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, kamel)
	// Generate kamel specific directory
	folder := filepath.Join(".", "docs", kamel.Name)
	os.Mkdir(folder, os.ModePerm)
	// Generate versions.json
	toolVersions, err := getKamelVersions()
	if err != nil {
		log.Fatal(err)
	}
	err = writeJson(folder, "versions.json", toolVersions)
	if err != nil {
		log.Fatal(err)
	}
	// Generate template.json
	template := generateKamelVersion("{{.Version}}")
	err = writeJson(folder, "template.json", template)
	if err != nil {
		log.Fatal(err)
	}
	// Generate latest.json
	latest := generateKamelVersion(kamel.Latest)
	err = writeJson(folder, "latest.json", latest)
	if err != nil {
		log.Fatal(err)
	}
}

func getKamel() (kamel Tool, err error) {
	return getLatestReleaseFromGithub("apache", "camel-k", "kamel", "Cli for Apacke Camel-K",
		"darwin_amd64",
		"linux_amd64",
		"windows_amd64")
}

func getKamelVersions() (toolVersions []string, err error) {
	releases, err := getReleasesFromGithub("apache", "camel-k", "kamel")
	if err != nil {
		return nil, err
	}
	for _, release := range releases {
		toolVersions = append(toolVersions, *release.TagName)
	}
	return toolVersions, nil
}

func generateKamelVersion(version string) (toolVersion ToolVersion) {
	baseUrl := "https://github.com/apache/camel-k/releases/download/" + version
	withOutV := strings.TrimPrefix(version, "v")

	return ToolVersion{
		Version: version,
		Platforms: map[string]ToolDownload{
			"darwin_amd64": {
				Url: baseUrl + "/camel-k-client-" + withOutV + "-mac-64bit.tar.gz",
			},
			"linux_amd64": {
				Url: baseUrl + "/camel-k-client-" + withOutV + "-linux-64bit.tar.gz",
			},
			"windows_amd64": {
				Url: baseUrl + "/camel-k-client-" + withOutV + "-windows-64bit.tar.gz",
			},
		},
	}
}
