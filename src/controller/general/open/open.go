package open

import (
	"github.com/faradey/madock/src/helper/cli/arg_struct"
	"github.com/faradey/madock/src/helper/cli/attr"
	"github.com/faradey/madock/src/helper/configs"
	"github.com/faradey/madock/src/helper/logger"
	"os/exec"
	"runtime"
)

func Execute() {
	args := attr.Parse(new(arg_struct.ControllerGeneralOpen)).(*arg_struct.ControllerGeneralOpen)

	projectConfig := configs.GetCurrentProjectConfig()
	hosts := configs.GetHosts(projectConfig)
	var cmd string
	var argsCommand []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		argsCommand = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	host := "https://" + hosts[0]["name"]
	if args.Service != "" {
		host = host + "/" + args.Service
	}
	argsCommand = append(argsCommand, host)
	err := exec.Command(cmd, argsCommand...).Start()
	if err != nil {
		logger.Fatal(err)
	}
}
