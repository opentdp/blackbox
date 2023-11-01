package capi

import (
	"encoding/base64"
	"os"

	"github.com/opentdp/go-helper/request"

	"github.com/opentdp/blackbox/client/cmd/args"
)

func openApi(url, body string) (string, error) {
	url = "https://blackbox.opentdp.org/api" + url
	data := base64.StdEncoding.EncodeToString([]byte(body))
	return request.TextPost(url, data, request.H{"TDP-Blackbox-Version": args.Version})
}

func writeFile(filePath, content string) error {
	return os.WriteFile(filePath, []byte(content), 0644)
}
