package main

import (
	"os/exec"
)

type Command struct {
	boxname    string
	command    string
	parameters string
}

func (c *Command) Scp() error {
	script := "/home/boxcmd/uci/"

	if c.command == "ssid" {
		script += "set_ssid.sh"
	} else if c.command == "macs" {
		script += "set_ipset_macs.sh"
	}

	scp := []string{
		script,
		c.boxname + ":/tmp",
	}
	_, err := exec.Command("scp", scp...).Output()
	return err
}

func (c *Command) Ssh() error {
	script := "/tmp/"

	if c.command == "ssid" {
		script += "set_ssid.sh"
	} else if c.command == "macs" {
		script += "set_ipset_macs.sh"
	}

	ssh := []string{
		c.boxname,
		script,
		c.parameters,
	}
	_, err := exec.Command("ssh", ssh...).Output()
	return err
}

func (c *Command) Exec() error {
	err := c.Scp()
	if err == nil {
		err = c.Ssh()
	}
	return err
}
