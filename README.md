# go-gemini-api

A Go library for interacting with the Gemini API. An example can be found in the examples directory or below.

```
package main

import (
	"fmt"
	"time"

	"github.com/eightseventhreethree/go-gemini-api/pkg/gemini"
	"github.com/eightseventhreethree/go-gemini-api/pkg/handlers"
)

func main() {
	api := gemini.NewClient(&gemini.ClientOpts{
		BaseURL:    "https://api.sandbox.gemini.com",
		Timeout:    time.Second.Round(10),
		RetryLimit: 3,
		RetryDelay: time.Second.Round(3),
	})

	resp, err := api.GetSymbols()
	handlers.CheckErrLog(err, "failed to GetSymbols")
	for _, v := range resp.Symbols {
		sym := &gemini.SymbolRequest{Name: v}
		symDetails, err := api.GetSymbolDetails(sym)
		handlers.CheckErrLog(err, "failed to call GetSymbolDetails")
		if symDetails.Status != gemini.Closed {
			tickerResp, err := api.GetTickerV2(sym)
			handlers.CheckErrLog(err, "failed to call GetTickerV1")
			tickerv1Resp, _ := api.GetTickerV1(sym)
			fmt.Printf("Symbol: %s -> Changes: %v -> Volume: %v \n", v, tickerResp.Changes, tickerv1Resp.Volume)
		}
	}
}
```

## Contributing:
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License:
[GPLv2](https://www.gnu.org/licenses/old-licenses/gpl-2.0.en.html)