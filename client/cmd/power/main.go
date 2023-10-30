package power

import (
	"os"
	"os/exec"

	"github.com/opentdp/blackbox/client/api"
	"github.com/opentdp/blackbox/client/cmd/args"
	"github.com/opentdp/go-helper/logman"
)

func Start() {
	if err := api.Join(); err != nil {
		logman.Fatal("Fetch config failed", "msg", err)
	}

	go forkFrps()
	go forkProber()

	metaServer()
}

func forkFrps() {
	etc := args.EtcDirectory + "/frpc.toml"
	cmd := exec.Command(args.ExecutablePath, "--config", etc)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr

	if err := cmd.Start(); err != nil {
		logman.Fatal("Start frpc failed", "msg", err)
	}
	cmd.Wait()
}

func forkProber() {
	etc := args.EtcDirectory + "/prober.yml"
	cmd := exec.Command(args.ExecutablePath, "--config.file", etc)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr

	if err := cmd.Start(); err != nil {
		logman.Fatal("Start prober failed", "msg", err)
	}
	cmd.Wait()
}
