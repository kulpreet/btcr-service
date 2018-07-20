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
	"os"
	"log"
	"net/http"
	
	"github.com/julienschmidt/httprouter"
	flags "github.com/jessevdk/go-flags"
)

type Result struct {
	Hrp string
	Magic int
	Height int
	Position int
	UtxoIndex int
}

type config struct {
	Username              string        `short:"u" long:"username" description:"Username for BTCD connections" required:"true"`
	Password              string        `short:"p" long:"password" description:"Password for BTCD connections" required:"true"`
	BtcdConnect           string        `long:"btcdconnect" description:"Host and post to connect to btcd at" required:"true"`
	BtcdCert              string        `long:"btcdcert" description:"File containing the certificate file" required:"true"`
}
var opts config

func init() {
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}
}

func main() {	
	router := httprouter.New()
    router.GET("/txref/:query/decode", decodetxref)
    router.GET("/txref/:query/txid", txref2txid)
    router.GET("/tx/:query", gettx)
	router.GET("/addr/:query/spends", searchtransactions)

	openWebsocket(opts)

	// Get the current block count.
	blockCount, err := BtcdClient.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Connection established, block count: %v\n", blockCount)	
	
    log.Fatal(http.ListenAndServe(":8080", router))	
}
