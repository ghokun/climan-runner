package tool

import (
	"log"
	"os"
	"path/filepath"
)

func init() {
	kn, err := getKn()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, kn)
	// Generate kn specific directory
	folder := filepath.Join(".", "docs", kn.Name)
	os.Mkdir(folder, os.ModePerm)
	// Generate versions.json
	toolVersions, err := getKnVersions()
	if err != nil {
		log.Fatal(err)
	}
	err = writeJson(folder, "versions.json", toolVersions)
	if err != nil {
		log.Fatal(err)
	}
	// Generate template.json
	template := generateKnVersion("{{.Version}}")
	err = writeJson(folder, "template.json", template)
	if err != nil {
		log.Fatal(err)
	}
	// Generate latest.json
	latest := generateKnVersion(kn.Latest)
	err = writeJson(folder, "latest.json", latest)
	if err != nil {
		log.Fatal(err)
	}
}

func getKn() (kn Tool, err error) {
	return getLatestTagFromGithub("knative", "client", "kn", "Knative cli",
		"darwin_amd64",
		"linux_amd64",
		"linux_arm64",
		"linux_ppc64le",
		"linux_s390x",
		"windows_amd64")
}

func getKnVersions() (toolVersions []string, err error) {
	tags, err := getTagsFromGithub("knative", "client", "kn")
	if err != nil {
		return nil, err
	}
	for _, tag := range tags {
		toolVersions = append(toolVersions, *tag.Name)
	}
	return toolVersions, nil
}

func generateKnVersion(version string) (toolVersion ToolVersion) {
	baseUrl := "https://github.com/knative/client/releases/download/" + version
	return ToolVersion{
		Version: version,
		Platforms: map[string]ToolDownload{
			"darwin_amd64": {
				Url:      baseUrl + "/kn-darwin-amd64",
				Checksum: baseUrl + "/checksums.txt",
				Alg:      "sha256",
			},
			"linux_amd64": {
				Url:      baseUrl + "/kn-linux-amd64",
				Checksum: baseUrl + "/checksums.txt",
				Alg:      "sha256",
			},
			"linux_arm64": {
				Url:      baseUrl + "/kn-linux-arm64",
				Checksum: baseUrl + "/checksums.txt",
				Alg:      "sha256",
			},
			"linux_ppc64le": {
				Url:      baseUrl + "/kn-linux-ppc64le",
				Checksum: baseUrl + "/checksums.txt",
				Alg:      "sha256",
			},
			"linux_s390x": {
				Url:      baseUrl + "/kn-linux-s390x",
				Checksum: baseUrl + "/checksums.txt",
				Alg:      "sha256",
			},
			"windows_amd64": {
				Url:      baseUrl + "/kn-windows-amd64.exe",
				Checksum: baseUrl + "/checksums.txt",
				Alg:      "sha256",
			},
		},
	}
}
