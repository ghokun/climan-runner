package tools

import (
	"context"
	"errors"
	"io"
	"net/http"
	"runtime"

	"github.com/google/go-github/v34/github"
)

func init() {
	Tools["kubectl"] = getKubectl()
}

func getKubectl() Tool {

	owner := "kubernetes"
	repository := "kubernetes"
	name := "kubectl"

	binaryTemplate := "https://storage.googleapis.com/kubernetes-release/release/{{.Version}}/kubernetes-client-{{.Os}}-{{.Arch}}{{.Suffix}}"
	//sha256Template := "https://dl.k8s.io/{{.Version}}/bin/{{.Os}}/{{.Arch}}/kubectl.sha256"

	return Tool{
		Name:        name,
		Description: "CLI Tools Version Manager",
		Support:     2047,
		GetLatest: func() (latest ToolVersion, err error) {
			response, err := http.Get("https://storage.googleapis.com/kubernetes-release/release/stable.txt")
			if err != nil {
				return
			}
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			version := string(body)

			url, _ := parseDownloadTemplate(templateVariables{
				Owner:      owner,
				Repository: repository,
				Name:       name,
				Os:         runtime.GOOS,
				Arch:       runtime.GOARCH,
				Suffix:     ".tar.gz",
				Version:    version,
			}, binaryTemplate)

			return ToolVersion{
				Version: version,
				URL:     url,
			}, err
		},
		GetAll: func() (all ToolVersions, err error) {
			releases, response, err := client.Repositories.ListReleases(context.Background(), owner, repository, &github.ListOptions{
				Page:    0,
				PerPage: 100,
			})
			if err != nil {
				return
			}
			if err == nil && response.StatusCode == 200 {
				for _, release := range releases {
					url, _ := parseDownloadTemplate(templateVariables{
						Owner:      owner,
						Repository: repository,
						Name:       name,
						Os:         runtime.GOOS,
						Arch:       runtime.GOARCH,
						Suffix:     ".tar.gz",
						Version:    *release.TagName,
					}, binaryTemplate)
					all = append(all, ToolVersion{
						Version: *release.TagName,
						URL:     url,
					})
				}
				return
			}
			err = errors.New("GitHub API error")
			return
		},
		GetApiVersion: func() (version string, err error) {
			return "", nil
		},
		VerifyChecksum: func(version string) (result bool, err error) {
			return false, nil
		},
	}
}
