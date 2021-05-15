package tool

import "log"

func init() {
	climan, err := getCliman()
	if err != nil {
		log.Fatal(err)
	}
	Tools = append(Tools, climan)
}

func getCliman() (climan Tool, err error) {
	return getLatestReleaseFromGithub("ghokun", "climan", "climan", "Cloud tools cli manager", "darwin_amd64",
		"darwin_arm64",
		"linux_386",
		"linux_amd64",
		"linux_arm",
		"linux_arm64",
		"linux_ppc64le",
		"linux_s390x",
		"windows_386",
		"windows_amd64",
		"windows_arm64")
}
