package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/toorop/go-bittrex"
	"gopkg.in/mgo.v2"
)

//HandleFunc http request handler function
type HandleFunc func(w http.ResponseWriter, r *http.Request)

//ContextKey is type for context value keys
type ContextKey string

type Tick struct {
	Value     float64 `json:"value"`
	Timestamp int64   `json:"timestamp"`
}

func getMarketTick(api *bittrex.Bittrex, market string) []byte {
	ticker, err := api.GetTicker(market)
	if err != nil {
		jsonError, _ := json.Marshal(map[string]string{"error": err.Error()})
		return jsonError
	}
	timestamp := time.Now()
	tick := Tick{Value: ticker.Last, Timestamp: timestamp.Unix()}
	jsonResponse, _ := json.Marshal(tick)
	return jsonResponse
}

func mongoInjector(mongo *mgo.Session, handler HandleFunc) HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), ContextKey("mongo"), mongo)
		r = r.WithContext(ctx)
		handler(w, r)
	}
}

func main() {
	apiClient := bittrex.New("a9be0dc25f7b4c36945744b0b8ae3f96", "00345964c1974fde936e8cae400801bf")
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/api/btc", func(w http.ResponseWriter, r *http.Request) {
		w.Write(getMarketTick(apiClient, "USDT-BTC"))
	})
	serveMux.HandleFunc("/api/eth", func(w http.ResponseWriter, r *http.Request) {
		w.Write(getMarketTick(apiClient, "USDT-ETH"))
	})
	serveMux.HandleFunc("/api/ltc", func(w http.ResponseWriter, r *http.Request) {
		w.Write(getMarketTick(apiClient, "USDT-LTC"))
	})

	http.ListenAndServe("0.0.0.0:3000", serveMux)
}
