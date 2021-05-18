package tool

import (
	"log"
	"os"
	"path/filepath"
)

func init() {
	kind, err := getKind()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, kind)
	// Generate kind specific directory
	folder := filepath.Join(".", "docs", kind.Name)
	os.Mkdir(folder, os.ModePerm)
	// Generate versions.json
	toolVersions, err := getKindVersions()
	if err != nil {
		log.Fatal(err)
	}
	err = writeJson(folder, "versions.json", toolVersions)
	if err != nil {
		log.Fatal(err)
	}
	// Generate template.json
	template := generateKindVersion("{{.Version}}")
	err = writeJson(folder, "template.json", template)
	if err != nil {
		log.Fatal(err)
	}
	// Generate latest.json
	latest := generateKindVersion(kind.Latest)
	err = writeJson(folder, "latest.json", latest)
	if err != nil {
		log.Fatal(err)
	}
}

func getKind() (kind Tool, err error) {
	return getLatestReleaseFromGithub("kubernetes-sigs", "kind", "kind", "Kubernetes in Docker",
		"darwin_amd64",
		"linux_amd64",
		"linux_arm64",
		"linux_ppc64le",
		"windows_amd64")
}

func getKindVersions() (toolVersions []string, err error) {
	releases, err := getReleasesFromGithub("kubernetes-sigs", "kind", "kind")
	if err != nil {
		return nil, err
	}
	for _, release := range releases {
		toolVersions = append(toolVersions, *release.TagName)
	}
	return toolVersions, nil
}

func generateKindVersion(version string) (toolVersion ToolVersion) {
	baseUrl := "https://github.com/kubernetes-sigs/kind/releases/download/" + version
	return ToolVersion{
		Version: version,
		Platforms: map[string]ToolDownload{
			"darwin_amd64": {
				Url:      baseUrl + "/kind-darwin-amd64",
			},
			"linux_amd64": {
				Url:      baseUrl + "/kind-linux-amd64",
			},
			"linux_arm64": {
				Url:      baseUrl + "/kind-linux-arm64",
			},
			"linux_ppc64le": {
				Url:      baseUrl + "/kind-linux-ppc64le",
			},
			"windows_amd64": {
				Url:      baseUrl + "/kind-windows-amd64",
			},
		},
	}
}
