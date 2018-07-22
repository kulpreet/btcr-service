# Simple HTTP Service to query txrefs

## Install

`go get github.com/kulpreet/btcr-service`

## Start Service

```
cd $GOPATH/github.com/kulpreet/btcr-service
go build
./btcr-service
```

## Resolve a BTCR DID

`https://localhost:8080/txref/<txref>/resolve`


## Decoding a Txref

`https://localhost:8080/txref/<TxRef>/decode`

## Get Txid from Txref

`https://localhost:8080/txref/<TxRef>/txid`

## Get decoded Tx from Txid

`https://localhost:8080/tx/<Txid>`

## Txid to UTXOs for the address in Txid

`https://localhost:8080/addr/<addr>/spends`
