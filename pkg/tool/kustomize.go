package tool

import (
	"fmt"
	"log"
	"strings"

	"github.com/ghokun/climan-runner/pkg/platform"
	"golang.org/x/mod/semver"
)

func init() {
	kustomize, err := getKustomize()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, kustomize)
}

func getKustomize() (kustomize Tool, err error) {

	owner := "kubernetes-sigs"
	repo := "kustomize"
	name := "kustomize"

	// kustomize repository has tags with prefixes
	// Github api has 100 element limit for each tag page
	// A better solution is needed to fix
	// TODO for later
	tags, err := getTagsFromGithubWithPage(owner, repo, name, 1, 100)
	if err != nil {
		return kustomize, err
	}
	secondPage, err := getTagsFromGithubWithPage(owner, repo, name, 2, 100)
	if err != nil {
		return kustomize, err
	}
	tags = append(tags, secondPage...)
	if len(tags) < 0 {
		return kustomize, fmt.Errorf("no tag found for %q", name)
	}
	var versions []string
	var latest string
	for _, tag := range tags {
		if strings.HasPrefix(tag.GetName(), name) {
			version := strings.TrimPrefix(tag.GetName(), "kustomize/")
			versions = append(versions, version)
			if semver.Compare(latest, version) == -1 {
				latest = version
			}
		}
	}
	return Tool{
		Name:        name,
		Description: "Customization of kubernetes YAML configurations",
		Supports: platform.CalculateSupportedPlatforms(
			[]string{"darwin_amd64",
				"linux_amd64",
				"linux_arm64",
				"windows_amd64"}),
		Latest: latest,
		GetVersions: func() ([]string, error) {
			return versions, nil
		},
		GenerateVersion: func(version string) (toolVersion ToolVersion) {
			baseUrl := "https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize%2F" + version
			checksum := baseUrl + "/checksums.txt"
			alg := "sha256"
			return ToolVersion{
				Version: version,
				Platforms: map[string]ToolDownload{
					"darwin_amd64": {
						Url:      baseUrl + "/kustomize_" + version + "_darwin_amd64.tar.gz",
						Checksum: checksum,
						Alg:      alg,
					},
					"linux_amd64": {
						Url:      baseUrl + "/kustomize_" + version + "_linux_amd64.tar.gz",
						Checksum: checksum,
						Alg:      alg,
					},
					"linux_arm64": {
						Url:      baseUrl + "/kustomize_" + version + "_linux_arm64.tar.gz",
						Checksum: checksum,
						Alg:      alg,
					},
					"windows_amd64": {
						Url:      baseUrl + "/kustomize_" + version + "_windows_amd64.tar.gz",
						Checksum: checksum,
						Alg:      alg,
					},
				},
			}
		},
	}, nil
}
