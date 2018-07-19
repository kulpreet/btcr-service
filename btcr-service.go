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
	"log"
	"net/http"

	txref "github.com/kulpreet/txref/util"
)

type Result struct {
	Hrp string
	Magic int
	Height int
	Position int
	UtxoIndex int
}

func decodeTxref(writer http.ResponseWriter,
	request *http.Request,
	params httprouter.Params) {

	query := params.ByName("query")
	log.Printf("in handler...%s", query)

	Hrp, Magic, Height, Position, UtxoIndex, err := txref.Decode(query)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
        return
	}

	result := Result{Hrp, Magic, Height, Position, UtxoIndex}
	log.Printf("Decoded as %v", result)
	writer.Header().Set("Content-Type", "application/json")	
	json.NewEncoder(writer).Encode(result)
}

func main() {
	router := httprouter.New()
    router.GET("/txref/:query/decode", decodeTxref)
	
    log.Fatal(http.ListenAndServe(":8080", router))
}