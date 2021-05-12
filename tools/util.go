package tools

import (
	"bytes"
	"errors"
	"runtime"
	"strings"
	"text/template"
)

var platforms = map[string]int{
	"darwin_amd64":  1,
	"darwin_arm64":  2,
	"linux_386":     4,
	"linux_amd64":   8,
	"linux_arm":     16,
	"linux_arm64":   32,
	"linux_ppc64le": 64,
	"linux_s390x":   128,
	"windows_386":   256,
	"windows_amd64": 512,
	"windows_arm64": 1024,
}

func currentPlatform() (platform int, err error) {
	current := runtime.GOOS + "_" + runtime.GOARCH
	if platform, ok := platforms[current]; ok {
		return platform, nil
	}
	return platform, errors.New("unsupported platform")
}

func checkToolSupport(supports int) (result bool, err error) {
	platform, err := currentPlatform()
	if err != nil {
		return false, err
	}
	return platform == platform&supports, nil
}

type templateVariables struct {
	Owner      string
	Repository string
	Name       string
	Os         string
	Arch       string
	Suffix     string
	Version    string
}

func parseDownloadTemplate(data interface{}, rawTemplate string) (parsedTemplate string, err error) {
	var tpl bytes.Buffer
	tmpl, err := template.New("template").Parse(rawTemplate)
	if err != nil {
		return parsedTemplate, err
	}
	if err = tmpl.Execute(&tpl, data); err != nil {
		return parsedTemplate, err
	}
	return tpl.String(), nil
}

func GenerateCharacter(numberOfCharacter int, character string) string {
	var sb strings.Builder
	for i := 0; i < numberOfCharacter; i++ {
		sb.WriteString(character)
	}
	return sb.String()
}
