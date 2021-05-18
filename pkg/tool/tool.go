package tool

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

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
	toolsFile := filepath.Join(".", "docs", "tools.json")
	err = ioutil.WriteFile(toolsFile, data, 0644)
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
		PerPage: 100,
	})
	if err != nil {
		return releases, err
	}
	if response.StatusCode == 200 {
		return releases, nil
	}
	return releases, errors.Unwrap(fmt.Errorf("error while fetcing latest version of %q", name))
}

func getTagsFromGithub(owner string, repo string, name string) (releases []*github.RepositoryTag, err error) {
	tags, response, err := client.Repositories.ListTags(context.Background(), owner, repo, &github.ListOptions{
		Page:    0,
		PerPage: 100,
	})
	if err != nil {
		return tags, err
	}
	if response.StatusCode == 200 {
		return tags, nil
	}
	return tags, errors.Unwrap(fmt.Errorf("error while fetcing latest version of %q", name))
}

func writeJson(folder string, filename string, data interface{}) (err error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.Join(folder, filename), bytes, 0644)
	if err != nil {
		return err
	}
	return nil
}
