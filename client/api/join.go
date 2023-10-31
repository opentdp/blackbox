package api

import (
	"errors"
	"strings"

	"github.com/opentdp/blackbox/client/cmd/args"
)

func Join() error {
	token := []string{
		args.VersionNumber,
		args.Metadata["Name"],
		args.Metadata["Owner"],
		args.Metadata["Region"],
		args.Metadata["ISP"],
		args.Metadata["Banner"],
	}

	body, err := httpPost("/api/join", strings.Join(token, ";"))
	if err != nil {
		return err
	}

	config := strings.Split(body, "===================")
	if len(config) < 2 {
		return errors.New(body)
	}

	config[0] = strings.Trim(config[0], "\n")
	err = fileWrite(args.EtcDirectory+"/frpc.toml", config[0])
	if err != nil {
		return err
	}

	config[1] = strings.Trim(config[1], "\n")
	err = fileWrite(args.EtcDirectory+"/prober.yml", config[1])
	if err != nil {
		return err
	}

	return nil
}
