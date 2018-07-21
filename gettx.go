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
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/btcjson"
	"log"
	"net/http"
)

func GetTxFromId(txid string) (tx *btcjson.TxRawResult) {
	hash, err := chainhash.NewHashFromStr(txid)
	if err != nil {
		log.Printf("Error getting tx hash %v\n", err)
	}
	tx, err = BtcdClient.GetRawTransactionVerbose(hash)
	if err != nil {
		log.Printf("Error finding tx %v\n", err)
	}
	log.Printf("Found tx %v\n", tx)
	return
}


func gettx(writer http.ResponseWriter,
	request *http.Request,
	params httprouter.Params) {

	query := params.ByName("query")
	log.Printf("in gettx...%s", query)

	tx := GetTxFromId(query)
	
	writer.Header().Set("Content-Type", "application/json")	
	json.NewEncoder(writer).Encode(tx)
}
