package tool

import (
	"log"
	"os"
	"path/filepath"
)

func init() {
	climan, err := getCliman()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, climan)
	// Generate climan specific directory
	folder := filepath.Join(".", "docs", climan.Name)
	os.Mkdir(folder, os.ModePerm)
	// Generate versions.json
	toolVersions, err := getClimanVersions()
	if err != nil {
		log.Fatal(err)
	}
	err = writeJson(folder, "versions.json", toolVersions)
	if err != nil {
		log.Fatal(err)
	}
	// Generate template.json
	template := generateClimanVersion("{{.Version}}")
	err = writeJson(folder, "template.json", template)
	if err != nil {
		log.Fatal(err)
	}
	// Generate latest.json
	latest := generateClimanVersion(climan.Latest)
	err = writeJson(folder, "latest.json", latest)
	if err != nil {
		log.Fatal(err)
	}
}

func getCliman() (climan Tool, err error) {
	return getLatestReleaseFromGithub("ghokun", "climan", "climan", "Cloud tools cli manager",
		"darwin_amd64",
		"darwin_arm64",
		"linux_386",
		"linux_amd64",
		"linux_arm",
		"linux_arm64",
		"linux_ppc64le",
		"linux_s390x",
		"windows_386",
		"windows_amd64",
		"windows_arm")
}

func getClimanVersions() (toolVersions []string, err error) {
	releases, err := getReleasesFromGithub("ghokun", "climan", "climan")
	if err != nil {
		return nil, err
	}
	for _, release := range releases {
		toolVersions = append(toolVersions, *release.TagName)
	}
	return toolVersions, nil
}

func generateClimanVersion(version string) (toolVersion ToolVersion) {
	baseUrl := "https://github.com/ghokun/climan/releases/download/" + version
	return ToolVersion{
		Version: version,
		Platforms: map[string]ToolDownload{
			"darwin_amd64": {
				Url:      baseUrl + "/climan-darwin-amd64",
				Checksum: baseUrl + "/climan-checksums.txt",
				Alg:      "sha256",
			},
			"darwin_arm64": {
				Url:      baseUrl + "/climan-darwin-arm64",
				Checksum: baseUrl + "/climan-checksums.txt",
				Alg:      "sha256",
			},
			"linux_386": {
				Url:      baseUrl + "/climan-linux-386",
				Checksum: baseUrl + "/climan-checksums.txt",
				Alg:      "sha256",
			},
			"linux_amd64": {
				Url:      baseUrl + "/climan-linux-amd64",
				Checksum: baseUrl + "/climan-checksums.txt",
				Alg:      "sha256",
			},
			"linux_arm": {
				Url:      baseUrl + "/climan-linux-arm",
				Checksum: baseUrl + "/climan-checksums.txt",
				Alg:      "sha256",
			},
			"linux_arm64": {
				Url:      baseUrl + "/climan-linux-arm64",
				Checksum: baseUrl + "/climan-checksums.txt",
				Alg:      "sha256",
			},
			"linux_ppc64le": {
				Url:      baseUrl + "/climan-linux-ppc64le",
				Checksum: baseUrl + "/climan-checksums.txt",
				Alg:      "sha256",
			},
			"linux_s390x": {
				Url:      baseUrl + "/climan-linux-s390x",
				Checksum: baseUrl + "/climan-checksums.txt",
				Alg:      "sha256",
			},
			"windows_386": {
				Url:      baseUrl + "/climan-windows-386.exe",
				Checksum: baseUrl + "/climan-checksums.txt",
				Alg:      "sha256",
			},
			"windows_amd64": {
				Url:      baseUrl + "/climan-windows-amd64.exe",
				Checksum: baseUrl + "/climan-checksums.txt",
				Alg:      "sha256",
			},
			"windows_arm": {
				Url:      baseUrl + "/climan-windows-arm.exe",
				Checksum: baseUrl + "/climan-checksums.txt",
				Alg:      "sha256",
			},
		},
	}
}
