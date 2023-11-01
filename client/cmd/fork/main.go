package fork

import (
	"os"

	"github.com/opentdp/go-helper/logman"

	"github.com/opentdp/blackbox/client/module/capi"
	"github.com/opentdp/blackbox/client/module/glue"
)

func Checkin() {
	glueApps()

	if err := capi.Join(); err != nil {
		logman.Fatal("Fetch config failed", "msg", err)
	}

	go forkFrpClient()
	go forkPrometheusBlackbox()
}

func glueApps() {
	if len(os.Args) < 2 {
		return
	}
	if os.Args[1] == "--config" {
		glue.FrpClient()
		os.Exit(0)
	}
	if os.Args[1] == "--config.file" {
		glue.PrometheusBlackbox()
		os.Exit(0)
	}
}
