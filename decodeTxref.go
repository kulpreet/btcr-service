package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"

	txref "github.com/kulpreet/txref/util"
)

func decodetxref(writer http.ResponseWriter,
	request *http.Request,
	params httprouter.Params) {

	query := params.ByName("query")
	log.Printf("in decodeTxref...%s", query)

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
