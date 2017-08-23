package main

import (
	"context"
	"net/http"

	"github.com/toorop/go-bittrex"
	"gopkg.in/mgo.v2"
)

//HandleFunc http request handler function
type HandleFunc func(w http.ResponseWriter, r *http.Request)

//ContextKey is type for context value keys
type ContextKey string

func getBitcoinToRouble(w http.ResponseWriter, r *http.Request) {

}

func mongoInjector(mongo *mgo.Session, handler HandleFunc) HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), ContextKey("mongo"), mongo)
		r = r.WithContext(ctx)
		handler(w, r)
	}
}

func main() {
	mongodb, err := mgo.Dial("mongo")
	if err != nil {
		panic(err.Error())
	}
	apiClient := bittrex.New("a9be0dc25f7b4c36945744b0b8ae3f96", "00345964c1974fde936e8cae400801bf")
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/api/btc", func(w http.ResponseWriter, r *http.Request) {
		apiClient.GetMarketSummary("USDT-BTC")
	})
	serveMux.HandleFunc("/api/doge", func(w http.ResponseWriter, r *http.Request) {
		apiClient.GetMarketSummary("USDT-DOGE")
	})
	serveMux.HandleFunc("/api/ltc", mongoInjector(mongodb, func(w http.ResponseWriter, r *http.Request) {
		apiClient.GetMarketSummary("USDT-LTC")
	}))
	http.ListenAndServe("0.0.0.0:3000", serveMux)
}
