package main

import (
	"os"

	"github.com/opentdp/blackbox/client/cmd/args"
	"github.com/opentdp/blackbox/client/cmd/frpc"
	"github.com/opentdp/blackbox/client/cmd/power"
	"github.com/opentdp/blackbox/client/cmd/prober"
	"github.com/opentdp/go-helper/logman"
)

func main() {
	if len(os.Args) == 0 {
		logman.Fatal("Invalid args")
	}

	switch os.Args[1] {
	case "--config":
		frpc.Start()
	case "--config.file":
		prober.Start()
	default:
		args.Token = os.Args[1]
		if len(os.Args) > 2 {
			args.EtcDirectory = os.Args[2]
		}
		power.Start()
	}
}
