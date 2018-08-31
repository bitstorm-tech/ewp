package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
)

type config struct {
	ProxyHost  string
	ProxyPort  int
	ProxyUser  string
	Default    bool
	TTLSeconds int
}

var configFileName = ".proxy-manager.json"
var delim = byte('\n')
var configs = make(map[int]config)
var stdin = bufio.NewReader(os.Stdin)

func main() {
	c := config{}
	c.ProxyUser = readString("Proxy User (optional): ")
	c.ProxyHost = readString("Proxy Host: ")
	c.ProxyPort = readInt("Proxy Port (80): ", 80)
	c.Default = readBool("Default proxy (y/N): ", false)
	c.TTLSeconds = readInt("TTL in seconds (optional): ", 0)
	writeConfigToFile(c)
}

func readString(prompt string) string {
	fmt.Print(prompt)
	s, err := stdin.ReadString(delim)

	if err != nil {
		panic(err)
	}

	return strings.Replace(s, "\n", "", -1)
}

func readInt(prompt string, d int) int {
	s := readString(prompt)

	if len(s) == 0 {
		return d
	}

	i, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return i
}

func readBool(prompt string, d bool) bool {
	s := readString(prompt)

	if len(s) == 0 {
		return d
	}

	return strings.EqualFold(s, "y")
}

func getHomeDirectory() string {
	u, err := user.Current()
	if err != nil {
		fmt.Printf("Can't get user home directory, use current working directory (%v)\n", err)
		return "."
	}

	return u.HomeDir
}

func writeConfigToFile(c config) {
	f, err := os.Create("proxy-manager.json")
	defer f.Close()
	j, err := json.Marshal(c)
	f.Write(j)
	f.Sync()

	if err != nil {
		panic(err)
	}
}
