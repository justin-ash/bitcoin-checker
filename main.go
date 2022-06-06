package main

import (
	"bitcoin-checker/Config"
	"bitcoin-checker/Routes"
	"bitcoin-checker/Utils"
)

func main() {
	Config.DatabaseInit()
	go Utils.StartPolling()
	Routes.Init()
}
