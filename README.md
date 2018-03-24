# Quarkhash Golang Bindings

## Introduction

Quarkhash is an hash function designed by Colin Percival & used in several crypto currencies.

This library allows using Quarkhash in any Golang software.

## Installation

```bash
$ go get github.com/mycroft/goquarkhash
$ cd $GOPATH/src/github.com/mycroft/goquarkhash
$ make all
```

## Usage

```golang
package main

import (
	"fmt"
	quark "github.com/mycroft/goquarkhash"
)

func main() {
	hash := quark.QuarkHash(make([]byte, 32))
	fmt.Printf("%x\n", hash)
}
```

## Authors

Please see source code files.
