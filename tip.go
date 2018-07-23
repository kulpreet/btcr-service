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
	"encoding/json"
	"log"
	"strconv"
	
	"github.com/julienschmidt/httprouter"
	txref "github.com/kulpreet/txref/util"
)

func tip(writer http.ResponseWriter,
	request *http.Request,
	params httprouter.Params) {

	var result = make(map[string]string)
	
	query := params.ByName("query")
	log.Printf("resolving to did...%s", query)

	spendsOnly := true

	queryString := request.URL.Query()
	spendsOnlyParam := queryString["spendsOnly"]
	if len(spendsOnlyParam) > 0 && spendsOnlyParam[0] == "false" {
		spendsOnly = false
	}

	_, _, Height, Position, UtxoIndex, err := txref.Decode(query)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
        return
	}

	blockHash, err := BtcdClient.GetBlockHash(int64(Height))
	if err != nil {
		log.Printf("Error finding blockhash %v\n", blockHash)
	}
	log.Printf("Found blockhash %v\n", blockHash)
	
	block, err := BtcdClient.GetBlockVerbose(blockHash)
	if err != nil {
		log.Printf("Error finding block %v\n", blockHash.String())
	}

	txid := block.Tx[Position]
	result["txid"] = txid
	result["utxo_index"] = strconv.Itoa(UtxoIndex)

	log.Printf("Found tx: %v", txid)
	
	tx := getTxFromTxid(txid)

	didAddrs := getDidAddressFromTx(tx)
	log.Printf("didAddrs: %v\n", didAddrs)
	
	// try to follow the tip
	tipchain, err := followTipFromTx(tx, spendsOnly)
	if err != nil {
		log.Printf("Error following tip %v\n", err)
	}

	// // is the tip spent
	// if tipSpent(tipchain) {
	// } else {
		
	// }		
	
	writer.Header().Set("Content-Type", "application/json")	
	json.NewEncoder(writer).Encode(tipchain)
}
