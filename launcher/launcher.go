package launcher

import (
	"os/exec"

	"github.com/dadrian/detour/config"
	"github.com/sirupsen/logrus"
)

func BuildLaunchCallback(browser *config.BrowserDefinition) func() {
	profile := browser.Profile
	return func() {
		LaunchFirefox(profile)
	}
}

func LaunchFirefox(profileName string) *exec.Cmd {
	cmd := exec.Command("firefox", "-P", profileName)
	logrus.Debugf("launching %s", cmd)
	cmd.Start()
	cmd.Process.Release()
	logrus.Debugf("launched process pid: %d", cmd.Process.Pid)
	return cmd
}
