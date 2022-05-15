package main

import (
	"EngineREST/routes"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func main() {
	//Read params
	//args := os.Args[1:]

	//const node = "https://engine.rishipanthee.com"

	//if len(args) < 2 {
	//	fmt.Println("Usage: ./engineREST <PORTNUMBER>")
	//	os.Exit(1)
	//}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You have accessed the Hive Engine REST API. All data fetches that could be done using post requests to a JSONRPC server can be done here but using simple GET requests."))
	})
	r.Mount("/blockchain", routes.BlockchainRouter())
	r.Mount("/nft", routes.NFTRouter())
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("The page you attempted to access doesn't exist. We don't have proper documentation yet. We'll let you know when we do."))
	})
	http.ListenAndServe(":3000", r)

}
