package api

import (
	"encoding/base64"
	"os"

	"github.com/opentdp/blackbox/client/cmd/args"
	"github.com/opentdp/go-helper/request"
)

func httpPost(url, body string) (string, error) {
	url = "https://blackbox.opentdp.org" + url
	data := base64.StdEncoding.EncodeToString([]byte(body))
	return request.TextPost(url, data, request.H{"TDP-Blackbox-Version": args.Version})
}

func fileWrite(filePath, content string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err = file.WriteString(content); err != nil {
		return err
	}
	return nil
}
