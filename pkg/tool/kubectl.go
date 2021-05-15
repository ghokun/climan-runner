package tool

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/ghokun/climan-runner/pkg/platform"
)

func init() {
	kubectl, err := getKubectl()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, kubectl)
	toolVersions, err := getKubectlVersions()
	if err != nil {
		log.Fatal(err)
	}
	var versions []string
	for _, toolVersion := range toolVersions {
		versions = append(versions, toolVersion.Version)
		data, err := json.Marshal(toolVersion)
		if err != nil {
			log.Fatal(err)
		}
		os.Mkdir("./docs/"+kubectl.Name, os.ModePerm)
		location := "./docs/" + kubectl.Name + "/" + toolVersion.Version + ".json"
		err = ioutil.WriteFile(location, data, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
	allVersions, err := json.Marshal(versions)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("./docs/"+kubectl.Name+"/versions.json", allVersions, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func getKubectl() (kubectl Tool, err error) {
	response, err := http.Get("https://storage.googleapis.com/kubernetes-release/release/stable.txt")
	if err != nil {
		return kubectl, err
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			return kubectl, err
		}
		return Tool{
			Name:        "kubectl",
			Description: "Kubernetes command line tool",
			Supports: platform.CalculateSupportedPlatforms(
				[]string{"darwin_amd64",
					"darwin_arm64",
					"linux_386",
					"linux_amd64",
					"linux_arm",
					"linux_arm64",
					"linux_ppc64le",
					"linux_s390x",
					"windows_386",
					"windows_amd64"}),
			Latest: string(bodyBytes),
		}, nil
	}
	return kubectl, errors.New("error while fetcing latest version of kubectl")
}

func getKubectlVersions() (toolVersions []ToolVersion, err error) {
	releases, err := getReleasesFromGithub("kubernetes", "kubernetes", "kubectl")
	if err != nil {
		return nil, err
	}
	for _, release := range releases {
		toolVersion, err := generateKubectlVersion(*release.TagName)
		if err != nil {
			return nil, err
		}
		toolVersions = append(toolVersions, toolVersion)
	}
	return toolVersions, nil
}

func generateKubectlVersion(version string) (toolVersion ToolVersion, err error) {
	baseUrl := "https://storage.googleapis.com/kubernetes-release/release/" + version
	return ToolVersion{
		Version: version,
		Platforms: map[string]ToolDownload{
			"darwin_amd64": {
				Url: baseUrl + "/bin/darwin/amd64/kubectl",
			},
			"darwin_arm64": {
				Url: baseUrl + "/bin/darwin/arm64/kubectl",
			},
			"linux_386": {
				Url: baseUrl + "/bin/linux/386/kubectl",
			},
			"linux_amd64": {
				Url: baseUrl + "/bin/linux/amd64/kubectl",
			},
			"linux_arm": {
				Url: baseUrl + "/bin/linux/arm/kubectl",
			},
			"linux_arm64": {
				Url: baseUrl + "/bin/linux/arm64/kubectl",
			},
			"linux_ppc64le": {
				Url: baseUrl + "/bin/linux/ppc64le/kubectl",
			},
			"linux_s390x": {
				Url: baseUrl + "/bin/linux/s390x/kubectl",
			},
			"windows_386": {
				Url: baseUrl + "/bin/windows/386/kubectl.exe",
			},
			"windows_amd64": {
				Url: baseUrl + "/bin/windows/amd64/kubectl.exe",
			},
		},
	}, nil
}
