package main

import (
	"github.com/manishmeganathan/peerchat/src"
	"github.com/sirupsen/logrus"
)

func main() {
	// Create a new P2PHost
	p2phost := src.NewP2P()
	logrus.Infoln(p2phost.Host.Addrs())
	logrus.Infoln(p2phost.Host.ID())
}
