/* Copyright (c) 2018 Kulpreet Singh
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */

package main

import (
	"net/http"
	"log"
	"os"
	
	flags "github.com/jessevdk/go-flags"
	"github.com/julienschmidt/httprouter"
)

type Result struct {
	Hrp string
	Magic int
	Height int
	Position int
	UtxoIndex int
}

type config struct {
	Username              string        `short:"u" long:"username" description:"Username for BTCD connections"`
	Password              string        `short:"p" long:"password" description:"Password for BTCD connections"`
	BtcdConnect           string        `long:"btcdconnect" description:"Host and post to connect to btcd at"`
	BtcdCert              string        `long:"btcdcert" description:"File containing the certificate file"`
	Listen                string        `short:"l" long:"listen" description:"Port to listen on for HTTP"`
	Config                string        `short:"C" long:"config" description:"Ini Config file" required:"true"`	
}

var opts config

func loadConfig() {
	
	p := flags.NewParser(&opts, flags.Default)
	
    _, err := p.Parse()
    if err != nil {
        log.Println(err)
		os.Exit(1)
    }	

	err = flags.NewIniParser(p).ParseFile(opts.Config)
    if err != nil {
        log.Println(err)
    }
	
}

func main() {

	loadConfig()

	log.Printf("parsed: %v\n", opts)
	if opts.Username == "" ||
		opts.Password == "" ||
		opts.BtcdConnect == "" ||
		opts.Listen == "" {
		log.Println("Missing configuration, please check your config.ini file")
		os.Exit(1)
	}
	
	router := httprouter.New()
    router.GET("/txref/:query/decode", decodetxref)
    router.GET("/txref/:query/txid", txref2txid)
    router.GET("/tx/:query", gettx)
	router.GET("/addr/:query/spends", searchtransactions)
	router.GET("/txref/:query/resolve", resolvetodid)
	router.GET("/txref/:query/tip", tip)

	openWebsocket(opts)

	// Get the current block count.
	blockCount, err := BtcdClient.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Connection established, block count: %v\n", blockCount)	
	
    log.Fatal(http.ListenAndServe(":" + opts.Listen, router))	
}
