package main

import (
	"context"
	"github.com/ronaldoafonso/boxcmd/gcommand"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Countn't dial: %v.", err)
	}
	defer conn.Close()

	c := gcommand.NewRemoteCommandClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	cmd := gcommand.Command{
		Boxname: "788a20298f81.z3n.com.br",
		Command: "macs",
		Params:  "11:22:33:44:55:66 66:55:44:33:22:11",
	}
	rc, err := c.ExecCommand(ctx, &cmd)
	if err != nil {
		log.Fatalf("Error in ExecCommand: %v.", err)
	}
	log.Printf("Success: %v.", rc.GetReturnMsg())
}
