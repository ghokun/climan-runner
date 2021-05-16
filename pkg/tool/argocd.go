package tool

import (
	"log"
	"os"
	"path/filepath"
)

func init() {
	argocd, err := getArgocd()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, argocd)
	// Generate argocd specific directory
	folder := filepath.Join(".", "docs", argocd.Name)
	os.Mkdir(folder, os.ModePerm)
	// Generate versions.json
	toolVersions, err := getArgocdVersions()
	if err != nil {
		log.Fatal(err)
	}
	err = writeJson(folder, "versions.json", toolVersions)
	if err != nil {
		log.Fatal(err)
	}
	// Generate template.json
	template := generateArgocdVersion("{{.Version}}")
	err = writeJson(folder, "template.json", template)
	if err != nil {
		log.Fatal(err)
	}
	// Generate latest.json
	latest := generateArgocdVersion(argocd.Latest)
	err = writeJson(folder, "latest.json", latest)
	if err != nil {
		log.Fatal(err)
	}
}

func getArgocd() (argocd Tool, err error) {
	return getLatestReleaseFromGithub("argoproj", "argo-cd", "argocd", "Declarative continuous deployment for Kubernetes",
		"darwin_amd64",
		"linux_amd64",
		"windows_amd64")
}

func getArgocdVersions() (toolVersions []string, err error) {
	releases, err := getReleasesFromGithub("argoproj", "argo-cd", "argocd")
	if err != nil {
		return nil, err
	}
	for _, release := range releases {
		toolVersions = append(toolVersions, *release.TagName)
	}
	return toolVersions, nil
}

func generateArgocdVersion(version string) (toolVersion ToolVersion) {
	baseUrl := "https://github.com/argoproj/argo-cd/releases/download/" + version
	return ToolVersion{
		Version: version,
		Platforms: map[string]ToolDownload{
			"darwin_amd64": {
				Url: baseUrl + "/argocd-darwin-amd64",
			},
			"linux_amd64": {
				Url: baseUrl + "/argocd-linux-amd64",
			},
			"windows_amd64": {
				Url: baseUrl + "/argocd-windows-amd64.exe",
			},
		},
	}
}
