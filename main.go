package main

import (
	"peerchat/p2p"
	"github.com/sirupsen/logrus"
)

func main() {
	// Create a new P2PHost
	h := p2p.NewP2P()
	for _, add := range h.Host.Addrs(){
		logrus.Infoln(add)
	}
	logrus.Infoln(h.Host.ID())
}
