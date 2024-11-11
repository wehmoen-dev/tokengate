# Axie Infinity Tokengate Client

This package provides a simple client that wraps the tokengate API. 

## Installation

```bash
go get github.com/wehmoen-dev/tokengate
```

## Usage

```go
package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/wehmoen-dev/tokengate"
	"log"
)

var testAddress = common.HexToAddress("0x2d62c27ce2e9e66bb8a667ce1b60f7cb02fa9810")

func main() {
	client := tokengate.New()
	hasMystic, err := client.HasAxieFromCollection(testAddress, tokengate.MysticAxie)
	if err != nil {
		log.Fatal(err)
	}

	if hasMystic {
		log.Printf("Address %s has a mystic axie", testAddress.Hex())
	} else {
		log.Printf("Address %s does not have a mystic axie", testAddress.Hex())
	}
}
```

## Notes

Land functions are currently not implemented. I first have to update the upstream API proxy.