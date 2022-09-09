package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

func RestRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You have accessed the Hive Engine REST API. All data fetches that could be done using post requests to a JSONRPC server can be done here but using simple GET requests."))
	})
	r.Mount("/blockchain", BlockchainRouter())
	r.Mount("/nft", NFTRouter())
	r.Mount("/stats", StatsRouter())
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("The page you attempted to access doesn't exist. We don't have proper documentation yet. We'll let you know when we do."))
	})
	return r
}
