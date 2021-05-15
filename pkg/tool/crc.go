package tool

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/ghokun/climan-runner/pkg/platform"
)

func init() {
	crc, err := getCrc()
	if err != nil {
		log.Fatal(err)
	}
	Tools[crc.Name] = crc
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
			Description: "Code Ready Containers - Local single node Openshift",
			Supports: platform.CalculateSupportedPlatforms(
				[]string{"darwin_amd64",
					"linux_amd64",
					"windows_amd64"}),
			Latest: data["version"].CrcVersion,
		}, nil
	}
	return crc, errors.New("error while fetcing latest version of crc")
}
