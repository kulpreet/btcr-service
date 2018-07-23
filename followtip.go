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
	"errors"
	"log"
	
	"github.com/btcsuite/btcd/btcjson"
)

type TipChainElement struct {
	Transaction *btcjson.SearchRawTransactionsResult
	InTipChain bool
}

func followTipFromTxid(txid string, spendsOnly bool) (tipchain []TipChainElement, err error) {
	tx := getTxFromTxid(txid)
	return followTipFromTx(tx, spendsOnly)
}

func isSpentBy(txid string, b *btcjson.SearchRawTransactionsResult) (uint32, bool) {
	for _, vin := range b.Vin {
		if vin.Txid == txid {
			return vin.Vout, true
		}				
	}
	return 0, false
}

/* A very simple loop to find transactions by blocktime.  Will only
work for Test #1 — BTCR DID MVP (Minimum Viable Product) Valid &
Unrevoked DID Document Example from
https://github.com/w3c-ccg/did-hackathon-2018

Note: txs are sorted in chronological order when received from btcd
*/
func findTipChainFromTxs(
	tx *btcjson.TxRawResult,
	txs []*btcjson.SearchRawTransactionsResult,
	spendsOnly bool) (
		tipchain []TipChainElement, err error) {
	
	if len(txs) == 0 {
		err = errors.New("No tip found")
		return
	}
	tipTxid := tx.Txid
	for _, relatedTx := range txs {
		element := TipChainElement{relatedTx, false}
		if spending, ok := isSpentBy(tipTxid, relatedTx); ok {
			log.Printf("txid %v is spent in %v, %v", tipTxid, relatedTx.Txid, spending)
			tipTxid = relatedTx.Txid
			element.InTipChain = true
			tipchain = append(tipchain, element)
		} else if !spendsOnly {
			tipchain = append(tipchain, element)			
		}
	}
	return
}

func followTipFromTx(tx *btcjson.TxRawResult, spendsOnly bool) (tipchain []TipChainElement, err error) {
	// follow the 0th vout for now
	// also assume there is only one address in the vout
	addr := tx.Vout[0].ScriptPubKey.Addresses[0]
	allTxs, err := searchForAddr(addr)
	if err != nil {
		log.Printf("Failed following tip, error finding address in tx vout %v\n", tx)
		log.Printf("Error %v\n", err)
		return
	}
	tipchain, err = findTipChainFromTxs(tx, allTxs, spendsOnly)
	if err != nil {
		log.Printf("Failed to find tip %v\n", err)
		return
	}
	return
}
