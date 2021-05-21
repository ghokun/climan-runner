package tool

import (
	"log"
)

func init() {
	faascli, err := getFaascli()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, faascli)
	generateToolSpecificFiles("faas-cli", faascli.Latest, getFaascliVersions, generateFaascliVersion)
}

func getFaascli() (faascli Tool, err error) {
	return getLatestReleaseFromGithub("openfaas", "faas-cli", "faas-cli", "Cli for OpenFaaS",
		"darwin_amd64",
		"linux_amd64",
		"linux_arm",
		"linux_arm64",
		"windows_amd64")
}

func getFaascliVersions() (toolVersions []string, err error) {
	releases, err := getReleasesFromGithub("openfaas", "faas-cli", "faas-cli")
	if err != nil {
		return nil, err
	}
	for _, release := range releases {
		toolVersions = append(toolVersions, *release.TagName)
	}
	return toolVersions, nil
}

func generateFaascliVersion(version string) (toolVersion ToolVersion) {
	baseUrl := "https://github.com/openfaas/faas-cli/releases/download/" + version
	return ToolVersion{
		Version: version,
		Platforms: map[string]ToolDownload{
			"darwin_amd64": {
				Url:      baseUrl + "/faas-cli-darwin",
				Checksum: baseUrl + "/faas-cli-darwin.sha256",
				Alg:      "sha256",
			},
			"linux_amd64": {
				Url:      baseUrl + "/faas-cli",
				Checksum: baseUrl + "/faas-cli.sha256",
				Alg:      "sha256",
			},
			"linux_arm": {
				Url:      baseUrl + "/faas-cli-armhf",
				Checksum: baseUrl + "/faas-cli-armhf.sha256",
				Alg:      "sha256",
			},
			"linux_arm64": {
				Url:      baseUrl + "/faas-cli-arm64",
				Checksum: baseUrl + "/faas-cli-arm64.sha256",
				Alg:      "sha256",
			},
			"windows_amd64": {
				Url:      baseUrl + "/faas-cli.exe",
				Checksum: baseUrl + "/faas-cli.exe.sha256",
				Alg:      "sha256",
			},
		},
	}
}
