package main

import (
	"fmt"
	"os"
	"os/user"
)

var configFileName = ".proxy-manager.conf"

type config struct {
	proxyHost string
	proxyPort int
	proxyUser string
}

func main() {
	fmt.Printf("Home directory: %v\n", getHomeDirectory())
}

func createDefaultConfig() (config, error) {
	return config{}, nil
}

func getHomeDirectory() string {
	u, err := user.Current()
	if err != nil {
		fmt.Println("Can't find user home directory, use working directory")
		return "."
	}

	return u.HomeDir
}

func readConfigFile() config {
	c := getHomeDirectory() + configFileName
	_, err := os.OpenFile(c, os.O_RDWR, 0755)
	if err != nil {
		fmt.Println("No config file found, create a default one")
		if d, err := createDefaultConfig(); err == nil {
			return d
		}
	}

	return config{}
}
