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
}

type crcData struct {
	CrcVersion       string `json:"crcVersion,omitempty"`
	GitSha           string `json:"gitSha,omitempty"`
	OpenshiftVersion string `json:"openshiftVersion,omitempty"`
}

func getCrc() (crc Tool, err error) {
	name := "crc"
	url := "https://mirror.openshift.com/pub/openshift-v4/clients/crc/"
	response, err := http.Get(url + "latest/release-info.json")
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
			Name:        name,
			Description: "Local single node Openshift",
			Supports: platform.CalculateSupportedPlatforms(
				[]string{"darwin_amd64",
					"linux_amd64",
					"windows_amd64"}),
			Latest: data["version"].CrcVersion,
			GetVersions: func() ([]string, error) {
				return getVersionsFromMirrorOpenshift(name, url)
			},
			GenerateVersion: func(version string) (toolVersion ToolVersion) {
				baseUrl := url + strings.TrimPrefix(version, "v")
				checksum := baseUrl + "/sha256sum.txt"
				alg := "sha256"
				return ToolVersion{
					Version: version,
					Platforms: map[string]ToolDownload{
						"darwin_amd64": {
							Url:      baseUrl + "/crc-macos-amd64.tar.xz",
							Checksum: checksum,
							Alg:      alg,
						},
						"linux_amd64": {
							Url:      baseUrl + "/crc-linux-amd64.tar.xz",
							Checksum: checksum,
							Alg:      alg,
						},
						"windows_amd64": {
							Url:      baseUrl + "/crc-windows-amd64.zip",
							Checksum: checksum,
							Alg:      alg,
						},
					},
				}
			},
		}, nil
	}
	return crc, errors.New("error while fetcing latest version of crc")
}
