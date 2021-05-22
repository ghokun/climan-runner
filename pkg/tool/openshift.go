package tool

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/ghokun/climan-runner/pkg/platform"
)

func init() {
	oc, openshiftinstall, err := getOpenshiftTools()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, oc)
	Tools = append(Tools, openshiftinstall)
}

func getOpenshiftTools() (oc Tool, openshiftinstall Tool, err error) {

	url := "https://mirror.openshift.com/pub/openshift-v4/clients/ocp/"
	latest, err := getLatestVersion(url+"stable/", "openshift-client-linux-", ".tar.gz")

	oc.Name = "oc"
	oc.Description = "Openshift command line interface"
	oc.Supports = platform.CalculateSupportedPlatforms([]string{
		"darwin_amd64",
		"linux_amd64",
		"linux_ppc64le",
		"linux_s390x",
		"windows_amd64",
	})
	oc.Latest = latest
	oc.GetVersions = func() ([]string, error) {
		return getVersionsFromMirrorOpenshift("oc", url)
	}
	oc.GenerateVersion = func(version string) (toolVersion ToolVersion) {
		baseUrl := "https://mirror.openshift.com/pub/openshift-v4/"
		alg := "sha256"
		return ToolVersion{
			Version: version,
			Platforms: map[string]ToolDownload{
				"darwin_amd64": {
					Url:      baseUrl + "clients/ocp/" + version + "/openshift-client-mac.tar.gz",
					Checksum: baseUrl + "clients/ocp/" + version + "/sha256sum.txt",
					Alg:      alg,
				},
				"linux_amd64": {
					Url:      baseUrl + "clients/ocp/" + version + "/openshift-client-linux.tar.gz",
					Checksum: baseUrl + "clients/ocp/" + version + "/sha256sum.txt",
					Alg:      alg,
				},
				"linux_ppc64le": {
					Url:      baseUrl + "ppc64le/clients/ocp/" + version + "/openshift-client-linux.tar.gz",
					Checksum: baseUrl + "ppc64le/clients/ocp/" + version + "/sha256sum.txt",
					Alg:      alg,
				},
				"linux_s390x": {
					Url:      baseUrl + "s390x/clients/ocp/" + version + "/openshift-client-linux.tar.gz",
					Checksum: baseUrl + "s390x/clients/ocp/" + version + "/sha256sum.txt",
					Alg:      alg,
				},
				"windows_amd64": {
					Url:      baseUrl + "clients/ocp/" + version + "/openshift-client-windows.tar.gz",
					Checksum: baseUrl + "clients/ocp/" + version + "/sha256sum.txt",
					Alg:      alg,
				},
			},
		}
	}

	openshiftinstall.Name = "openshift-install"
	openshiftinstall.Description = "Openshift installer"
	openshiftinstall.Supports = platform.CalculateSupportedPlatforms([]string{
		"darwin_amd64",
		"linux_amd64",
		"linux_ppc64le",
		"linux_s390x",
	})
	openshiftinstall.Latest = latest
	openshiftinstall.GetVersions = func() ([]string, error) {
		return getVersionsFromMirrorOpenshift("openshift-install", url)
	}
	openshiftinstall.GenerateVersion = func(version string) (toolVersion ToolVersion) {
		baseUrl := "https://mirror.openshift.com/pub/openshift-v4/"
		alg := "sha256"
		return ToolVersion{
			Version: version,
			Platforms: map[string]ToolDownload{
				"darwin_amd64": {
					Url:      baseUrl + "clients/ocp/" + version + "/openshift-install-mac.tar.gz",
					Checksum: baseUrl + "clients/ocp/" + version + "/sha256sum.txt",
					Alg:      alg,
				},
				"linux_amd64": {
					Url:      baseUrl + "clients/ocp/" + version + "/openshift-install-linux.tar.gz",
					Checksum: baseUrl + "clients/ocp/" + version + "/sha256sum.txt",
					Alg:      alg,
				},
				"linux_ppc64le": {
					Url:      baseUrl + "ppc64le/clients/ocp/" + version + "/openshift-install-linux.tar.gz",
					Checksum: baseUrl + "ppc64le/clients/ocp/" + version + "/sha256sum.txt",
					Alg:      alg,
				},
				"linux_s390x": {
					Url:      baseUrl + "s390x/clients/ocp/" + version + "/openshift-install-linux.tar.gz",
					Checksum: baseUrl + "s390x/clients/ocp/" + version + "/sha256sum.txt",
					Alg:      alg,
				},
			},
		}
	}
	return oc, openshiftinstall, err
}

func getLatestVersion(url string, prefix string, suffix string) (version string, err error) {
	res, err := http.Get(url)
	if err != nil {
		return version, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return version, fmt.Errorf("error while scraping latest version of openshift tools")
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return version, fmt.Errorf("error while scraping %q", url)
	}
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		if text, ok := s.Attr("href"); ok {
			if len(text) > 0 && strings.HasPrefix(text, prefix) {
				version = strings.TrimSuffix(strings.TrimPrefix(text, prefix), suffix)
			}
		}
	})
	return version, err
}
