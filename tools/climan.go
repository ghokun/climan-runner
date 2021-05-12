package tools

import (
	"context"
	"errors"
	"runtime"

	"github.com/google/go-github/v34/github"
)

func init() {
	climan := getCliman()
	if supports, err := checkToolSupport(climan.Support); err == nil && supports {
		Tools["climan"] = climan
	}
}

func getCliman() Tool {

	owner := "ghokun"
	repository := "climan"
	name := "climan"
	suffix := ""
	if runtime.GOOS == "windows" {
		suffix = ".exe"
	}
	binaryTemplate := "https://github.com/{{.Owner}}/{{.Repository}}/releases/download/{{.Version}}/{{.Name}}-{{.Os}}-{{.Arch}}{{.Suffix}}"

	return Tool{
		Name:        name,
		Description: "CLI Tools Version Manager",
		Support:     2047,
		GetLatest: func() (latest ToolVersion, err error) {
			release, response, err := client.Repositories.GetLatestRelease(context.Background(), owner, repository)
			if err != nil {
				return
			}
			if err == nil && response.StatusCode == 200 {
				latest.Version = *release.TagName
				latest.URL, err = parseDownloadTemplate(templateVariables{
					Owner:      owner,
					Repository: repository,
					Name:       name,
					Os:         runtime.GOOS,
					Arch:       runtime.GOARCH,
					Suffix:     suffix,
					Version:    *release.TagName,
				}, binaryTemplate)
				return
			}
			// TODO Handle api errors
			err = errors.New("GitHub API error")
			return
		},
		GetAll: func() (all ToolVersions, err error) {
			releases, response, err := client.Repositories.ListReleases(context.Background(), owner, repository, &github.ListOptions{})
			if err != nil {
				return
			}
			if response.StatusCode == 200 {
				for _, release := range releases {
					url, _ := parseDownloadTemplate(templateVariables{
						Owner:      owner,
						Repository: repository,
						Name:       name,
						Os:         runtime.GOOS,
						Arch:       runtime.GOARCH,
						Suffix:     suffix,
						Version:    *release.TagName,
					}, binaryTemplate)
					all = append(all, ToolVersion{
						Version: *release.TagName,
						URL:     url,
					})
				}
				return
			}
			return []ToolVersion{}, nil
		},
		GetApiVersion: func() (version string, err error) {
			return "", nil
		},
		VerifyChecksum: func(version string) (result bool, err error) {
			return false, nil
		},
	}
}
