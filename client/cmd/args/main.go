package args

import (
	"os"

	"github.com/opentdp/go-helper/logman"
)

var Token string

var Version = "0.2.0"
var VersionNumber = "2"

var EtcDirectory string
var ExecutablePath string

var MetaServerListen = "127.0.0.1:9000"

func init() {
	var err error

	EtcDirectory = os.TempDir()

	if ExecutablePath, err = os.Executable(); err != nil {
		logman.Fatal("Find executable failed", "msg", err)
	}
}
