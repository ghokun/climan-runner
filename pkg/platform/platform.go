package platform

import (
	"errors"
	"runtime"
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

func CurrentPlatform() (platform int, err error) {
	current := runtime.GOOS + "_" + runtime.GOARCH
	if platform, ok := platforms[current]; ok {
		return platform, nil
	}
	return platform, errors.New("unsupported platform")
}
