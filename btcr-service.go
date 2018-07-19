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
