package routes

import (
	"encoding/json"
	"github.com/cfoxon/hiveenginego"
	"github.com/go-chi/chi"
	"net/http"
)

func NFTRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/getSymbolAllNFT", getSymbolAllNFT())
	return r
}

func getSymbolAllNFT() http.HandlerFunc {
	const node = "https://engine.rishipanthee.com"
	herpc := hiveenginego.NewHiveEngineRpc(node)
	return func(w http.ResponseWriter, r *http.Request) {
		symbol := r.URL.Query().Get("symbol")
		allNFTs, _ := herpc.GetSymbolAllNft(symbol)
		js, _ := json.Marshal(allNFTs)
		w.Write(js)
	}
}
