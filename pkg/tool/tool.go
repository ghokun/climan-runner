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

func GenerateEachTool() (err error) {
	return nil
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
