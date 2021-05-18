package tool

import (
	"log"
	"os"
	"path/filepath"
)

func init() {
	linkerd2, err := getLinkerd2()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, linkerd2)
	// Generate linkerd2 specific directory
	folder := filepath.Join(".", "docs", linkerd2.Name)
	os.Mkdir(folder, os.ModePerm)
	// Generate versions.json
	toolVersions, err := getLinkerd2Versions()
	if err != nil {
		log.Fatal(err)
	}
	err = writeJson(folder, "versions.json", toolVersions)
	if err != nil {
		log.Fatal(err)
	}
	// Generate template.json
	template := generateLinkerd2Version("{{.Version}}")
	err = writeJson(folder, "template.json", template)
	if err != nil {
		log.Fatal(err)
	}
	// Generate latest.json
	latest := generateLinkerd2Version(linkerd2.Latest)
	err = writeJson(folder, "latest.json", latest)
	if err != nil {
		log.Fatal(err)
	}
}

func getLinkerd2() (linkerd2 Tool, err error) {
	return getLatestReleaseFromGithub("linkerd", "linkerd2", "linkerd2", "Ultralight, security-first service mesh for Kubernetes",
		"darwin_amd64",
		"darwin_arm64",
		"linux_amd64",
		"linux_arm",
		"linux_arm64",
		"windows_amd64")
}

func getLinkerd2Versions() (toolVersions []string, err error) {
	releases, err := getReleasesFromGithub("linkerd", "linkerd2", "linkerd2")
	if err != nil {
		return nil, err
	}
	for _, release := range releases {
		toolVersions = append(toolVersions, *release.TagName)
	}
	return toolVersions, nil
}

func generateLinkerd2Version(version string) (toolVersion ToolVersion) {
	baseUrl := "https://github.com/linkerd/linkerd2/releases/download/" + version
	return ToolVersion{
		Version: version,
		Platforms: map[string]ToolDownload{
			"darwin_amd64": {
				Url:      baseUrl + "/linkerd2-cli-" + version + "-darwin",
				Checksum: baseUrl + "/linkerd2-cli-" + version + "-darwin.sha256",
				Alg:      "sha256",
			},
			"darwin_arm64": {
				Url:      baseUrl + "/linkerd2-cli-" + version + "-darwin-arm64",
				Checksum: baseUrl + "/linkerd2-cli-" + version + "-darwin-arm64.sha256",
				Alg:      "sha256",
			},
			"linux_amd64": {
				Url:      baseUrl + "/linkerd2-cli-" + version + "-linux-amd64",
				Checksum: baseUrl + "/linkerd2-cli-" + version + "-linux-amd64.sha256",
				Alg:      "sha256",
			},
			"linux_arm": {
				Url:      baseUrl + "/linkerd2-cli-" + version + "-linux-arm",
				Checksum: baseUrl + "/linkerd2-cli-" + version + "-linux-arm.sha256",
				Alg:      "sha256",
			},
			"linux_arm64": {
				Url:      baseUrl + "/linkerd2-cli-" + version + "-linux-arm64",
				Checksum: baseUrl + "/linkerd2-cli-" + version + "-linux-arm64.sha256",
				Alg:      "sha256",
			},
			"windows_amd64": {
				Url:      baseUrl + "/linkerd2-cli-" + version + "-windows.exe",
				Checksum: baseUrl + "/linkerd2-cli-" + version + "-windows.exe.sha256",
				Alg:      "sha256",
			},
		},
	}
}