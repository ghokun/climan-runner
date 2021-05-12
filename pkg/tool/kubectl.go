package tool

import (
	"errors"
	"io"
	"log"
	"net/http"
)

func init() {
	kubectl, err := getKubectl()
	if err != nil {
		log.Fatal(err)
	}
	Tools[kubectl.Name] = kubectl
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
			Supports:    2047,
			Latest:      string(bodyBytes),
		}, nil
	}
	return kubectl, errors.New("error while fetcing latest version of kubectl")
}
