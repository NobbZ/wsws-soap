package stockquote_test

import (
	"fmt"
	"testing"

	"github.com/NobbZ/wsws-soap/stockquote"
)

func TestStockQuote(*testing.T) {
	service := stockquote.NewStockQuotePortType("http://localhost:8080/stockquote", false, nil)
	seasons, err := service.GetLastTradePrice(&stockquote.TradePriceRequest{TickerSymbol: "SIE"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Trade Price: %+v\n", seasons)
}
