package cmd

import (
	"os"
	"os/exec"
)

func FrpcStart(binPath, etcPath string) {
	cmd := exec.Command(binPath, "--config", etcPath+"/frpc.toml")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	cmd.Wait()
}

func ProberStart(binPath, etcPath string) {
	cmd := exec.Command(binPath, "--config.file", etcPath+"/prober.yml")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	cmd.Wait()
}
