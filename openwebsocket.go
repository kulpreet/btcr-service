package main

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcutil"
)

const (
	username string = "jacktestnet"
	password string =  "111111"
	rpcURL = "host.opdup.com:18443"
	certFile = "rpc-opdup.cert"
)

var BtcdClient *rpcclient.Client

func openWebsocket() {
	// Connect to local btcd RPC server using websockets.
	btcdHomeDir := btcutil.AppDataDir("btcd", false)
	certs, err := ioutil.ReadFile(filepath.Join(btcdHomeDir, certFile))
	if err != nil {
		log.Fatal(err)
	}
	connCfg := &rpcclient.ConnConfig{
		Host:         rpcURL,
		Endpoint:     "ws",
		User:         username,
		Pass:         password,
		Certificates: certs,
	}
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}

	BtcdClient = client
}
