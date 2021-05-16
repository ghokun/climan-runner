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
	folder := "./docs/" + kubectl.Name
	os.Mkdir(folder, os.ModePerm)
	// Generate versions.json
	toolVersions, err := getKubectlVersions()
	if err != nil {
		log.Fatal(err)
	}
	allVersions, err := json.Marshal(toolVersions)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(folder+"/versions.json", allVersions, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// Generate template.json
	template := generateKubectlVersion("{{.Version}}")
	templateData, err := json.Marshal(template)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(folder+"/template.json", templateData, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// Generate latest.json
	latest := generateKubectlVersion(kubectl.Latest)
	latestData, err := json.Marshal(latest)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(folder+"/latest.json", latestData, 0644)
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

func getKubectlVersions() (toolVersions []string, err error) {
	releases, err := getReleasesFromGithub("kubernetes", "kubernetes", "kubectl")
	if err != nil {
		return nil, err
	}
	for _, release := range releases {
		toolVersions = append(toolVersions, *release.TagName)
	}
	return toolVersions, nil
}

func generateKubectlVersion(version string) (toolVersion ToolVersion) {
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
	}
}
