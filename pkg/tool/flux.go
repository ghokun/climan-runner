package tool

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	flux, err := getFlux()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, flux)
	// Generate flux specific directory
	folder := filepath.Join(".", "docs", flux.Name)
	os.Mkdir(folder, os.ModePerm)
	// Generate versions.json
	toolVersions, err := getFluxVersions()
	if err != nil {
		log.Fatal(err)
	}
	err = writeJson(folder, "versions.json", toolVersions)
	if err != nil {
		log.Fatal(err)
	}
	// Generate template.json
	template := generateFluxVersion("{{.Version}}")
	err = writeJson(folder, "template.json", template)
	if err != nil {
		log.Fatal(err)
	}
	// Generate latest.json
	latest := generateFluxVersion(flux.Latest)
	err = writeJson(folder, "latest.json", latest)
	if err != nil {
		log.Fatal(err)
	}
}

func getFlux() (flux Tool, err error) {
	return getLatestReleaseFromGithub("fluxcd", "flux2", "flux", "Cli for Flux",
		"darwin_amd64",
		"darwin_arm64",
		"linux_amd64",
		"linux_arm",
		"linux_arm64",
		"windows_386",
		"windows_amd64")
}

func getFluxVersions() (toolVersions []string, err error) {
	releases, err := getReleasesFromGithub("fluxcd", "flux2", "flux")
	if err != nil {
		return nil, err
	}
	for _, release := range releases {
		toolVersions = append(toolVersions, *release.TagName)
	}
	return toolVersions, nil
}

func generateFluxVersion(version string) (toolVersion ToolVersion) {
	withOutV := strings.TrimPrefix(version, "v")
	baseUrl := "https://github.com/fluxcd/flux2/releases/download/" + version + "/flux_" + withOutV
	return ToolVersion{
		Version: version,
		Platforms: map[string]ToolDownload{
			"darwin_amd64": {
				Url:      baseUrl + "_darwin_amd64.tar.gz",
				Checksum: baseUrl + "_checksums.txt",
				Alg:      "sha256",
			},
			"darwin_arm64": {
				Url:      baseUrl + "_darwin_arm64.tar.gz",
				Checksum: baseUrl + "_checksums.txt",
				Alg:      "sha256",
			},
			"linux_amd64": {
				Url:      baseUrl + "_linux_amd64.tar.gz",
				Checksum: baseUrl + "_checksums.txt",
				Alg:      "sha256",
			},
			"linux_arm": {
				Url:      baseUrl + "_linux_arm.tar.gz",
				Checksum: baseUrl + "_checksums.txt",
				Alg:      "sha256",
			},
			"linux_arm64": {
				Url:      baseUrl + "_linux_arm64.tar.gz",
				Checksum: baseUrl + "_checksums.txt",
				Alg:      "sha256",
			},
			"windows_amd64": {
				Url:      baseUrl + "_windows_amd64.tar.gz",
				Checksum: baseUrl + "_checksums.txt",
				Alg:      "sha256",
			},
			"windows_386": {
				Url:      baseUrl + "_windows_386.tar.gz",
				Checksum: baseUrl + "_checksums.txt",
				Alg:      "sha256",
			},
		},
	}
}
