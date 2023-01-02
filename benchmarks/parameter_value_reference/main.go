package main

import (
	"config-struct/conf"
	"fmt"
)

const (
	configFileEnvVariableName = "CONFIG_FILE"
	defaultConfigFileName     = "config"
)

func main() {
	cfg, err := conf.LoadConfiguration("", "")
	if err != nil {
		panic(err)

	}
	fmt.Println(cfg)
}
