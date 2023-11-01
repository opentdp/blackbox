package args

import (
	"os"
	"strings"

	"github.com/opentdp/go-helper/logman"
)

var Version = "0.2.2"
var VersionNumber = "2"

var ConfigPath string
var ExecutablePath string

var Metadata map[string]string
var MetaServerListen = "127.0.0.1:9000"

var FrpClientConfig string
var PrometheusBlackboxConfig string

func init() {
	var err error

	// get ExecutablePath
	ExecutablePath, err = os.Executable()
	if err != nil {
		logman.Fatal("Find executable failed", "msg", err)
	}

	// get All ConfigPath
	ConfigPath = os.TempDir()
	if len(os.Args) > 1 && !strings.Contains(os.Args[1], "--") {
		ConfigPath = os.Args[1]
	}
	FrpClientConfig = ConfigPath + "/frp_client.toml"
	PrometheusBlackboxConfig = ConfigPath + "/prometheus_blackbox.yml"

	// get Node Metadata
	Metadata = map[string]string{
		"Name":   os.Getenv("NODE_NAME"),
		"Owner":  os.Getenv("NODE_OWNER"),
		"Region": os.Getenv("NODE_REGION"),
		"ISP":    os.Getenv("NODE_ISP"),
		"Banner": os.Getenv("NODE_BANNER"),
	}
}
