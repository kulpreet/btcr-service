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
	"github.com/julienschmidt/httprouter"
	"time"	
	"log"
	"net/http"
)

type Result struct {
	Hrp string
	Magic int
	Height int
	Position int
	UtxoIndex int
}

var httpClient *http.Client

const (
    MaxIdleConnections int = 20
    RequestTimeout     int = 5
)

const (
	Username string = "jacktestnet"
	Password string =  "111111"
)
func init() {
    httpClient = createHTTPClient()
}

// createHTTPClient for connection re-use
func createHTTPClient() *http.Client {
    client := &http.Client{
        Transport: &http.Transport{
            MaxIdleConnsPerHost: MaxIdleConnections,
        },
        Timeout: time.Duration(RequestTimeout) * time.Second,
    }

    return client
}

func main() {
	router := httprouter.New()
    router.GET("/txref/:query/decode", decodeTxref)
    router.GET("/txref/:query/txid", getTxid)
	
    log.Fatal(http.ListenAndServe(":8080", router))	
}
