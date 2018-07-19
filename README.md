# Simple HTTP Service to query txrefs

## Install

`go get github.com/kulpreet/btcr-service`

## Start Service

```
cd $GOPATH/github.com/kulpreet/btcr-service
go build
./btcr-service
```

## Txref to Txid

### Decoding a Txref

`https://localhost:8080/txref/<TxRef>/decode`

### Getting Txid from Txref

`https://localhost:8080/txref/<TxRef>/txid`

## Txid to UTXOs for the address in Txid


