package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
	"syscall"
)

type config struct {
	ProxyHost  string
	ProxyPort  int
	ProxyUser  string
	Default    bool
	TTLSeconds int
}

var configFileName = "ewp-config.json"
var delim = byte('\n')
var configs = make(map[int]config)
var stdin = bufio.NewReader(os.Stdin)
var create = flag.Bool("c", false, "create a new config")
var help = flag.Bool("h", false, "show this help")

func main() {
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	if *create {
		c := createConfigFromStdin()
		writeConfigToFile(c)
		os.Exit(0)
	}

	execCommand()
}

func readStringFromStdin(prompt string) string {
	fmt.Print(prompt)

	s, err := stdin.ReadString(delim)
	if err != nil {
		panic(err)
	}

	return strings.Replace(s, "\n", "", -1)
}

func readIntFromStdin(prompt string, d int) int {
	s := readStringFromStdin(prompt)
	if len(s) == 0 {
		return d
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

func readBoolFromStdin(prompt string, d bool) bool {
	s := readStringFromStdin(prompt)
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
	f, err := os.Create(configFileName)
	defer f.Close()
	j, err := json.Marshal(c)
	f.Write(j)
	f.Sync()

	if err != nil {
		panic(err)
	}
}

func createConfigFromStdin() config {
	c := config{}
	c.ProxyUser = readStringFromStdin("Proxy User (optional): ")
	c.ProxyHost = readStringFromStdin("Proxy Host: ")
	c.ProxyPort = readIntFromStdin("Proxy Port (80): ", 80)
	c.Default = readBoolFromStdin("Default proxy (y/N): ", false)
	c.TTLSeconds = readIntFromStdin("TTL in seconds (optional): ", 0)

	return c
}

func readConfigFromFile() config {
	d, err := ioutil.ReadFile(configFileName)
	if err != nil {
		fmt.Printf("No %s file found, please create one first\n", configFileName)
		os.Exit(0)
	}

	c := new(config)
	err = json.Unmarshal(d, c)
	if err != nil {
		panic(err)
	}

	return *c
}

func setEnvironment(password string) {
	c := readConfigFromFile()
	var p string

	if len(c.ProxyUser) > 0 {
		p = c.ProxyUser
	}

	if len(p) > 0 && len(password) > 0 {
		p += ":" + password
	}

	if len(p) > 0 {
		p += "@"
	}

	p += c.ProxyHost + ":" + fmt.Sprint(c.ProxyPort)

	os.Setenv("HTTP_PROXY", p)
	os.Setenv("HTTPS_PROXY", p)
}

func execCommand() {
	password := readStringFromStdin("Proxy password (optional): ")
	setEnvironment(password)
	args := flag.Args()
	if len(args) == 0 {
		return
	}
	exec, err := exec.LookPath(args[0])
	if err != nil {
		panic(err)
	}
	env := os.Environ()
	syscall.Exec(exec, args, env)
}
