package tool

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/ghokun/climan-runner/pkg/platform"
	"github.com/google/go-github/v35/github"
)

type Tool struct {
	Name            string                   `json:"name,omitempty"`
	Description     string                   `json:"description,omitempty"`
	Supports        int                      `json:"supports,omitempty"`
	Latest          string                   `json:"latest,omitempty"`
	GetVersions     func() ([]string, error) `json:"-"`
	GenerateVersion func(string) ToolVersion `json:"-"`
}

type ToolVersion struct {
	Version   string                  `json:"version,omitempty"`
	Supports  int                     `json:"supports,omitempty"`
	Platforms map[string]ToolDownload `json:"platforms,omitempty"`
}

type ToolVersions struct {
	Supports int      `json:"supports,omitempty"`
	Versions []string `json:"versions,omitempty"`
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
	log.Println("Generating tools.json")
	err = ioutil.WriteFile(toolsFile, data, 0644)
	for _, tool := range Tools {
		log.Printf("Generating files for %q", tool.Name)
		generateToolSpecificFiles(tool.Name, tool.Latest, tool.Supports, tool.GetVersions, tool.GenerateVersion)
	}
	return err
}

func generateToolSpecificFiles(toolName string, latest string, supports int, getVersions func() ([]string, error), generateVersion func(string) ToolVersion) {
	folder := filepath.Join(".", "docs", toolName)
	os.Mkdir(folder, os.ModePerm)

	// Generate versions.json
	versionsJson, err := getVersions()
	if err != nil {
		log.Fatal(err)
	}
	err = writeJson(folder, "versions.json", ToolVersions{Versions: versionsJson, Supports: supports})
	if err != nil {
		log.Fatal(err)
	}

	// Generate template.json
	templateJson := generateVersion("{{.Version}}")
	templateJson.Supports = supports
	err = writeJson(folder, "template.json", templateJson)
	if err != nil {
		log.Fatal(err)
	}

	// Generate latest.json
	latestJson := generateVersion(latest)
	latestJson.Supports = supports
	err = writeJson(folder, "latest.json", latestJson)
	if err != nil {
		log.Fatal(err)
	}
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
	return releases, errors.Unwrap(fmt.Errorf("error while fetcing releases of %q", name))
}

func getTagsFromGithub(owner string, repo string, name string) (tags []*github.RepositoryTag, err error) {
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
	return tags, errors.Unwrap(fmt.Errorf("error while fetcing tags of %q", name))
}

func getVersionsFromMirrorOpenshift(name string, url string) (toolVersions []string, err error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, errors.Unwrap(fmt.Errorf("error while fetcing releases of %q", name))
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, errors.Unwrap(fmt.Errorf("error while scraping %q", url))
	}
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		if len(text) > 0 && text[0] >= '0' && text[0] <= '9' {
			toolVersions = append(toolVersions, strings.ReplaceAll(text, "/", ""))
		}
	})
	return toolVersions, err
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
