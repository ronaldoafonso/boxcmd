package main

import (
	"context"
	"github.com/ronaldoafonso/boxcmd/gcommand"
	"log"
)

var server = gcommand.RemoteCommandService{
	ExecCommand: ExecCommand,
}

func ExecCommand(ctx context.Context, in *gcommand.Command) (*gcommand.ReturnMsg, error) {
	cmd := Command{
		boxname:    in.GetBoxname(),
		command:    in.GetCommand(),
		parameters: in.GetParams(),
	}

	msg := "Ok"
	err := cmd.Exec()
	if err != nil {
		msg = err.Error()
	}

	log.Printf("ExecCommand(%v): %v.", cmd, msg)
	return &gcommand.ReturnMsg{ReturnMsg: msg}, err
}
