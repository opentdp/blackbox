package main

import (
	"os"

	"github.com/opentdp/blackbox/client/cmd/args"
	"github.com/opentdp/blackbox/client/cmd/frpc"
	"github.com/opentdp/blackbox/client/cmd/power"
	"github.com/opentdp/blackbox/client/cmd/prober"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--config":
			frpc.Start()
			return
		case "--config.file":
			prober.Start()
			return
		default:
			args.EtcDirectory = os.Args[1]
		}
	}
	power.Start()
}
