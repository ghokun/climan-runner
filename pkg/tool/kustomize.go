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
	desc := "Customization of kubernetes YAML configurations"

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
			if strings.HasPrefix(*tag.Name, "kustomize") {
				return Tool{
					Name:        name,
					Description: desc,
					Supports: platform.CalculateSupportedPlatforms([]string{"darwin_amd64",
						"linux_amd64",
						"linux_arm64",
						"windows_amd64"}),
					Latest: strings.Split(*tag.Name, "/")[1],
				}, nil
			}
		}
	}
	return kustomize, errors.Unwrap(fmt.Errorf("error while fetcing latest version of %q", name))
}
