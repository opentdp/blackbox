package main

import (
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"blackbox-node/cmd/frpc"
	"blackbox-node/cmd/prober"
)

func main() {
	token := "2;"
	if len(os.Args) > 1 {
		token = os.Args[1]
	}
	etcPath := os.TempDir()
	if len(os.Args) > 2 {
		etcPath = os.Args[2]
	}
	if token == "--config" {
		frpc.Start()
		return
	}
	if token == "--config.file" {
		prober.Start()
		return
	}
	if err := joinCluster(token, etcPath); err != nil {
		panic(err)
	}
	if binPath, err := os.Executable(); err == nil {
		go frpcStart(binPath, etcPath)
		proberStart(binPath, etcPath)
	} else {
		panic(err)
	}
}

func joinCluster(token, etcPath string) error {
	body, err := apiPost("/api/join", "2;"+token)
	if err != nil {
		return err
	}
	if !strings.Contains(body, "===================") {
		return errors.New("join cluster failed\n" + body)
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

func frpcStart(binPath, etcPath string) {
	cmd := exec.Command(binPath, "--config", etcPath+"/frpc.toml")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	cmd.Wait()
}

func proberStart(binPath, etcPath string) {
	cmd := exec.Command(binPath, "--config.file", etcPath+"/prober.yml")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	cmd.Wait()
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
