package routes

import (
	"encoding/json"
	"github.com/cfoxon/hiveenginego"
	"github.com/go-chi/chi"
	"net/http"
)

func StatsRouter() http.Handler {
	const node = "https://engine.rishipanthee.com"
	herpc := hiveenginego.NewHiveEngineRpc(node)
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		currentBlockInfo, _ := herpc.GetLatestBlockInfo()
		js, _ := json.Marshal(currentBlockInfo)
		w.Write(js)
	})
	return r
}
