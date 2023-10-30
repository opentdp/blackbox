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
		"Name":   os.Getenv("NODE_NAME"),
		"Owner":  os.Getenv("NODE_OWNER"),
		"Region": os.Getenv("NODE_REGION"),
		"ISP":    os.Getenv("NODE_ISP"),
		"Banner": os.Getenv("NODE_BANNER"),
	}
}
