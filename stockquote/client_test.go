package stockquote_test

import (
	"fmt"
	"testing"

	"github.com/NobbZ/wsws-soap/stockquote"
)

func TestStockQuote(t *testing.T) {
	service := stockquote.NewStockQuotePortType("http://localhost:8080/stockquote", false, nil)
	seasons, err := service.GetLastTradePrice(&stockquote.TradePriceRequest{TickerSymbol: "SIE"})
	if err != nil {
		panic(err)
	}
	if seasons.Price != 5.0 {
		t.Error("Expected", 5.0, "but got", seasons.Price)
	}
	fmt.Printf("Trade Price: %+v\n", seasons)
}
