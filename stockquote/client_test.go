package stockquote_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/NobbZ/wsws-soap/server"
	"github.com/NobbZ/wsws-soap/stockquote"
)

func TestStockQuote(t *testing.T) {
	go server.Start("127.0.0.1:9999")
	time.Sleep(10 * time.Millisecond) // Wait a bit for the server to start

	service := stockquote.NewStockQuotePortType("http://localhost:9999/stockquote", false, nil)
	seasons, err := service.GetLastTradePrice(&stockquote.TradePriceRequest{TickerSymbol: "SIE"})

	if err != nil {
		panic(err)
	}
	if seasons.Price != 5.0 {
		t.Error("Expected", 5.0, "but got", seasons.Price)
	}
	fmt.Printf("Trade Price: %+v\n", seasons)
}
