package tool

import (
	"log"
)

func init() {
	kind, err := getKind()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, kind)
}

func getKind() (kind Tool, err error) {
	owner := "kubernetes-sigs"
	repo := "kind"
	name := "kind"
	kind, err = getLatestReleaseFromGithub(owner, repo, name, "Kubernetes in Docker",
		"darwin_amd64",
		"linux_amd64",
		"linux_arm64",
		"linux_ppc64le",
		"windows_amd64")
	kind.GetVersions = func() (toolVersions []string, err error) {
		releases, err := getReleasesFromGithub(owner, repo, name)
		if err != nil {
			return nil, err
		}
		for _, release := range releases {
			toolVersions = append(toolVersions, *release.TagName)
		}
		return toolVersions, nil
	}
	kind.GenerateVersion = func(version string) (toolVersion ToolVersion) {
		baseUrl := "https://github.com/kubernetes-sigs/kind/releases/download/" + version
		return ToolVersion{
			Version: version,
			Platforms: map[string]ToolDownload{
				"darwin_amd64": {
					Url: baseUrl + "/kind-darwin-amd64",
				},
				"linux_amd64": {
					Url: baseUrl + "/kind-linux-amd64",
				},
				"linux_arm64": {
					Url: baseUrl + "/kind-linux-arm64",
				},
				"linux_ppc64le": {
					Url: baseUrl + "/kind-linux-ppc64le",
				},
				"windows_amd64": {
					Url: baseUrl + "/kind-windows-amd64",
				},
			},
		}
	}
	return kind, err
}
