package tool

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/ghokun/climan-runner/pkg/platform"
	"github.com/google/go-github/v35/github"
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

	tags, response, err := client.Repositories.ListTags(context.Background(), owner, repo, &github.ListOptions{
		Page:    0,
		PerPage: 1000,
	})
	if err != nil {
		return kustomize, err
	}
	if len(tags) < 0 {
		return kustomize, errors.Unwrap(fmt.Errorf("no tag found for %q", name))
	}
	if response.StatusCode == 200 {
		for _, tag := range tags {
			// This is the way
			if strings.HasPrefix(*tag.Name, name) {
				return Tool{
					Name:        name,
					Description: "Customization of kubernetes YAML configurations",
					Supports: platform.CalculateSupportedPlatforms(
						[]string{"darwin_amd64",
							"linux_amd64",
							"linux_arm64",
							"windows_amd64"}),
					Latest: strings.TrimPrefix(*tag.Name, "kustomize/"),
					GetVersions: func() (toolVersions []string, err error) {
						releases, err := getReleasesFromGithub(owner, repo, name)
						if err != nil {
							return nil, err
						}
						for _, release := range releases {
							if strings.HasPrefix(release.GetTagName(), name) {
								toolVersions = append(toolVersions, strings.TrimPrefix(release.GetTagName(), "kustomize/"))
							}
						}
						return toolVersions, nil
					},
					GenerateVersion: func(version string) (toolVersion ToolVersion) {
						baseUrl := "https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize%2F" + version
						alg := "sha256"
						return ToolVersion{
							Version: version,
							Platforms: map[string]ToolDownload{
								"darwin_amd64": {
									Url:      baseUrl + "/kustomize_" + version + "_darwin_amd64.tar.gz",
									Checksum: baseUrl + "/checksums.txt",
									Alg:      alg,
								},
								"linux_amd64": {
									Url:      baseUrl + "/kustomize_" + version + "_linux_amd64.tar.gz",
									Checksum: baseUrl + "/checksums.txt",
									Alg:      alg,
								},
								"linux_arm64": {
									Url:      baseUrl + "/kustomize_" + version + "_linux_arm64.tar.gz",
									Checksum: baseUrl + "/checksums.txt",
									Alg:      alg,
								},
								"windows_amd64": {
									Url:      baseUrl + "/kustomize_" + version + "_windows_amd64.tar.gz",
									Checksum: baseUrl + "/checksums.txt",
									Alg:      alg,
								},
							},
						}
					},
				}, nil
			}
		}
	}
	return kustomize, errors.Unwrap(fmt.Errorf("error while fetcing latest version of %q", name))
}

// TODO fix getting versions
