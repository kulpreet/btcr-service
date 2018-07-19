package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"

	"github.com/btcsuite/btcd/btcjson"	
	txref "github.com/kulpreet/txref/util"
)

func getTxid(writer http.ResponseWriter,
	request *http.Request,
	params httprouter.Params) {

	query := params.ByName("query")
	log.Printf("in getTxid...%s", query)
	
	_, _, Height, _, _, err := txref.Decode(query)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
        return
	}

	cmd := btcjson.NewGetBlockHashCmd(int64(Height))
	marshalled, err := btcjson.MarshalCmd(100, cmd)
	if err != nil {
		log.Fatalf("Error marshalling command %s", err)
	}

	log.Printf("Marshalled command %s", marshalled)

	jsonResponse, err := sendCommand(marshalled)
	if err != nil {
		log.Printf("Error sending command %s\n", err)
	}

	var blockHash string
	if err = json.Unmarshal(jsonResponse.Result, &blockHash); err != nil {
		log.Printf("Unexpected result type: %T\n", jsonResponse.Result)
		return
	}

	writer.Header().Set("Content-Type", "application/json")	
	json.NewEncoder(writer).Encode(blockHash)
}
