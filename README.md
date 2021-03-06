# googlefinance-client-go [![GoDoc](https://godoc.org/github.com/pdevty/googlefinance-client-go?status.svg)](https://godoc.org/github.com/pdevty/googlefinance-client-go)

googlefinance-client-go is a Go client library for google finance.

## Installation

execute:

    $ go get github.com/pdevty/googlefinance-client-go

## Usage

```go
package main

import (
	"context"
	"fmt"
	gf "github.com/pdevty/googlefinance-client-go"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	prices, err := gf.GetPrices(ctx, &gf.Query{P: "1Y", I: "86400", X: "TYO", Q: "7203"})
	if err != nil {
		panic(err)
	}

	fmt.Println(prices)
}
```

Refer to [godoc](http://godoc.org/github.com/pdevty/googlefinance-client-go) for more infomation.

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request