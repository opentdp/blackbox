package api

import (
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
)

func Join(token, etcPath string) error {
	body, err := apiPost("/api/join", "2;"+token)
	if err != nil {
		return err
	}
	if !strings.Contains(body, "===================") {
		return errors.New("Join cluster failed\n" + body)
	}
	config := strings.Split(body, "===================")
	err = writeFile(etcPath+"/frpc.toml", config[0])
	if err != nil {
		return err
	}
	err = writeFile(etcPath+"/prober.yml", config[1])
	if err != nil {
		return err
	}
	return nil
}

func apiPost(url, textData string) (string, error) {
	url = "https://blackbox.opentdp.org" + url
	encodedData := base64.StdEncoding.EncodeToString([]byte(textData))
	resp, err := http.Post(url, "text/plain", strings.NewReader(encodedData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func writeFile(filePath, content string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}
