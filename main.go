package main

import (
	"log"
)

func main() {
	cmds := []Command{
		Command{
			"788a20298f81.z3n.com.br",
			"ssid",
			"'z3n'",
		},
		Command{
			"788a20298f81.z3n.com.br",
			"macs",
			"11:11:11:11:11:11 22:22:22:22:22:22",
		},
	}

	for _, cmd := range cmds {
		err := cmd.Exec()
		if err != nil {
			log.Printf("%v error: %v.", cmd.command, err)
		}
	}
}
