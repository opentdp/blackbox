package fork

import (
	"os"
	"os/exec"

	"github.com/opentdp/go-helper/logman"

	"github.com/opentdp/blackbox/client/cmd/args"
	"github.com/opentdp/blackbox/client/module/capi"
	"github.com/opentdp/blackbox/client/module/glue"
)

func Checkin() {
	if len(os.Args) > 1 {
		glue.Start(os.Args[1])
	}

	if err := capi.Join(); err != nil {
		logman.Fatal("Fetch config failed", "msg", err)
	}

	go forkFrpClient()
	go forkPrometheusBlackbox()
}

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
