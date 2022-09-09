package routes

import (
	"encoding/json"
	"github.com/cfoxon/hiveenginego"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func BlockchainRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/getLatestBlockInfo", getLatestBlockInfo())
	r.Get("/getBlockRange", getBlockRange())
	return r
}

func getLatestBlockInfo() http.HandlerFunc {
	const node = "https://engine.rishipanthee.com"
	herpc := hiveenginego.NewHiveEngineRpc(node)
	return func(w http.ResponseWriter, r *http.Request) {
		currentBlockInfo, _ := herpc.GetLatestBlockInfo()
		js, _ := json.Marshal(currentBlockInfo)
		w.Write(js)
	}
}

func getBlockRange() http.HandlerFunc {
	const node = "https://engine.rishipanthee.com"
	herpc := hiveenginego.NewHiveEngineRpc(node)
	return func(w http.ResponseWriter, r *http.Request) {
		start, _ := strconv.Atoi(r.URL.Query().Get("start"))
		end, _ := strconv.Atoi(r.URL.Query().Get("end"))
		blockRange, _ := herpc.GetBlockRange(start, end)
		var arr []json.RawMessage
		for _, block := range blockRange {
			arr = append(arr, block)
		}
		js, _ := json.Marshal(arr)
		w.Write(js)
	}
}
