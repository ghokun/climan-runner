package tool

import (
	"log"
)

func init() {
	minikube, err := getMinikube()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, minikube)
	generateToolSpecificFiles("minikube", minikube.Latest, getMinikubeVersions, generateMinikubeVersion)
}

func getMinikube() (minikube Tool, err error) {
	return getLatestReleaseFromGithub("kubernetes", "minikube", "minikube", "Run Kubernetes locally",
		"darwin_amd64",
		"darwin_arm64",
		"linux_amd64",
		"linux_arm",
		"linux_arm64",
		"linux_ppc64le",
		"linux_s390x",
		"windows_amd64")
}

func getMinikubeVersions() (toolVersions []string, err error) {
	releases, err := getReleasesFromGithub("kubernetes", "minikube", "minikube")
	if err != nil {
		return nil, err
	}
	for _, release := range releases {
		toolVersions = append(toolVersions, *release.TagName)
	}
	return toolVersions, nil
}

func generateMinikubeVersion(version string) (toolVersion ToolVersion) {
	baseUrl := "https://github.com/kubernetes/minikube/releases/download/" + version
	return ToolVersion{
		Version: version,
		Platforms: map[string]ToolDownload{
			"darwin_amd64": {
				Url:      baseUrl + "/minikube-darwin-amd64",
				Checksum: baseUrl + "/minikube-darwin-amd64.sha256",
				Alg:      "sha256",
			},
			"darwin_arm64": {
				Url:      baseUrl + "/minikube-darwin-arm64",
				Checksum: baseUrl + "/minikube-darwin-arm64.sha256",
				Alg:      "sha256",
			},
			"linux_amd64": {
				Url:      baseUrl + "/minikube-linux-amd64",
				Checksum: baseUrl + "/minikube-linux-amd64.sha256",
				Alg:      "sha256",
			},
			"linux_arm": {
				Url:      baseUrl + "/minikube-linux-arm",
				Checksum: baseUrl + "/minikube-linux-arm.sha256",
				Alg:      "sha256",
			},
			"linux_arm64": {
				Url:      baseUrl + "/minikube-linux-arm64",
				Checksum: baseUrl + "/minikube-linux-arm64.sha256",
				Alg:      "sha256",
			},
			"linux_ppc64le": {
				Url:      baseUrl + "/minikube-linux-ppc64le",
				Checksum: baseUrl + "/minikube-linux-ppc64le.sha256",
				Alg:      "sha256",
			},
			"linux_s390x": {
				Url:      baseUrl + "/minikube-linux-s390x",
				Checksum: baseUrl + "/minikube-linux-s390x.sha256",
				Alg:      "sha256",
			},
			"windows_amd64": {
				Url:      baseUrl + "/minikube.exe",
				Checksum: baseUrl + "/minikube.exe.sha256",
				Alg:      "sha256",
			},
		},
	}
}
