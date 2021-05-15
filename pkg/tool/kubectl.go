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
			Description: "Kubernetes command line interface",
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
