package launcher

import (
	"os/exec"

	"github.com/dadrian/detour/config"
	"github.com/sirupsen/logrus"
)

func BuildLaunchCallback(browser *config.BrowserDefinition, url string) func() {
	profile := browser.Profile
	return func() {
		LaunchFirefox(profile, url)
	}
}

func LaunchFirefox(profileName, url string) *exec.Cmd {
	cmd := exec.Command("firefox", "-P", profileName, url)
	logrus.Debugf("launching %s", cmd)
	cmd.Start()
	cmd.Process.Release()
	logrus.Debugf("launched process pid: %d", cmd.Process.Pid)
	return cmd
}
