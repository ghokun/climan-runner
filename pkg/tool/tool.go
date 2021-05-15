package tool

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/ghokun/climan-runner/pkg/platform"
	"github.com/google/go-github/v35/github"
)

type Tool struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Supports    int    `json:"supports,omitempty"`
	Latest      string `json:"latest,omitempty"`
}

type ToolVersion struct {
	Version   string                  `json:"version,omitempty"`
	Platforms map[string]ToolDownload `json:"platforms,omitempty"`
}

type ToolDownload struct {
	Url      string `json:"url,omitempty"`
	Checksum string `json:"checksum,omitempty"`
	Alg      string `json:"alg,omitempty"`
}

var (
	Tools  = []Tool{}
	client = github.NewClient(nil)
)

func GenerateTools() (err error) {
	data, err := json.Marshal(Tools)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("./docs/tools.json", data, 0644)
	return err
}

func getLatestReleaseFromGithub(owner string, repo string, name string, desc string, platforms ...string) (t Tool, err error) {
	release, response, err := client.Repositories.GetLatestRelease(context.Background(), owner, repo)
	if err != nil {
		return t, err
	}
	if response.StatusCode == 200 {
		return Tool{
			Name:        name,
			Description: desc,
			Supports:    platform.CalculateSupportedPlatforms(platforms),
			Latest:      *release.TagName,
		}, nil
	}
	return t, errors.Unwrap(fmt.Errorf("error while fetcing latest version of %q", name))
}

func getLatestTagFromGithub(owner string, repo string, name string, desc string, platforms ...string) (t Tool, err error) {
	tags, response, err := client.Repositories.ListTags(context.Background(), owner, repo, &github.ListOptions{
		Page:    0,
		PerPage: 1,
	})
	if err != nil {
		return t, err
	}
	if len(tags) < 0 {
		return t, errors.Unwrap(fmt.Errorf("no tag found for %q", name))
	}
	if response.StatusCode == 200 {
		return Tool{
			Name:        name,
			Description: desc,
			Supports:    platform.CalculateSupportedPlatforms(platforms),
			Latest:      *tags[0].Name,
		}, nil
	}
	return t, errors.Unwrap(fmt.Errorf("error while fetcing latest version of %q", name))
}

func getReleasesFromGithub(owner string, repo string, name string) (releases []*github.RepositoryRelease, err error) {
	releases, response, err := client.Repositories.ListReleases(context.Background(), owner, repo, &github.ListOptions{
		Page:    0,
		PerPage: 1000,
	})
	if err != nil {
		return releases, err
	}
	if response.StatusCode == 200 {
		return releases, nil
	}
	return releases, errors.Unwrap(fmt.Errorf("error while fetcing latest version of %q", name))
}
