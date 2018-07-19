package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
	
	txref "github.com/kulpreet/txref/util"
)

func txref2txid(writer http.ResponseWriter,
	request *http.Request,
	params httprouter.Params) {

	var result = make(map[string]string)
	
	query := params.ByName("query")
	log.Printf("in decodeTxref...%s", query)

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

	writer.Header().Set("Content-Type", "application/json")	
	json.NewEncoder(writer).Encode(result)
}
