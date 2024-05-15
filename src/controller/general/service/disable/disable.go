package disable

import (
	"github.com/faradey/madock/src/controller/general/rebuild"
	"github.com/faradey/madock/src/controller/general/service"
	"github.com/faradey/madock/src/helper/cli/arg_struct"
	"github.com/faradey/madock/src/helper/cli/attr"
	"github.com/faradey/madock/src/helper/cli/fmtc"
	"github.com/faradey/madock/src/helper/configs"
)

func Execute() {
	args := attr.Parse(new(arg_struct.ControllerGeneralServiceDisable)).(*arg_struct.ControllerGeneralServiceDisable)

	if len(args.Args) == 0 {
		fmtc.ErrorLn("Service name(s) is required")
		return
	}
	for _, name := range args.Args {
		if service.IsService(name) {
			serviceName := service.GetByShort(name) + "/enabled"
			projectName := configs.GetProjectName()
			projectConfig := configs.GetProjectConfig(projectName)
			configs.SetParam(projectName, serviceName, "false", projectConfig["activeScope"], "")
			if args.Global {
				configs.SetParam(projectName, serviceName, "false", "default", configs.MainConfigCode)
			}
		}
	}

	rebuild.Execute()
}
