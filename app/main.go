package main

import (
	"BE-ASSIGMENT/app/cmd"
	"github.com/labstack/gommon/log"
)

func main() {
	log.Infof("starting application..........")
	cmd.Execute()
}
