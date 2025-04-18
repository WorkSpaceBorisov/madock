package cli

import (
	"github.com/faradey/madock/src/helper/cli"
	"github.com/faradey/madock/src/helper/configs"
	"github.com/faradey/madock/src/helper/docker"
	"github.com/faradey/madock/src/helper/logger"
	"os"
	"os/exec"
)

func Execute() {
	flag := cli.NormalizeCliCommandWithJoin(os.Args[2:])
	projectName := configs.GetProjectName()
	projectConf := configs.GetCurrentProjectConfig()
	service := "php"
	if projectConf["platform"] == "pwa" {
		service = "nodejs"
	}

	service, user, workdir := cli.GetEnvForUserServiceWorkdir(service, "www-data", "")

	interactivePlusTTY := "-it"

	if os.Getenv("MADOCK_TTY_ENABLED") == "0" {
		interactivePlusTTY = "-i"
	}

	if workdir != "" {
		workdir = "cd " + workdir + " && "
	}

	cmd := exec.Command("docker", "exec", interactivePlusTTY, "-u", user, docker.GetContainerName(projectConf, projectName, service), "bash", "-c", workdir+flag)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		logger.Fatal(err)
	}
}
