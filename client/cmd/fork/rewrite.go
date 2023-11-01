package fork

import (
	"os"
	"os/exec"

	"github.com/opentdp/go-helper/logman"

	"github.com/opentdp/blackbox/client/cmd/args"
)

func forkFrpClient() {
	etc := args.FrpClientConfig
	cmd := exec.Command(args.ExecutablePath, "--config", etc)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr

	if err := cmd.Start(); err != nil {
		logman.Fatal("Start frpc failed", "msg", err)
	}
	cmd.Wait()
}

func forkPrometheusBlackbox() {
	etc := args.PrometheusBlackboxConfig
	cmd := exec.Command(args.ExecutablePath, "--config.file", etc)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr

	if err := cmd.Start(); err != nil {
		logman.Fatal("Start prometheus blackbox failed", "msg", err)
	}
	cmd.Wait()
}
