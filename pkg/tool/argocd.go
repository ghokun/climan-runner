package tool

import (
	"log"
)

func init() {
	argocd, err := getArgocd()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, argocd)
}

func getArgocd() (argocd Tool, err error) {
	owner := "argoproj"
	repo := "argo-cd"
	name := "argocd"
	argocd, err = getLatestReleaseFromGithub(owner, repo, name, "Declarative continuous deployment for Kubernetes",
		"darwin_amd64",
		"linux_amd64",
		"windows_amd64")
	argocd.GetVersions = func() (toolVersions []string, err error) {
		releases, err := getReleasesFromGithub(owner, repo, name)
		if err != nil {
			return nil, err
		}
		for _, release := range releases {
			toolVersions = append(toolVersions, *release.TagName)
		}
		return toolVersions, nil
	}
	argocd.GenerateVersion = func(version string) (toolVersion ToolVersion) {
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
	return argocd, err
}
