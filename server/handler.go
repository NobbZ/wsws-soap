package server

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/NobbZ/wsws-soap/stockquote"
)

func handleGetLastTradePrice(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting last TradePrice")
	buf, _ := ioutil.ReadAll(r.Body)
	soapEnvelopeRequest := new(stockquote.SOAPEnvelope)
	soapRequest := new(stockquote.TradePriceRequest)
	soapEnvelopeRequest.Body = stockquote.SOAPBody{Content: soapRequest}

	xml.Unmarshal(buf, soapEnvelopeRequest)

	soapEnvelopeResponse := new(stockquote.SOAPEnvelope)
	soapResponse := new(stockquote.TradePrice)
	soapResponse.Price = 5.0
	soapEnvelopeResponse.Body = stockquote.SOAPBody{Content: soapResponse}

	encoder := xml.NewEncoder(w)
	encoder.Encode(soapEnvelopeResponse)
}

func handlePostStockquote(w http.ResponseWriter, r *http.Request) {
	switch r.Header["Soapaction"][0] {
	case "http://example.com/GetLastTradePrice":
		handleGetLastTradePrice(w, r)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		handlePostStockquote(w, r)
	default:
		panic("I don't GET it")
	}
}
