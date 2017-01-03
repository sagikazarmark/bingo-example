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

    bin.Endpoints = []*bingo.Endpoint{
		&bingo.Endpoint{
			Method:      "GET",
			Path:        "/endpoint",
			Description: "Lorem ipsum dolor",
		},
		&bingo.Endpoint{
			Method:      "GET",
			Path:        "/endpoint/:param",
			Description: "Lorem ipsum dolor",
			Parameters: map[string]string{
				"param": "100",
			},
			Handler: func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
				fmt.Fprintf(w, "You wrote: %s", p.ByName("param"))
			},
		},
		&bingo.Endpoint{
			Method:      "POST",
			Path:        "/endpoint/post",
			Description: "...sit amet",
		},
	}

	bin.Groups = []*bingo.Group{
		&bingo.Group{
			Name:        "Group",
			Description: "Lorem ipsum dolor",
			Endpoints: []*bingo.Endpoint{
				&bingo.Endpoint{
					Method:      "GET",
					Path:        "/group/endpoint",
					Description: "Lorem ipsum dolor",
				},
				&bingo.Endpoint{
					Method:      "GET",
					Path:        "/group/endpoint/:param",
					Description: "Lorem ipsum dolor",
					Parameters: map[string]string{
						"param": "100",
					},
					Handler: func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
						fmt.Fprintf(w, "You wrote: %s", p.ByName("param"))
					},
				},
				&bingo.Endpoint{
					Method:      "POST",
					Path:        "/group/endpoint/post",
					Description: "...sit amet",
				},
			},
		},
	}

	log.Println("Starting server on *:" + port)
	log.Fatal(http.ListenAndServe(":"+port, bingo.NewHandler(bin)))
}
