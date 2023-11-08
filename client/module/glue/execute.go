package glue

import (
	"os"

	"github.com/opentdp/go-helper/logman"

	"github.com/opentdp/blackbox/client/module/glue/frp_client"
	"github.com/opentdp/blackbox/client/module/glue/prometheus_blackbox"
)

func Start(arg1 string) {
	switch arg1 {
	case "--config":
		frpClient()
	case "--config.file":
		prometheusBlackbox()
	}
}

func frpClient() {
	frp_client.Execute()
	logman.Info("frp client exited")
	os.Exit(0)
}

func prometheusBlackbox() {
	prometheus_blackbox.Execute()
	logman.Info("prometheus blackbox exited")
	os.Exit(0)
}
