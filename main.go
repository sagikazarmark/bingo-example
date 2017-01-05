package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sagikazarmark/bingo"
	"github.com/julienschmidt/httprouter"
	flag "github.com/spf13/pflag"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8080", "Custom port")
}

func main() {
	bin, _ := bingo.New("bingo", "Bingo example application", "Lorem ipsum")

	AddEndpoints(bin, "")

	group, _ := bingo.NewGroup("Group", "Lorem ipsum")
	AddEndpoints(group, "/group")
	bin.AddGroup(group)

	log.Println("Starting server on *:" + port)
	log.Fatal(http.ListenAndServe(":"+port, bingo.NewHandler(bin)))
}

func AddEndpoints(collector bingo.EndpointCollector, pathPrefix string) {
	endpoint, _ := bingo.NewEndpoint(
		"GET",
		pathPrefix + "/lorem",
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			fmt.Fprintf(w, "Lorem ipsum dolor")
		},
	)
	endpoint.Description = "Lorem ipsum dolor"
	collector.AddEndpoint(endpoint)

	endpoint, _ = bingo.NewEndpoint(
		"GET",
		pathPrefix + "/number/:num",
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			fmt.Fprintf(w, "You wrote: %s", p.ByName("num"))
		},
	)
	endpoint.Description = "Reports back your number"
	endpoint.Parameters.Set("num", "100")
	collector.AddEndpoint(endpoint)

	endpoint, _ = bingo.NewEndpoint(
		"GET",
		pathPrefix + "/query?key=:value",
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			fmt.Fprintf(w, "You wrote: %s", r.URL.RawQuery)
		},
	)
	endpoint.Description = "Reports back your query"
	endpoint.Query.Set("key", "myvalue")
	collector.AddEndpoint(endpoint)

	endpoint, _ = bingo.NewEndpoint(
		"POST",
		pathPrefix + "/post",
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			fmt.Fprintf(w, "You posted")
		},
	)
	endpoint.Description = "Reports that you posted"
	collector.AddEndpoint(endpoint)
}
