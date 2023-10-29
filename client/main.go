package main

import (
	"os"

	"github.com/opentdp/blackbox/client/api"
	"github.com/opentdp/blackbox/client/cmd"
	"github.com/opentdp/blackbox/client/cmd/frpc"
	"github.com/opentdp/blackbox/client/cmd/prober"
)

func main() {
	token := "2;"
	if len(os.Args) > 1 {
		token = os.Args[1]
	}
	etcPath := os.TempDir()
	if len(os.Args) > 2 {
		etcPath = os.Args[2]
	}
	if token == "--config" {
		frpc.Start()
		return
	}
	if token == "--config.file" {
		prober.Start()
		return
	}
	if err := api.Join(token, etcPath); err != nil {
		panic(err)
	}
	if binPath, err := os.Executable(); err == nil {
		go cmd.FrpcStart(binPath, etcPath)
		cmd.ProberStart(binPath, etcPath)
	} else {
		panic(err)
	}
}
