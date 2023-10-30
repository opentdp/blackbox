package args

import (
	"os"

	"github.com/opentdp/go-helper/logman"
)

var Version = "0.2.0"
var VersionNumber = "2"

var EtcDirectory string
var ExecutablePath string

var NodeMetaInfo map[string]string

var MetaServerListen = "127.0.0.1:9000"

func init() {
	var err error

	// get EtcDirectory
	EtcDirectory = os.TempDir()

	// get ExecutablePath
	ExecutablePath, err = os.Executable()
	if err != nil {
		logman.Fatal("Find executable failed", "msg", err)
	}

	// get NodeMetaInfo
	NodeMetaInfo = map[string]string{
		"name":   os.Getenv("NODE_NAME"),
		"owner":  os.Getenv("NODE_OWNER"),
		"region": os.Getenv("NODE_REGION"),
		"isp":    os.Getenv("NODE_ISP"),
		"banner": os.Getenv("NODE_BANNER"),
	}
}
