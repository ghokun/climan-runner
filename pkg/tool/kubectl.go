package tool

import (
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/ghokun/climan-runner/pkg/platform"
)

func init() {
	kubectl, err := getKubectl()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, kubectl)
	generateToolSpecificFiles("kubectl", kubectl.Latest, getKubectlVersions, generateKubectlVersion)
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
				Url:      baseUrl + "/bin/darwin/amd64/kubectl",
				Checksum: baseUrl + "/bin/darwin/amd64/kubectl.sha256",
				Alg:      "sha256",
			},
			"darwin_arm64": {
				Url:      baseUrl + "/bin/darwin/arm64/kubectl",
				Checksum: baseUrl + "/bin/darwin/arm64/kubectl.sha256",
				Alg:      "sha256",
			},
			"linux_386": {
				Url:      baseUrl + "/bin/linux/386/kubectl",
				Checksum: baseUrl + "/bin/linux/386/kubectl.sha256",
				Alg:      "sha256",
			},
			"linux_amd64": {
				Url:      baseUrl + "/bin/linux/amd64/kubectl",
				Checksum: baseUrl + "/bin/linux/amd64/kubectl.sha256",
				Alg:      "sha256",
			},
			"linux_arm": {
				Url:      baseUrl + "/bin/linux/arm/kubectl",
				Checksum: baseUrl + "/bin/linux/arm/kubectl.sha256",
				Alg:      "sha256",
			},
			"linux_arm64": {
				Url:      baseUrl + "/bin/linux/arm64/kubectl",
				Checksum: baseUrl + "/bin/linux/arm64/kubectl.sha256",
				Alg:      "sha256",
			},
			"linux_ppc64le": {
				Url:      baseUrl + "/bin/linux/ppc64le/kubectl",
				Checksum: baseUrl + "/bin/linux/ppc64le/kubectl.sha256",
				Alg:      "sha256",
			},
			"linux_s390x": {
				Url:      baseUrl + "/bin/linux/s390x/kubectl",
				Checksum: baseUrl + "/bin/linux/s390x/kubectl.sha256",
				Alg:      "sha256",
			},
			"windows_386": {
				Url:      baseUrl + "/bin/windows/386/kubectl.exe",
				Checksum: baseUrl + "/bin/windows/386/kubectl.exe.sha256",
				Alg:      "sha256",
			},
			"windows_amd64": {
				Url:      baseUrl + "/bin/windows/amd64/kubectl.exe",
				Checksum: baseUrl + "/bin/windows/amd64/kubectl.exe.sha256",
				Alg:      "sha256",
			},
		},
	}
}
