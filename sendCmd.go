package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/btcsuite/btcd/btcjson"	
)

func sendCommand(marshalledCmd []byte)  (jsonResponse btcjson.Response, err error) {

	req, err := http.NewRequest("POST", "https://host.opdup.com:18443", bytes.NewBuffer(marshalledCmd))
	req.SetBasicAuth(Username, Password)

	response, err := httpClient.Do(req)
    if err != nil && response == nil {
        log.Fatalf("Error sending request to API endpoint. %+v", err)
    }	

    defer response.Body.Close()
	
	body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatalf("Couldn't parse response body. %+v", err)
    }
	
    log.Println("Response Body:", string(body))
	if err = json.Unmarshal(body, &jsonResponse); err != nil {
		fmt.Printf("Error unmarshalling response: %s\n", err)
		return
	}

	if jsonResponse.Error != nil {
		fmt.Printf("JSON response error: %s\n", jsonResponse.Error)
		err = jsonResponse.Error
		return 
	}

	return
}
