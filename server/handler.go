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
	if buf, err := ioutil.ReadAll(r.Body); err != nil {
		log.Println("There was an error in reading the request!", err)
	} else {
		soapEnvelopeRequest := new(stockquote.SOAPEnvelope)
		soapRequest := new(stockquote.TradePriceRequest)
		soapEnvelopeRequest.Body = stockquote.SOAPBody{Content: soapRequest}

		if err = xml.Unmarshal(buf, soapEnvelopeRequest); err != nil {
			log.Println("There was an error in parsing the request!", err)
		}

		soapEnvelopeResponse := new(stockquote.SOAPEnvelope)
		soapResponse := new(stockquote.TradePrice)
		soapResponse.Price = 5.0
		soapEnvelopeResponse.Body = stockquote.SOAPBody{Content: soapResponse}

		encoder := xml.NewEncoder(w)
		if err = encoder.Encode(soapEnvelopeResponse); err != nil {
			log.Println("There was an error in writing the response!", err)
		}
	}
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
