package glue

import (
	"github.com/opentdp/go-helper/logman"

	"github.com/opentdp/blackbox/client/module/glue/frp_client"
	"github.com/opentdp/blackbox/client/module/glue/prometheus_blackbox"
)

func FrpClient() {
	frp_client.Execute()
	logman.Info("frp client exited")
}

func PrometheusBlackbox() {
	prometheus_blackbox.Execute()
	logman.Info("prometheus blackbox exited")
}
