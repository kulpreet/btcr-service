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

This is still a WIP, see /tip for now

`https://localhost:8080/txref/<txref>/resolve`

## Following a tip

`https://localhost:8080/txref/<txref>/tip`

The above will list all transactions matching the address in the vout
from the txref. We assume there is only one address in the vout for
now, as we are only focussed on P2PKH for the current MVP.

By default this endpoint only shows spending transactions from the
list of transactions originating at txref.

The same endpoint can also show all the transactions that were
considered for following the tip by passing a spendsOnly query string
parameter, spendsOnly, as false. For example: 

`https://localhost:8080/txref/<txref>/tip?spendsOnly=false`


## Decoding a Txref

`https://localhost:8080/txref/<TxRef>/decode`

## Get Txid from Txref

`https://localhost:8080/txref/<TxRef>/txid`

## Get decoded Tx from Txid

`https://localhost:8080/tx/<Txid>`

## Txid to UTXOs for the address in Txid

`https://localhost:8080/addr/<addr>/spends`
