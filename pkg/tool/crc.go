package tool

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/ghokun/climan-runner/pkg/platform"
)

func init() {
	crc, err := getCrc()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, crc)
	generateToolSpecificFiles("crc", crc.Latest, getCrcVersions, generateCrcVersion)
}

type crcData struct {
	CrcVersion       string `json:"crcVersion,omitempty"`
	GitSha           string `json:"gitSha,omitempty"`
	OpenshiftVersion string `json:"openshiftVersion,omitempty"`
}

func getCrc() (crc Tool, err error) {
	response, err := http.Get("https://mirror.openshift.com/pub/openshift-v4/clients/crc/latest/release-info.json")
	if err != nil {
		return crc, err
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		var data map[string]crcData
		decoder := json.NewDecoder(response.Body)
		err = decoder.Decode(&data)
		if err != nil {
			return crc, err
		}
		return Tool{
			Name:        "crc",
			Description: "Local single node Openshift",
			Supports: platform.CalculateSupportedPlatforms(
				[]string{"darwin_amd64",
					"linux_amd64",
					"windows_amd64"}),
			Latest: data["version"].CrcVersion,
		}, nil
	}
	return crc, errors.New("error while fetcing latest version of crc")
}

// TODO get from mirror openshift
func getCrcVersions() (toolVersions []string, err error) {
	releases, err := getReleasesFromGithub("code-ready", "crc", "crc")
	if err != nil {
		return nil, err
	}
	for _, release := range releases {
		// Versions starting with 0 are not hosted on openshift mirror
		if !strings.HasPrefix(*release.TagName, "0.") {
			toolVersions = append(toolVersions, *release.TagName)
		}
	}
	return toolVersions, nil
}

func generateCrcVersion(version string) (toolVersion ToolVersion) {
	baseUrl := "https://mirror.openshift.com/pub/openshift-v4/clients/crc/" + strings.TrimPrefix(version, "v")
	return ToolVersion{
		Version: version,
		Platforms: map[string]ToolDownload{
			"darwin_amd64": {
				Url:      baseUrl + "/crc-macos-amd64.tar.xz",
				Checksum: baseUrl + "/sha256sum.txt",
				Alg:      "sha256",
			},
			"linux_amd64": {
				Url:      baseUrl + "/crc-linux-amd64.tar.xz",
				Checksum: baseUrl + "/sha256sum.txt",
				Alg:      "sha256"},
			"windows_amd64": {
				Url:      baseUrl + "/crc-windows-amd64.zip",
				Checksum: baseUrl + "/sha256sum.txt",
				Alg:      "sha256"},
		},
	}
}
